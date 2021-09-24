package endpoint

import (
	"github.com/amirnejati/map-services/geofence/pkg/service"
	kitendpoint "github.com/go-kit/kit/endpoint"
)

// EndpointsSet collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type EndpointsSet struct {
	Point2DistrictEndpoint kitendpoint.Endpoint
}

// NewEndpointsSet returns a EndpointsSet struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func NewEndpointsSet(svc service.GeofenceService, mdw map[string][]kitendpoint.Middleware) EndpointsSet {
	eps := EndpointsSet{
		Point2DistrictEndpoint: MakePoint2DistrictEndpoint(svc),
	}
	for _, m := range mdw["Point2District"] {
		eps.Point2DistrictEndpoint = m(eps.Point2DistrictEndpoint)
	}
	return eps
}
