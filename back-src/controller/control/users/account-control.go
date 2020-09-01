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

func RegisterEmployer(emp existence.Employer, Db *database.Database) error {
	if !Db.EmployerTable.DoesEmployerExistWithUsername(emp.Username) {
		if !Db.EmployerTable.DoesEmployerExistWithEmail(emp.Email) {
			emp.ShownName = emp.Username
			return Db.EmployerTable.InsertEmployer(emp)
		}
		return errors.New("duplicate email: " + emp.Email)
	}
	return errors.New("duplicate username: " + emp.Username)
}

func RegisterFreelancer(frl existence.Freelancer, Db *database.Database) error {
	if !Db.FreelancerTable.DoesFreelancerExistWithUsername(frl.Username) {
		if !Db.FreelancerTable.DoesFreelancerExistWithEmail(frl.Email) {
			frl.ShownName = frl.Username
			frl.AccountType = existence.FreelancerBronze
			return Db.FreelancerTable.InsertFreelancer(frl)
		}
		return errors.New("duplicate email: " + frl.Email)
	}
	return errors.New("duplicate username: " + frl.Username)
}

func Login(loginReq data.LoginRequest, usernameGetter func() (string, error), passwordGetter func(string) (string, error), Db *database.Database) (token string, error error) {
	if username, err := usernameGetter(); err == nil {
		if realPassword, err := passwordGetter(username); err == nil {
			if realPassword == loginReq.Password {
				newToken, err := MakeNewAuthToken(username, loginReq.IsFreelancer, Db)
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

func MakeNewAuthToken(username string, isFreelancer bool, Db *database.Database) (token string, e error) {
	token, err := Db.AuthTokenTable.MakeNewAuth(username, libs.GetRandomString(AuthTokenSize, func(token string) bool {
		if isDuplicate, err := Db.AuthTokenTable.IsThereAuthWithToken(token); err == nil {
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
