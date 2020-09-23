package databases

import (
	"back-src/controller/utils/libs"
	"back-src/model/redis-sessions/orm"
	"back-src/view/data"
	"time"
)

const filterDbId = 6

type RedisFilterDb struct {
	conn *RedisCli
}

const (
	filterDbKeysExpireMinutes = 10
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

func (db *RedisFilterDb) IsThereFilter(userWithRole string) bool {
	cmd := db.conn.Exists(userWithRole)
	return libs.Ternary(cmd.Val() == 0, false, true).(bool)
}

func (db *RedisFilterDb) GetFilter(userWithRole string) (data.Filter, []string, error) {
	if values, err := db.conn.HGetAll(userWithRole).Result(); err != nil {
		return data.Filter{}, nil, err
	} else {
		filter, projectIds := orm.UnHashFilter(values)
		return filter, projectIds, err
	}
}

func (db *RedisFilterDb) ExtendFilterExpiry(userWithRole string) error {
	return db.conn.Expire(userWithRole, filterDbKeysExpireMinutes).Err()
}
