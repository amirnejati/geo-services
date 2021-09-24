package service

import (
	"context"
	kitlog "github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//var ErrRepo = errors.New("unable to handle repoMongo request")

type repoMongo struct {
	db       *mongo.Database
	collName string
	logger   kitlog.Logger
}

func NewRepoMongo(db *mongo.Database, collName string, logger kitlog.Logger) (RepositoryMongo, error) {
	return &repoMongo{
		db:       db,
		collName: collName,
		logger:   kitlog.With(logger, "repoMongo", "mongodb"),
	}, nil
}

func (repoMongo *repoMongo) Point2District(ctx context.Context, point Point) (District, error) {
	coll := repoMongo.db.Collection(repoMongo.collName)

	query := bson.M{
		"polygons": bson.M{
			"$geoIntersects": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float32{point.Longitude, point.Latitude},
				},
			},
		},
	}

	data := coll.FindOne(ctx, query)
	var dist District

	if err := data.Decode(&dist); err != nil {
		return District{}, err
	}

	return dist, nil
}
