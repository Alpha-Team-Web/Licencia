package tables

import (
	"back-src/model/existence"
	"github.com/go-pg/pg"
)

type ReviewTable struct {
	conn *pg.DB
}

func (table *ReviewTable) AddFreelancerReview(review existence.FreelancerEmployerReview) error {
	_, err := table.conn.Model(&review).Insert()
	return err
}

func (table *ReviewTable) EditFreelancerReview(review existence.FreelancerEmployerReview) error {
	_, err := table.conn.Model(&review).Where("project_id = ?", review.ProjectID).Update()
	return err
}

func (table *ReviewTable) AddEmployerReview(review existence.EmployerFreelancerReview) error {
	_, err := table.conn.Model(&review).Insert()
	return err
}

func (table *ReviewTable) EditEmployerReview(review existence.EmployerFreelancerReview) error {
	_, err := table.conn.Model(&review).Where("project_id = ?", review.ProjectID).Update()
	return err
}

func (table *ReviewTable) HasEmployerReviewed(projectId string) (bool, error) {
	var reviews []existence.EmployerFreelancerReview
	err := table.conn.Model(&reviews).Where("project_id = ?", projectId).Select()
	if err != nil {
		return false, err
	}
	return len(reviews) != 0, nil
}

func (table *ReviewTable) HasFreelancerReviewed(projectId string) (bool, error) {
	var reviews []existence.FreelancerEmployerReview
	err := table.conn.Model(&reviews).Where("project_id = ?", projectId).Select()
	if err != nil {
		return false, err
	}
	return len(reviews) != 0, nil
}
