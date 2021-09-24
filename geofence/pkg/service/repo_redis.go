package service

import (
	"context"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-redis/redis/v8"
)

type repoRedis struct {
	db     *redis.Client
	logger kitlog.Logger
}

func NewRepoRedis(db *redis.Client, logger kitlog.Logger) (RepositoryRedis, error) {
	return &repoRedis{
		db:     db,
		logger: kitlog.With(logger, "repoRedis", "Redis"),
	}, nil
}

func (repoRedis *repoRedis) SetPriceEfficient(ctx context.Context, districtNo uint8, priceCoEfficient float32) error {
	//districtId := strconv.FormatInt(int64(districtNo), 10)
	res := repoRedis.db.Set(ctx, string(districtNo), priceCoEfficient, 0)
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}

func (repoRedis *repoRedis) GetPriceEfficient(ctx context.Context, districtNo uint8) (float32, error) {
	//districtId := strconv.FormatInt(int64(districtNo), 10)
	res := repoRedis.db.Get(ctx, string(districtNo))
	if res.Err() != nil {
		return 0, res.Err()
	}
	return res.Float32()
}
