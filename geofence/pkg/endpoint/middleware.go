package endpoint

import (
	"context"
	"errors"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"time"
)

// LoggingMiddleware returns an endpoint middleware that logs the duration of each invocation, and the resulting error, if any.
func LoggingMiddleware(logger log.Logger) kitendpoint.Middleware {
	return func(next kitendpoint.Endpoint) kitendpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Log(
					"transport_error", err,
					"took", time.Since(begin),
				)
			}(time.Now())
			return next(ctx, request)
		}
	}
}

// ValidatePoint checks coordinates to be valid
func ValidatePoint() kitendpoint.Middleware {
	return func(next kitendpoint.Endpoint) kitendpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			pointRequest, ok := request.(Point2DistrictRequest)
			if !ok ||
				pointRequest.Latitude < -90.0 || pointRequest.Latitude > 90.0 ||
				pointRequest.Longitude < -180.0 || pointRequest.Longitude > 180.0 {
				return nil, ErrPointFailure
			}
			return next(ctx, request)
		}
	}
}

var (
	// ErrPointFailure signifies that an auth token was missing or invalid
	ErrPointFailure = errors.New("latitude or longitude was missing or invalid")
)
