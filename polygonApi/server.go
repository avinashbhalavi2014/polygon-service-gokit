package polygonApi

import (
	"context"
	"encoding/json"
	"net/http"
	"polygon-service-gokit/common"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeErrorResponse),
	}

	r.Methods("POST").Path("/api/create-polygon").Handler(httptransport.NewServer(
		endpoints.CreatePolygon,
		decodeCreateRequest,
		encodeCreateResponse,
		options...,
	))

	r.Methods("GET").Path("/api/get-polygon").Handler(httptransport.NewServer(
		endpoints.GetPolygon,
		decodeGetRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodeCreateRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req CreatePolygonRequest
	if e := json.NewDecoder(r.Body).Decode(&req.FeatureCollection); e != nil {
		return nil, e
	}
	return req, nil
}

func encodeCreateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err, ok := response.(customErr); ok && err.error() != nil {
		encodeErrorResponse(ctx, err.error(), w)
		return nil
	}
	w.WriteHeader(http.StatusCreated)
	return nil
}

func decodeGetRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	return GetPolygonRequest{ID: id, Name: name}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err, ok := response.(customErr); ok && err.error() != nil {
		encodeErrorResponse(ctx, err.error(), w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

// commonMiddleware...
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func encodeErrorResponse(ctx context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("content-Type", "application/Json; charset=utf-8")
	w.WriteHeader(getStatusCode(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

type customErr interface {
	error() error
}

// getStatusCode funciton return the relevant status code
func getStatusCode(err error) int {
	switch err {
	case common.ErrInvalidKeyValue, common.ErrInvalidLocaleValue,
		common.ErrInvalidValue, common.ErrRecordNotFound,
		common.ErrRequestParamBody, common.ErrBadRouting,
		common.ErrDuplicateEntry:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
