package redis_sessions

import "back-src/model/redis-sessions/databases"

type RedisApi struct {
	AuthTokenDb *databases.RedisAuthTokenDb
	ProfileDb   *databases.RedisProfileDb
	FilterDb    *databases.RedisFilterDb
}

const (
	redisHostAddr = "localhost:6379"
	redisPassword = ""
)

func NewRedisApi() *RedisApi {
	return &RedisApi{
		AuthTokenDb: databases.NewRedisAuthTokenDB(
			redisHostAddr,
			redisPassword,
		),
		ProfileDb: databases.NewRedisProfileDB(
			redisHostAddr,
			redisPassword,
		),
		FilterDb: databases.NewRedisFilterDb(
			redisHostAddr,
			redisPassword,
		),
	}
}
