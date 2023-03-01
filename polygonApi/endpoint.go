package polygonApi

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreatePolygon endpoint.Endpoint
	GetPolygon    endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreatePolygon: makeCreatePolygonEndpoint(s),
		GetPolygon:    makeGetPolygonEndpoint(s),
	}
}

func makeCreatePolygonEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(CreatePolygonRequest)
		if !ok {
			return nil, errors.New("invalid json input")
		}
		err := s.CreatePolygon(ctx, req.FeatureCollection)
		return nil, err
	}
}

func makeGetPolygonEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPolygonRequest)
		polygon, err := s.GetPolygon(ctx, req.ID, req.Name)
		return PolygonResponse{FeatureCollection: polygon}, err
	}
}
