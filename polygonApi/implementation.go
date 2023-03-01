package polygonApi

import (
	"context"
	"encoding/json"
	"polygon-service-gokit/common"

	pq "github.com/lib/pq"
	"github.com/rs/zerolog"
)

type service struct {
	repository Repository
	logger     zerolog.Logger
}

func NewService(rep Repository, logger zerolog.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s *service) CreatePolygon(ctx context.Context, fc *FeatureCollection) error {

	log := s.logger.With().Str("method", "CreatePolygon").Logger()
	log.Info().Msgf("request body: feature collection: %v", fc)

	err := s.repository.Save(ctx, fc)
	if err != nil {
		if err.(*pq.Error).Code == pq.ErrorCode(common.DB_UNIQUE_CONSTRAINT_VIOLATION) {
			log.Error().Err(err).Msg(common.ErrDuplicateEntry.Error())
			return common.ErrDuplicateEntry
		}
		log.Error().Err(err).Msg(common.ErrQueryRepository.Error())
		return common.ErrQueryRepository
	}

	return nil
}

func (s *service) GetPolygon(ctx context.Context, id string, name string) (*FeatureCollection, error) {

	result, err := s.repository.GetPolygon(ctx, id, name)
	if err != nil {
		return nil, err
	}

	var geom geometry
	err = json.Unmarshal([]byte(result.Geom), &geom)
	if err != nil {
		return nil, err
	}

	var fc FeatureCollection

	fc.Type = "FeatureCollection"
	fc.Features = append(fc.Features, &Feature{
		Type:     "Feature",
		Geometry: geom,
		Properties: properties{
			Name: result.Name,
		},
	})

	return &fc, nil
}
