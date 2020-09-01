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

func (table *ProjectTable) EditProject(id string, project existence.Project) error {
	if _, err := table.Model(&project).Column("duration", "start_date", "min_budget", "max_budget", "description").Where("id = ?", id).Update(); err != nil {
		return err
	}
	return nil
}

func (table *ProjectTable) GetProject(id string) (existence.Project, error) {
	project := &existence.Project{}
	if err := table.Model(project).Where("id = ?", id).Select(); err != nil {
		return existence.Project{}, err
	} else {
		return *project, nil
	}
}

func (table *ProjectTable) GetOpenProjects() ([]existence.Project, error) {
	projects := &[]existence.Project{}
	if err := table.Model(projects).Where("project_status = ?", existence.Open).Select(); err != nil {
		return nil, err
	}
	return *projects, nil
}
