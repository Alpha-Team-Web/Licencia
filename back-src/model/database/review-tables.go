package database

import "back-src/model/existence"

func (db *Database) AddFreelancerReview(review existence.FreelancerEmployerReview) error {
	_, err := db.db.Model(&review).Insert()
	return err
}

func (db *Database) EditFreelancerReview(review existence.FreelancerEmployerReview) error {
	_, err := db.db.Model(&review).Where("project_id = ?", review.ProjectID).Update()
	return err
}

func (db *Database) AddEmployerReview(review existence.EmployerFreelancerReview) error {
	_, err := db.db.Model(&review).Insert()
	return err
}

func (db *Database) EditEmployerReview(review existence.EmployerFreelancerReview) error {
	_, err := db.db.Model(&review).Where("project_id = ?", review.ProjectID).Update()
	return err
}

func (db *Database) HasEmployerReviewed(projectId string) (bool, error) {
	var reviews []existence.EmployerFreelancerReview
	err := db.db.Model(&reviews).Where("project_id = ?", projectId).Select()
	if err != nil {
		return false, err
	}
	return len(reviews) != 0, nil
}

func (db *Database) HasFreelancerReviewed(projectId string) (bool, error) {
	var reviews []existence.FreelancerEmployerReview
	err := db.db.Model(&reviews).Where("project_id = ?", projectId).Select()
	if err != nil {
		return false, err
	}
	return len(reviews) != 0, nil
}
