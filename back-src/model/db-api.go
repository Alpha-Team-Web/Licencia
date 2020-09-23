package model

import (
	redis_sessions "back-src/model/redis-sessions"
	"back-src/model/sql"
)

type DbApi struct {
	SqlDb   *sql.Database
	RedisDb *redis_sessions.RedisApi
}
