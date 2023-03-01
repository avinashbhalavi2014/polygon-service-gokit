package polygonApi

import (
	"context"
	"encoding/json"
	"fmt"
	"polygon-service-gokit/util"

	"github.com/jmoiron/sqlx"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/ewkb"
	"github.com/paulmach/orb/geojson"
)

type repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) Save(ctx context.Context, item *FeatureCollection) error {

	id, err := util.GenerateID()
	if err != nil {
		return err
	}

	byteData, err := json.Marshal(item)
	if err != nil {
		return err
	}

	fc := geojson.NewFeatureCollection()
	err = json.Unmarshal(byteData, &fc)
	if err != nil {
		return err
	}

	for _, feature := range fc.Features {

		polygon := feature.Geometry.(orb.Polygon)

		_, err := r.db.DB.Exec("INSERT INTO areas (id, name, polygon) VALUES ($1,$2,ST_GeomFromEWKB($3))", id, feature.Properties.MustString("name"), ewkb.Value(polygon, 4326))
		if err != nil {
			return err
		}

	}

	return nil
}

func (r *repo) GetPolygon(ctx context.Context, id string, name string) (*Result, error) {

	query := `SELECT id, name, ST_AsGeoJSON(polygon) as geom FROM areas`

	if id != "" && name != "" {
		query += fmt.Sprintf(" WHERE id = '%s' AND name = '%s'", id, name)
	} else if id != "" {
		query += fmt.Sprintf(" WHERE id = '%s'", id)
	} else if name != "" {
		query += fmt.Sprintf(" WHERE name = '%s'", name)
	}

	row := r.db.DB.QueryRow(query)

	var result Result

	err := row.Scan(&result.ID, &result.Name, &result.Geom)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
