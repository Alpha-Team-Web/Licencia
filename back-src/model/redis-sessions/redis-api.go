package redis_sessions

import "back-src/model/redis-sessions/databases"

type RedisApi struct {
	AuthTokenDB *databases.RedisAuthTokenDb
	ProfileDB   *databases.RedisProfileDb
}

const (
	redisHostAddr = "localhost:6379"
	redisPassword = ""
)

func NewRedisApi() *RedisApi {
	return &RedisApi{
		AuthTokenDB: databases.NewRedisAuthTokenDB(
			redisHostAddr,
			redisPassword,
		),
		ProfileDB: databases.NewRedisProfileDB(
			redisHostAddr,
			redisPassword,
		),
	}
}
