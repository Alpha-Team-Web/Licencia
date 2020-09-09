package tables

import (
	"back-src/model/existence"
	"github.com/go-pg/pg"
	"time"
)

type EmployerTable struct {
	conn *pg.DB
}

func NewEmployerTable(db *pg.DB) *EmployerTable {
	return &EmployerTable{db}
}

func (table *EmployerTable) DoesEmployerExistWithUsername(username string) bool {
	resultSet := &[]existence.Employer{}
	_ = table.conn.Model(resultSet).Where("username = ?", username).Select()
	return len(*resultSet) != 0
}

func (table *EmployerTable) GetEmployerShownNameByUsername(username string) (string, error) {
	emp := existence.Employer{}
	err := table.conn.Model(&emp).Where("username = ?", username).Column("shown_name").Select()
	if err != nil {
		return "", err
	}
	return emp.ShownName, nil
}

func (table *EmployerTable) DoesEmployerExistWithEmail(email string) bool {
	resultSet := &[]existence.Employer{}
	_ = table.conn.Model(resultSet).Where("email = ?", email).Select()
	return len(*resultSet) != 0
}

func (table *EmployerTable) InsertEmployer(emp existence.Employer) error {
	_, err := table.conn.Model(&emp).Insert()
	return err
}

func (table *EmployerTable) UpdateEmployerProfile(username string, emp existence.Employer) error {
	if _, err := table.conn.Model(&emp).Column("shown_name", "description", "first_name", "last_name", "phone_number", "address").Where("username = ?", username).Update(); err != nil {
		return err
	}
	return nil
}

func (table *EmployerTable) UpdateEmployerPassword(username string, oldPass string, newPass string) error {
	emp := existence.Employer{}
	emp.Password = newPass
	if _, err := table.conn.Model(&emp).Column("password").Where("username = ?", username).Update(); err != nil {
		return err
	}
	return nil
}

func (table *EmployerTable) UpdateEmployerProjects(username string, emp existence.Employer) error {
	if _, err := table.conn.Model(&emp).Column("project_ids").Where("username = ?", username).Update(); err != nil {
		return err
	}
	return nil
}

func (table *EmployerTable) GetEmployer(username string) (existence.Employer, error) {
	emp := new(existence.Employer)
	err := table.conn.Model(emp).Where("username = ?", username).Select()
	return *emp, err
}

func (table *EmployerTable) GetEmployerProjects(username string) ([]existence.Project, error) {
	emp := new(existence.Employer)
	err := table.conn.Model(emp).Where("username = ?", username).Select()
	if err != nil {
		return nil, err
	}

	projectIDs := emp.ProjectIds
	projects := make([]existence.Project, len(projectIDs))
	for i := range projectIDs {
		project := new(existence.Project)
		table.conn.Model(project).Where("id = ?", i).Select()
		projects = append(projects, *project)
	}
	return projects, nil
}

func (table *EmployerTable) GetEmployerPasswordByUsername(username string) (string, error) {
	employer := existence.Employer{}
	err := table.conn.Model(&employer).Where("username = ?", username).Column("password").Select()
	return employer.Password, err
}

func (table *EmployerTable) GetEmployerUsernameByEmail(email string) (string, error) {
	employer := existence.Employer{}
	err := table.conn.Model(&employer).Where("email = ?", email).Column("username").Select()
	return employer.Username, err
}

func (table *EmployerTable) ExtendProject(id string, date time.Time) error {
	project := existence.Project{
		FinishDate: date,
	}
	_, err := table.conn.Model(&project).Column("finish_date").Where("id = ?", id).Update()
	return err
}
