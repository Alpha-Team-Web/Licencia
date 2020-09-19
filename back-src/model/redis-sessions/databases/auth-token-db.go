package databases

import (
	"back-src/model/existence"
	"back-src/model/redis-sessions/orm"
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
	if stats := authTokenDb.conn.FlushDB(); stats.Err() != nil {
		panic(stats.Err())
	}
	return authTokenDb
}

func (db *RedisAuthTokenDb) MakeNewAuth(username string, isFreelancer bool, initialTime time.Time, token string) (string, error) {
	if cmd := db.conn.SAdd(tokenSetKey, token); cmd.Err() != nil {
		return "", cmd.Err()
	}

	if stats := db.conn.HMSet(
		token,
		orm.HashAuthToken(existence.AuthToken{
			Token:        token,
			Username:     username,
			InitialTime:  initialTime,
			IsFreelancer: isFreelancer,
		},
		),
	); stats.Err() != nil {
		return "", stats.Err()
	}
	return token, nil
}

func (db *RedisAuthTokenDb) IsThereAuthWithToken(token string) (bool, error) {
	return db.conn.SIsMember(tokenSetKey, token).Result()
}

func (db *RedisAuthTokenDb) GetAuthByToken(token string) (existence.AuthToken, error) {
	if values, err := db.conn.HGetAll(token).Result(); err != nil {
		return existence.AuthToken{}, err
	} else {
		return orm.UnHashAuthToken(values), nil
	}
}

func (db *RedisAuthTokenDb) ExpireAuth(token string) error {
	if cmd := db.conn.SRem(tokenSetKey, token); cmd.Err() != nil {
		return cmd.Err()
	}
	if cmd := db.conn.Del(token); cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func (db *RedisAuthTokenDb) GetAllTokens() ([]existence.AuthToken, error) {
	if tokens, err := db.conn.SMembers(tokenSetKey).Result(); err != nil {
		return []existence.AuthToken{}, err
	} else {
		authz := []existence.AuthToken{}
		for _, token := range tokens {
			auth, _ := db.GetAuthByToken(token)
			authz = append(authz, auth)
		}
		return authz, nil
	}
}

func (db *RedisAuthTokenDb) GetUsernameByToken(token string) (string, error) {
	return db.conn.HGet(token, "username").Result()
}
