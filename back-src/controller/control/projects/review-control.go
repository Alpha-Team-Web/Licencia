package projects

import (
	"back-src/model/database"
	"back-src/model/existence"
	"errors"
)

func AddFreelancerReview(token string, review existence.FreelancerEmployerReview, db *database.Database) error {
	if username, err := db.AuthTokenTable.GetUsernameByToken(token); err != nil {
		return err
	} else {
		if realUsername, err := db.ProjectTable.GetFreelancerUsernameByProjectId(review.ProjectID); err != nil {
			return err
		} else if realUsername == username {
			if hasReviewed, err := db.ReviewTable.HasFreelancerReviewed(review.ProjectID); err != nil {
				return err
			} else if hasReviewed {
				return db.ReviewTable.EditFreelancerReview(review)
			} else {
				return db.ReviewTable.AddFreelancerReview(review)
			}
		} else {
			return errors.New("not involved in project the username: " + username)
		}
	}
}
