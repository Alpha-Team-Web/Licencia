package media

import (
	"back-src/controller/utils/libs"
	"back-src/model/database"
	"back-src/model/existence"
)

func AddUpdateProfileEvent(username string, isFreelancer bool, db *database.Database) error {
	event := existence.Event{
		Username:     username,
		IsFreelancer: isFreelancer,
		EventMessage: libs.Ternary(isFreelancer, existence.FEUpdateProfile, existence.EEUpdateProfile).(string),
		Id:           username,
	}
	return db.MediaTable.AddEvent(event)
}

func AddAddProjectEvent(username, projectId string, db *database.Database) error {
	event := existence.Event{
		Username:     username,
		IsFreelancer: false,
		EventMessage: existence.EEAddProject,
		Id:           projectId,
	}
	return db.MediaTable.AddEvent(event)
}

func AddAssignProjectEvent(username, projectId string, db *database.Database) error {
	event := existence.Event{
		Username:     username,
		IsFreelancer: false,
		EventMessage: existence.EEAddProject,
		Id:           projectId,
	}
	return db.MediaTable.AddEvent(event)
}
