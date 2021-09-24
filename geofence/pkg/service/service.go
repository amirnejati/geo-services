package service

import (
	"context"
	kitlog "github.com/go-kit/kit/log"
)

type GeofenceService interface {
	Point2District(ctx context.Context, latitude, longitude float32) (District, error)
}

type geofenceService struct {
	repository RepositoryMongo
	logger     kitlog.Logger
}

// NewGeofenceService returns a naive, stateless implementation of GeofenceService.
func NewGeofenceService(rep RepositoryMongo, logger kitlog.Logger) GeofenceService {
	return &geofenceService{repository: rep, logger: logger}
}

// NewGeofenceServiceMiddlewareEnabled returns a GeofenceService with all of the expected middleware wired in.
func NewGeofenceServiceMiddlewareEnabled(rep RepositoryMongo, logger kitlog.Logger, middleware []Middleware) GeofenceService {
	var svc GeofenceService = NewGeofenceService(rep, logger)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (s *geofenceService) Point2District(ctx context.Context, latitude float32, longitude float32) (District, error) {
	dist, err := s.repository.Point2District(ctx, Point{Latitude: latitude, Longitude: longitude})
	if err != nil {
		return District{}, err
	}
	return dist, nil
}
