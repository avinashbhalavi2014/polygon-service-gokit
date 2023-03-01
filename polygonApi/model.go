package polygonApi

import (
	"context"
)

type Areas struct {
	ID   string `json:"id,omitempty" db:"id"`
	Name string `json:"name" db:"name"`
	Geom string `json:"geom" db:"geom"`
}

type geometry struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"`
}

type properties struct {
	Name string `json:"name"`
}

type Feature struct {
	Type       string     `json:"type"`
	Geometry   geometry   `json:"geometry"`
	Properties properties `json:"properties"`
}

type FeatureCollection struct {
	Type     string     `json:"type"`
	Features []*Feature `json:"features"`
}

type Result struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Geom string `json:"geom"`
}

type Repository interface {
	Save(ctx context.Context, fc *FeatureCollection) error
	GetPolygon(ctx context.Context, ID string, name string) (*Result, error)
}
