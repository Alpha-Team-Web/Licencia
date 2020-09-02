package users

import (
	"back-src/controller/utils/data"
	"back-src/controller/utils/libs"
	"back-src/model/database"
	"back-src/model/existence"
	"errors"
)

const (
	AuthTokenSize = 20
)

func RegisterEmployer(emp existence.Employer, db *database.Database) error {
	if !db.EmployerTable.DoesEmployerExistWithUsername(emp.Username) {
		if !db.EmployerTable.DoesEmployerExistWithEmail(emp.Email) {
			emp.ShownName = emp.Username
			return db.EmployerTable.InsertEmployer(emp)
		}
		return errors.New("duplicate email: " + emp.Email)
	}
	return errors.New("duplicate username: " + emp.Username)
}

func RegisterFreelancer(frl existence.Freelancer, db *database.Database) error {
	if !db.FreelancerTable.DoesFreelancerExistWithUsername(frl.Username) {
		if !db.FreelancerTable.DoesFreelancerExistWithEmail(frl.Email) {
			frl.ShownName = frl.Username
			frl.AccountType = existence.FreelancerBronze
			return db.FreelancerTable.InsertFreelancer(frl)
		}
		return errors.New("duplicate email: " + frl.Email)
	}
	return errors.New("duplicate username: " + frl.Username)
}

func Login(loginReq data.LoginRequest, db *database.Database) (token string, error error) {
	if username, err := getUsernameGetter(loginReq.Id, loginReq.IsFreelancer, db)(); err == nil {
		if realPassword, err := getPasswordGetter(loginReq.IsFreelancer, db)(username); err == nil {
			if realPassword == loginReq.Password {
				newToken, err := MakeNewAuthToken(username, loginReq.IsFreelancer, db)
				if err == nil {
					token = newToken
					error = nil
				} else {
					error = err
				}
			} else {
				error = errors.New("invalid password: " + loginReq.Password)
			}
		} else {
			error = err
		}
	} else {
		error = err
	}
	return
}

func getUsernameGetter(Id string, isFreelancer bool, db *database.Database) func() (username string, error error) {
	if isFreelancer {
		return getUsernameById(Id, db.FreelancerTable.DoesFreelancerExistWithEmail, db.FreelancerTable.DoesFreelancerExistWithUsername, db.FreelancerTable.GetFreelancerUsernameByEmail)
	} else {
		return getUsernameById(Id, db.EmployerTable.DoesEmployerExistWithEmail, db.EmployerTable.DoesEmployerExistWithUsername, db.EmployerTable.GetEmployerUsernameByEmail)
	}
}

type doesExist func(string) bool
type getUsernameByEmail func(string) (string, error)

func getUsernameById(Id string, doesUserExistWithEmail doesExist, doesUserExistWithUsername doesExist, getUsername getUsernameByEmail) func() (string, error) {
	var username string
	var e error
	if libs.IsEmailValid(Id) {
		if doesUserExistWithEmail(Id) {
			if user, err := getUsername(Id); err == nil {
				username = user
			} else {
				e = err
			}
		} else {
			e = errors.New("not signed up email: " + Id)
		}
	} else {
		if doesUserExistWithUsername(Id) {
			username = Id
		} else {
			e = errors.New("not signed up username: " + Id)
		}
	}
	return func() (string, error) {
		return username, e
	}
}

func getPasswordGetter(isFreelancer bool, db *database.Database) func(string) (string, error) {
	if isFreelancer {
		return db.FreelancerTable.GetFreelancerPasswordByUsername
	} else {
		return db.EmployerTable.GetEmployerPasswordByUsername
	}
}

func MakeNewAuthToken(username string, isFreelancer bool, db *database.Database) (token string, e error) {
	token, err := db.AuthTokenTable.MakeNewAuth(username, libs.GetRandomString(AuthTokenSize, func(token string) bool {
		if isDuplicate, err := db.AuthTokenTable.IsThereAuthWithToken(token); err == nil {
			return isDuplicate
		} else {
			e = err
		}
		return false
	}), isFreelancer)
	if err != nil {
		e = err
	}
	return
}
