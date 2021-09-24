package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	dbMongo *mongo.Database
	dbRedis *redis.Client
	//dbSQL *sql.DB
)

func GetMongoDB() (*mongo.Database, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(mongoUri)
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Secondary()); err != nil {
		return nil, err
	}

	dbMongo = client.Database(mongoDbName)
	return dbMongo, nil
}

func GetRedisDB() *redis.Client {
	dbRedis = redis.NewClient(&redis.Options{
		Addr:     redisUri,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return dbRedis
}
