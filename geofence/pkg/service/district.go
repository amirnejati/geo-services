package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Point struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude" `
}

type District struct {
	PolygonId  primitive.ObjectID `json:"polygon_id" bson:"_id"`
	Name       string             `json:"name"  bson:"name"`
	DistrictNo uint8              `json:"district_no" bson:"region_id"`
}

type RepositoryMongo interface {
	Point2District(ctx context.Context, point Point) (District, error)
}

//type RepositorySQL interface {
//
//}

type RepositoryRedis interface {
}

//func (d District) Stringify() string {
//	b, err := json.Marshal(d)
//	if err != nil {
//		return "unsupported value type"
//	}
//	return string(b)
//}
