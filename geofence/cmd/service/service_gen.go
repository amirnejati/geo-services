package service

import (
	"github.com/amirnejati/map-services/geofence/pkg/endpoint"
	httptransport "github.com/amirnejati/map-services/geofence/pkg/http"
	"github.com/amirnejati/map-services/geofence/pkg/service"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/oklog/oklog/pkg/group"
)

func createService(endpoints endpoint.EndpointsSet) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	initGRPCHandler(endpoints, g)
	return g
}

func defaultHttpOptions(logger log.Logger) map[string][]kithttp.ServerOption {
	options := map[string][]kithttp.ServerOption{
		"Point2District": {
			kithttp.ServerErrorEncoder(httptransport.EncodeError),
			//kithttp.ServerErrorLogger(logger),
			//kithttp.ServerBefore(
			//	opentracing.HTTPToContext(tracer, "Point2District", logger),
			//),
		},
	}
	return options
}

func defaultGRPCOptions(logger log.Logger) map[string][]kitgrpc.ServerOption {
	options := map[string][]kitgrpc.ServerOption{
		"Point2District": {
			//kitgrpc.ServerErrorHandler(grpctransport.EncodeError),
			//kitgrpc.ServerErrorLogger(logger),
			//kitgrpc.ServerBefore(
			//	opentracing.HTTPToContext(tracer, "Point2District", logger),
			//),
		},
	}
	return options
}

func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}

func addDefaultEndpointMiddleware(logger log.Logger, mw map[string][]kitendpoint.Middleware) {
	mw["Point2District"] = []kitendpoint.Middleware{
		endpoint.LoggingMiddleware(log.With(logger, "method", "Point2District")),
		endpoint.ValidatePoint(),
	}
}
