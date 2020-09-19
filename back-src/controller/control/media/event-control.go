package media

import (
	"back-src/controller/utils/libs"
	"back-src/model/existence"
	"back-src/model/sql"
	"time"
)

func AddUpdateProfileEvent(username string, isFreelancer bool, db *sql.Database) error {
	return addNewEvent(username, isFreelancer, libs.Ternary(isFreelancer, existence.FEUpdateProfile, existence.EEUpdateProfile).(string), username, db)
}

func AddAddProjectEvent(username, projectId string, db *sql.Database) error {
	return addNewEvent(username, false, existence.EEAddProject, projectId, db)
}

func AddAssignProjectEvent(employerUsername, freelancerUsername, projectId string, db *sql.Database) error {
	err := addNewEvent(employerUsername, false, existence.EEAssignProject, projectId, db)
	job := func() error {
		return addNewEvent(freelancerUsername, true, existence.FETakeProject, projectId, db)
	}
	if err == nil {
		return job()
	} else {
		job()
		return err
	}
}

func AddExtendProjectEvent(username, projectId string, db *sql.Database) error {
	return addNewEvent(username, false, existence.EEExtendProject, projectId, db)
}
func AddCloseProjectEvent(employerUsername, freelancerUsername, projectId string, db *sql.Database) error {
	err := addNewEvent(employerUsername, false, existence.EECloseProject, projectId, db)
	job := func() error {
		return addNewEvent(freelancerUsername, true, existence.FEDeliverProject, projectId, db)
	}
	if err == nil {
		return job()
	} else {
		job()
		return err
	}
}

func AddFollowEvent(username string, id string, eventMessage string, isFreelancer bool, db *sql.Database) error {
	return addNewEvent(username, isFreelancer, eventMessage, id, db)
}

func AddRequestEvent(username string, id string, db *sql.Database) error {
	return addNewEvent(username, true, existence.FEReqProject, id, db)
}

func addNewEvent(username string, isFreelancer bool, eventMessage string, id string, db *sql.Database) error {
	event := existence.Event{
		Username:     username,
		IsFreelancer: isFreelancer,
		EventMessage: eventMessage,
		Id:           id,
		Time:         time.Now(),
	}
	return db.MediaTable.AddEvent(event)
}
