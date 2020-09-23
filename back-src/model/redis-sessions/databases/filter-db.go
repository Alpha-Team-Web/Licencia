package databases

import (
	"back-src/model/redis-sessions/orm"
	"back-src/view/data"
	"time"
)

const filterDbId = 6

type RedisFilterDb struct {
	conn *RedisCli
}

const (
	filterDbKeysExpireMinutes = 5
)

func NewRedisFilterDb(addr, password string) *RedisFilterDb {
	redisFilterDb := &RedisFilterDb{
		conn: NewRedisCli(addr, password, filterDbId),
	}
	if stats := redisFilterDb.conn.FlushDB(); stats.Err() != nil {
		panic(stats.Err())
	}
	return redisFilterDb
}

func (db *RedisFilterDb) AddFilterToUserWithRole(userWithRole string, filter data.Filter, projectIds []string) error {
	if err := db.conn.HMSet(userWithRole, orm.HashFilter(filter, projectIds)).Err(); err != nil {
		return err
	}
	if cmd := db.conn.Expire(userWithRole, time.Minute*filterDbKeysExpireMinutes); cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}
