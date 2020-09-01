package tables

import (
	"back-src/model/existence"
	"github.com/go-pg/pg"
)

type ProjectTable struct {
	*pg.DB
}

func (table *ProjectTable) GetFreelancerUsernameByProjectId(projectId string) (string, error) {
	project := existence.Project{}
	if err := table.Model(&project).Where("id = ?", projectId).Column("freelancer_username").Select(); err != nil {
		return "", err
	}
	return project.FreelancerUsername, nil
}

func (table *ProjectTable) GetEmployerUsernameByProjectId(projectId string) (string, error) {
	project := existence.Project{}
	if err := table.Model(&project).Where("id = ?", projectId).Column("employer_username").Select(); err != nil {
		return "", err
	}
	return project.EmployerUsername, nil
}

func (table *ProjectTable) AddProject(project existence.Project) error {
	if _, err := table.Model(&project).Insert(); err != nil {
		return err
	}
	return nil
}

func (table *ProjectTable) IsThereProjectWithId(projectId string) (bool, error) {
	var resultSet []existence.Project
	error := table.Model(&resultSet).Where("id = ?", projectId).Select()
	return len(resultSet) != 0, error
}

func (table *ProjectTable) GetProjectStatus(projectId string) (string, error) {
	project := existence.Project{}
	if err := table.Model(&project).Where("id = ?", projectId).Column("project_status").Select(); err != nil {
		return "", err
	}
	return project.ProjectStatus, nil
}

func (table *ProjectTable) GetProjectRequests(projectId string) (map[string]string, error) {
	project := existence.Project{}
	if err := table.Model(&project).Where("id = ?", projectId).Column("freelancer_requests_with_description").Select(); err != nil {
		return map[string]string{}, err
	}
	return project.FreelancerRequestsWithDescription, nil
}

func (table *ProjectTable) AddRequestToProject(projectId string, username string, description string) error {
	project := existence.Project{}
	if projectRequests, err := table.GetProjectRequests(projectId); err != nil {
		return err
	} else {
		project.FreelancerRequestsWithDescription = projectRequests
		project.FreelancerRequestsWithDescription[username] = description
		if _, err := table.Model(&project).Where("id = ?", projectId).Column("freelancer_requests_with_description").Update(); err != nil {
			return err
		}
		return nil
	}
}
