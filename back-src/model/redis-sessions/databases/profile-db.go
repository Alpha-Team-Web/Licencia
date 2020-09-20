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

func (db *RedisProfileDb) GetProfile(username string) (existence.Profile, error) {
	if values, err := db.conn.HGetAll(username).Result(); err != nil {
		return existence.Profile{}, err
	} else {
		return orm.UnHashProfileImage(values), nil
	}
}

func (db *RedisProfileDb) SetProfile(username string, profile existence.Profile) error {
	if cmd := db.conn.SAdd(profileSetKey, username); cmd.Err() != nil {
		return cmd.Err()
	}

	if stats := db.conn.HMSet(
		username,
		orm.HashProfileImage(profile),
	); stats.Err() != nil {
		return stats.Err()
	}
	return nil
}

func (db *RedisProfileDb) DeleteProfile(username string) error {
	if cmd := db.conn.SRem(profileSetKey, username); cmd.Err() != nil {
		return cmd.Err()
	}

	if cmd := db.conn.Del(username); cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func (db *RedisProfileDb) IsThereProfile(username string) bool {
	return db.conn.SIsMember(profileSetKey, username).Val()
}
