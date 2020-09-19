package media

import (
	"back-src/controller/utils/libs"
	"back-src/model/existence"
	"back-src/model/sql"
)

func Follow(token string, follow existence.Follow, db *sql.Database) error {
	if auth, err := db.AuthTokenTable.GetAuthByToken(token); err == nil {
		if isThere, _ := db.MediaTable.IsThereFollow(auth.Username, follow.FollowingUsername); !isThere {
			return modifyFollow(auth, follow, db.MediaTable.AddFollow, db)
		}
	} else {
		return err
	}
	return nil
}

func UnFollow(token string, follow existence.Follow, db *sql.Database) error {
	if auth, err := db.AuthTokenTable.GetAuthByToken(token); err == nil {
		return modifyFollow(auth, follow, db.MediaTable.RemoveFollow, db)
	} else {
		return err
	}
}

func modifyFollow(auth existence.AuthToken, follow existence.Follow, modifyFollow func(existence.Follow) error, db *sql.Database) error {
	follow.FollowerUsername = auth.Username
	follow.FollowerFreelancer = auth.IsFreelancer
	AddFollowEvent(follow.FollowerUsername, follow.FollowingUsername, libs.Ternary(follow.FollowerFreelancer, libs.Ternary(follow.FollowingFreelancer, existence.FEFollowFreelancer, existence.FEFollowEmployer).(string), libs.Ternary(follow.FollowingFreelancer, existence.EEFollowFreelancer, existence.EEFollowEmployer).(string)).(string), follow.FollowerFreelancer, db)
	return modifyFollow(follow)
}
