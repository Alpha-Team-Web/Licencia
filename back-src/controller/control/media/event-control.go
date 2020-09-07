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

func AddAssignProjectEvent(employerUsername, freelancerUsername, projectId string, db *database.Database) error {
	event := existence.Event{
		Username:     employerUsername,
		IsFreelancer: false,
		EventMessage: existence.EEAddProject,
		Id:           projectId,
	}
	err := db.MediaTable.AddEvent(event)
	event2 := existence.Event{
		Username:     freelancerUsername,
		IsFreelancer: true,
		EventMessage: existence.FETakeProject,
		Id:           projectId,
	}
	if err == nil {
		return db.MediaTable.AddEvent(event2)
	} else {
		db.MediaTable.AddEvent(event2)
		return err
	}
}

func AddExtendProjectEvent(username, projectId string, db *database.Database) error {
	event := existence.Event{
		Username:     username,
		IsFreelancer: false,
		EventMessage: existence.EEExtendProject,
		Id:           projectId,
	}
	return db.MediaTable.AddEvent(event)
}
func AddCloseProjectEvent(employerUsername, freelancerUsername, projectId string, db *database.Database) error {
	event := existence.Event{
		Username:     employerUsername,
		IsFreelancer: false,
		EventMessage: existence.EECloseProject,
		Id:           projectId,
	}
	err := db.MediaTable.AddEvent(event)
	event2 := existence.Event{
		Username:     freelancerUsername,
		IsFreelancer: true,
		EventMessage: existence.FEDeliverProject,
		Id:           projectId,
	}
	if err == nil {
		return db.MediaTable.AddEvent(event2)
	} else {
		db.MediaTable.AddEvent(event2)
		return err
	}
}

func AddFollowEvent(username string, id string, eventMessage string, isFreelancer bool, db *database.Database) error {
	event := existence.Event{
		Username:     username,
		IsFreelancer: isFreelancer,
		EventMessage: eventMessage,
		Id:           id,
	}
	return db.MediaTable.AddEvent(event)
}

func AddRequestEvent(username string, id string, db *database.Database) error {
	event := existence.Event{
		Username:     username,
		IsFreelancer: true,
		EventMessage: existence.FEReqProject,
		Id:           id,
	}
	return db.MediaTable.AddEvent(event)
}
