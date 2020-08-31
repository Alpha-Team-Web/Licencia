package projects

import (
	"back-src/model/database"
	"back-src/model/existence"
	"errors"
)

func AddFreelancerReview(token string, review existence.FreelancerEmployerReview, db *database.Database) error {
	if username, err := db.GetUsernameByToken(token); err != nil {
		return err
	} else {
		if realUsername, err := db.GetFreelancerUsernameByProjectId(review.ProjectID); err != nil {
			return err
		} else if realUsername == username {
			if hasReviewed, err := db.HasFreelancerReviewed(review.ProjectID); err != nil {
				return err
			} else if hasReviewed {
				return db.EditFreelancerReview(review)
			} else {
				return db.AddFreelancerReview(review)
			}
		} else {
			return errors.New("not involved in project the username: " + username)
		}
	}
}
