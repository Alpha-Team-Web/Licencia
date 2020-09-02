package tables

import (
	"back-src/model/existence"
	"github.com/go-pg/pg"
)

type ProjectTable struct {
	conn *pg.DB
}

func NewProjectTable(db *pg.DB) *ProjectTable {
	return &ProjectTable{db}
}

func (table *ProjectTable) GetFreelancerUsernameByProjectId(projectId string) (string, error) {
	project := existence.Project{}
	if err := table.conn.Model(&project).Where("id = ?", projectId).Column("freelancer_username").Select(); err != nil {
		return "", err
	}
	return project.FreelancerUsername, nil
}

func (table *ProjectTable) GetEmployerUsernameByProjectId(projectId string) (string, error) {
	project := existence.Project{}
	if err := table.conn.Model(&project).Where("id = ?", projectId).Column("employer_username").Select(); err != nil {
		return "", err
	}
	return project.EmployerUsername, nil
}

func (table *ProjectTable) AddProject(project existence.Project) error {
	if _, err := table.conn.Model(&project).Insert(); err != nil {
		return err
	}
	return nil
}

func (table *ProjectTable) EditProject(id string, project existence.Project) error {
	if _, err := table.conn.Model(&project).Column("duration", "start_date", "min_budget", "max_budget", "description").Where("id = ?", id).Update(); err != nil {
		return err
	}
	return nil
}

func (table *ProjectTable) GetProject(id string) (existence.Project, error) {
	project := &existence.Project{}
	if err := table.conn.Model(project).Where("id = ?", id).Select(); err != nil {
		return existence.Project{}, err
	} else {
		return *project, nil
	}
}

func (table *ProjectTable) GetProjectsByStatus(status string) ([]existence.Project, error) {
	projects := &[]existence.Project{}
	if err := table.conn.Model(projects).Where("project_status = ?", status).Select(); err != nil {
		return nil, err
	}
	return *projects, nil
}

func (table *ProjectTable) GetProjectIdsByStatusAndMaxBudget(status string, max, min float64) ([]string, error) {
	var projects []existence.Project
	query := table.conn.Model(&projects)
	query = query.Where("project_status = ?", status).Where("max_budget > ?", (min-0.001)).Where("max_budget < ?", (max + 0.001))
	query = query.Column("id")
	if err := query.Select(); err != nil {
		return nil, err
	}
	var ids []string
	for _, project := range projects {
		ids = append(ids, project.Id)
	}
	return ids, nil
}

func (table *ProjectTable) GetProjectDefinedColumns(id string, columns ...string) (existence.Project, error) {
	project := existence.Project{}
	if err := table.conn.Model(&project).Where("id = ?", id).Column(columns...).Select(); err != nil {
		return existence.Project{}, err
	}
	return project, nil
}

func (table *ProjectTable) IsThereProjectWithId(projectId string) (bool, error) {
	var resultSet []existence.Project
	error := table.conn.Model(&resultSet).Where("id = ?", projectId).Select()
	return len(resultSet) != 0, error
}

func (table *ProjectTable) GetProjectStatus(projectId string) (string, error) {
	project := existence.Project{}
	if err := table.conn.Model(&project).Where("id = ?", projectId).Column("project_status").Select(); err != nil {
		return "", err
	}
	return project.ProjectStatus, nil
}

func (table *ProjectTable) GetProjectRequests(projectId string) (map[string]string, error) {
	project := existence.Project{}
	if err := table.conn.Model(&project).Where("id = ?", projectId).Column("freelancer_requests_with_description").Select(); err != nil {
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
		if _, err := table.conn.Model(&project).Where("id = ?", projectId).Column("freelancer_requests_with_description").Update(); err != nil {
			return err
		}
		return nil
	}
}

func (table *ProjectTable) AddFreelancerToProject(username string, projectId string) error {
	project := existence.Project{FreelancerUsername: username}
	if _, err := table.conn.Model(&project).Column("freelancer_username").Where("id = ?", projectId).Update(); err != nil {
		return err
	}
	return nil
}

func (table *ProjectTable) SetProjectStatus(id string, status string) error {
	project, err := table.GetProject(id)
	if err != nil {
		return err
	}
	project.ProjectStatus = status
	if _, err := table.conn.Model(&project).Column("project_status").Where("id = ?", id).Update(); err != nil {
		return err
	}
	return nil
}

func (table *ProjectTable) DeleteProjectDescriptions(id string) error {
	project := existence.Project{FreelancerRequestsWithDescription: map[string]string{}}
	if _, err := table.conn.Model(&project).Column("freelancer_requests_with_description").Where("id = ?", id).Update(); err != nil {
		return err
	}
	return nil
}
