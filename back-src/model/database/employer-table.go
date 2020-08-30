package database

import (
	"back-src/model/existence"
)

func (db Database) DoesEmployerExistWithUsername(username string) bool {
	resultSet := &[]existence.Employer{}
	_ = db.db.Model(resultSet).Where("username = ?", username).Select()
	return len(*resultSet) != 0
}

func (db *Database) DoesEmployerExistWithEmail(email string) bool {
	resultSet := &[]existence.Employer{}
	_ = db.db.Model(resultSet).Where("email = ?", email).Select()
	return len(*resultSet) != 0
}

func (db *Database) InsertEmployer(emp existence.Employer) error {
	_, err := db.db.Model(&emp).Insert()
	return err
}

func (db *Database) UpdateEmployer(username string, emp existence.Employer) error {

	if _, err := db.db.Model(&emp).Column("password", "email", "description", "first_name", "last_name", "phone_number", "address").Where("username = ?", username).Update(); err != nil {
		return err
	}
	return nil
}

func (db *Database) GetEmployer(username string) (existence.Employer, error) {
	emp := new(existence.Employer)
	err := db.db.Model(emp).Where("username = ?", username).Select()
	return *emp, err
}

func (db *Database) GetEmployerProjects(username string) ([]existence.Project, error) {
	emp := new(existence.Employer)
	err := db.db.Model(emp).Where("username = ?", username).Select()
	if err != nil {
		return nil, err
	}

	projectIDs := emp.ProjectIds
	projects := make([]existence.Project, len(projectIDs))
	for i := range projectIDs {
		project := new(existence.Project)
		db.db.Model(project).Where("id = ?", i).Select()
		projects = append(projects, *project)
	}
	return projects, nil
}

func (db *Database) GetEmployerPasswordByUsername(username string) (string, error) {
	employer := existence.Employer{}
	err := db.db.Model(&employer).Where("username = ?", username).Column("password").Select()
	return employer.Password, err
}

func (db *Database) GetEmployerPasswordByEmail(email string) (string, error) {
	employer := existence.Employer{}
	err := db.db.Model(&employer).Where("email = ?", email).Column("password").Select()
	return employer.Password, err
}
