package databases

import (
	"back-src/model/existence"
	"time"
)

//DB Number 4

const authTokenDbId = 4

type RedisAuthTokenDb struct {
	conn *RedisCli
}

const (
	tokenSetKey = "tokens"
)

func NewRedisAuthTokenDB(addr, password string) *RedisAuthTokenDb {
	authTokenDb := &RedisAuthTokenDb{
		conn: NewRedisCli(addr, password, authTokenDbId),
	}
	//TODO(AT THE END)
	/*	if stats := authTokenDb.conn.FlushDB(); stats.Err() != nil {
		panic(stats.Err())
	}*/
	return authTokenDb
}

func (db *RedisAuthTokenDb) MakeNewAuth(username string, isFreelancer bool, initialTime time.Time, token string) (string, error) {
	if stats := db.conn.Set(tokenSetKey, token, 0); stats.Err() != nil {
		return "", stats.Err()
	}
	if stats := db.conn.HMSet(token, map[string]interface{}{
		"username":     username,
		"isFreelancer": isFreelancer,
		"initialTime":  initialTime,
	}); stats.Err() != nil {
		return "", stats.Err()
	}
	return token, nil
}

func (db *RedisAuthTokenDb) IsThereAuthWithToken(token string) (bool, error) {
	return db.conn.SIsMember(tokenSetKey, token).Result()
}

func (db *RedisAuthTokenDb) GetAuthByToken(token string) (existence.AuthToken, error) {
	if values, err := db.conn.HMGet(token, "username", "isFreelancer", "initialTime").Result(); err != nil {
		return existence.AuthToken{}, err
	} else {
		return existence.AuthToken{
			Token:        token,
			Username:     values[0].(string),
			IsFreelancer: values[1].(bool),
			InitialTime:  values[2].(time.Time),
		}, nil
	}
}

func (db *RedisAuthTokenDb) ExpireAuth(token string) error {
	if cmd := db.conn.SRem(tokenSetKey, token); cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}
