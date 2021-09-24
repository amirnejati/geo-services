package http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/amirnejati/map-services/geofence/pkg/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

// makePoint2DistrictHandler creates the handler logic
func makePoint2DistrictHandler(m *http.ServeMux, endpoints endpoint.EndpointsSet, options []kithttp.ServerOption) {
	m.Handle("/point2district",
		kithttp.NewServer(
			endpoints.Point2DistrictEndpoint,
			decodePoint2DistrictRequest,
			encodePoint2DistrictResponse,
			options...,
		),
	)
}

// decodePoint2DistrictRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePoint2DistrictRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if r.ContentLength <= 0 {
		return nil, errors.New("request without body is not acceptable")
	}
	req := endpoint.Point2DistrictRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodePoint2DistrictResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePoint2DistrictResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		EncodeError(ctx, e, w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
