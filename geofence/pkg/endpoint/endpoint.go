package endpoint

import (
	"context"
	"errors"
	"github.com/amirnejati/map-services/geofence/pkg/service"
	kitendpoint "github.com/go-kit/kit/endpoint"
)

type Point2DistrictRequest struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type Point2DistrictResponse struct {
	District service.District `json:"district"`
	Err      string           `json:"err,omitempty"`
}

// MakePoint2DistrictEndpoint returns an endpoint that invokes Point2District on the service.
func MakePoint2DistrictEndpoint(svc service.GeofenceService) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Point2DistrictRequest)
		dist, err := svc.Point2District(ctx, req.Latitude, req.Longitude)
		if err != nil {
			return Point2DistrictResponse{dist, err.Error()}, nil
		}
		return Point2DistrictResponse{
			District: dist,
			Err:      "",
		}, nil
	}
}

// Point2District implements Service. Primarily useful in a client.
func (es *EndpointsSet) Point2District(ctx context.Context, latitude float32, longitude float32) (service.District, error) {
	request := Point2DistrictRequest{
		Latitude:  latitude,
		Longitude: longitude,
	}
	response, err := es.Point2DistrictEndpoint(ctx, request)
	if err != nil {
		return service.District{}, err
	}

	getResp := response.(Point2DistrictResponse)
	if getResp.Err != "" {
		return service.District{}, errors.New(getResp.Err)
	}
	return response.(Point2DistrictResponse).District, nil
}
