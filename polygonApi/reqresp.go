package polygonApi

type (
	CreatePolygonRequest struct {
		FeatureCollection *FeatureCollection
	}

	GetPolygonRequest struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	PolygonResponse struct {
		FeatureCollection *FeatureCollection
		Err               error `json:"error,omitempty"`
	}
)
