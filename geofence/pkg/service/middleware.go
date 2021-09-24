package service

import (
	"context"
	"fmt"
	kitlog "github.com/go-kit/kit/log"
	"time"
)

// Middleware describes a service middleware.
type Middleware func(GeofenceService) GeofenceService

type loggingMiddleware struct {
	logger kitlog.Logger
	next   GeofenceService
}

// LoggingMiddleware takes a logger as a dependency and returns a GeofenceService Middleware.
func LoggingMiddleware(logger kitlog.Logger) Middleware {
	return func(next GeofenceService) GeofenceService {
		return &loggingMiddleware{logger, next}
	}
}

func (lmw loggingMiddleware) Point2District(ctx context.Context, latitude float32, longitude float32) (dist District, err error) {
	defer func(begin time.Time) {
		lmw.logger.Log(
			"method", "Point2District",
			"latitude", latitude,
			"longitude", longitude,
			"district", fmt.Sprintf("%+v", dist),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	dist, err = lmw.next.Point2District(ctx, latitude, longitude)
	return
}
