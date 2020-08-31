package database

import "back-src/model/existence"

func (db *Database) GetFreelancerUsernameByProjectId(projectId string) (string, error) {
	project := existence.Project{}
	if err := db.db.Model(&project).Where("id = ?", projectId).Column("freelancer_username").Select(); err != nil {
		return "", err
	}
	return project.FreelancerUsername, nil
}

func (db *Database) GetEmployerUsernameByProjectId(projectId string) (string, error) {
	project := existence.Project{}
	if err := db.db.Model(&project).Where("id = ?", projectId).Column("employer_username").Select(); err != nil {
		return "", err
	}
	return project.EmployerUsername, nil
}
