package http

import (
	"github.com/amirnejati/map-services/geofence/pkg/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on predefined paths.
func NewHTTPHandler(endpoints endpoint.EndpointsSet, options map[string][]kithttp.ServerOption) http.Handler {
	m := http.NewServeMux()
	makePoint2DistrictHandler(m, endpoints, options["Point2District"])
	return m
}
