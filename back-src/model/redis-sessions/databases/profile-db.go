package databases

import (
	"back-src/model/existence"
	"back-src/model/redis-sessions/orm"
)

//DB Number 5
const profileDbId = 5

type RedisProfileDb struct {
	conn *RedisCli
}

const (
	profileSetKey = "profile"
)

func NewRedisProfileDB(addr, password string) *RedisProfileDb {
	redisProfileDb := &RedisProfileDb{
		conn: NewRedisCli(addr, password, profileDbId),
	}
	if stats := redisProfileDb.conn.FlushDB(); stats.Err() != nil {
		panic(stats.Err())
	}
	return redisProfileDb
}

func (db *RedisProfileDb) GetProfile(userWithRole string) (existence.Profile, error) {
	if values, err := db.conn.HGetAll(userWithRole).Result(); err != nil {
		return existence.Profile{}, err
	} else {
		return orm.UnHashProfileImage(values), nil
	}
}

func (db *RedisProfileDb) SetProfile(userWithRole string, profile existence.Profile) error {
	if cmd := db.conn.SAdd(profileSetKey, userWithRole); cmd.Err() != nil {
		return cmd.Err()
	}

	if stats := db.conn.HMSet(
		userWithRole,
		orm.HashProfileImage(profile),
	); stats.Err() != nil {
		return stats.Err()
	}
	return nil
}

func (db *RedisProfileDb) DeleteProfile(userWithRole string) error {
	if cmd := db.conn.SRem(profileSetKey, userWithRole); cmd.Err() != nil {
		return cmd.Err()
	}

	if cmd := db.conn.Del(userWithRole); cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func (db *RedisProfileDb) IsThereProfile(userWithRole string) (bool, error) {
	return db.conn.SIsMember(profileSetKey, userWithRole).Result()
}
