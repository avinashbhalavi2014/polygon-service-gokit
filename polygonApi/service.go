package polygonApi

import (
	"context"
)

type Service interface {
	CreatePolygon(ctx context.Context, payload *FeatureCollection) error
	GetPolygon(ctx context.Context, id string, name string) (*FeatureCollection, error)
}
