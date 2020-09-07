package media

import (
	"back-src/model/database"
	"back-src/model/existence"
)

func Follow(token string, follow existence.Follow, db *database.Database) error {
	if auth, err := db.AuthTokenTable.GetAuthByToken(token); err == nil {
		if isThere, _ := db.MediaTable.IsThereFollow(auth.Username, follow.FollowingUsername); !isThere {
			return modifyFollow(auth, follow, db.MediaTable.AddFollow)
		}
	} else {
		return err
	}
	return nil
}

func UnFollow(token string, follow existence.Follow, db *database.Database) error {
	if auth, err := db.AuthTokenTable.GetAuthByToken(token); err == nil {
		return modifyFollow(auth, follow, db.MediaTable.RemoveFollow)
	} else {
		return err
	}
}

func modifyFollow(auth existence.AuthToken, follow existence.Follow, modifyFollow func(existence.Follow) error) error {
	follow.FollowerUsername = auth.Username
	follow.FollowerFreelancer = auth.IsFreelancer
	return modifyFollow(follow)
}
