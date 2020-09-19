package databases

import "github.com/go-redis/redis"

type RedisCli struct {
	*redis.Client
	DbId int
}

func NewRedisCli(addr, password string, dbId int) *RedisCli {
	return &RedisCli{
		Client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       dbId,
		}),
		DbId: dbId,
	}
}
