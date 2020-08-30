package users

import (
	"back-src/controller/control/utils/data"
	"back-src/controller/control/utils/libs"
	"back-src/model/database"
	"back-src/model/existence"
	"errors"
)

const (
	AuthTokenSize = 20
)

func RegisterEmployer(emp existence.Employer, Db *database.Database) error {
	if !Db.DoesEmployerExistWithUsername(emp.Username) {
		if !Db.DoesEmployerExistWithEmail(emp.Email) {
			return Db.InsertEmployer(emp)
		}
		return errors.New("duplicate email: " + emp.Email)
	}
	return errors.New("duplicate username: " + emp.Username)
}

func RegisterFreelancer(frl existence.Freelancer, Db *database.Database) error {
	if !Db.DoesFreelancerExistWithUsername(frl.Username) {
		if !Db.DoesEmployerExistWithEmail(frl.Email) {
			return Db.InsertFreelancer(frl)
		}
		return errors.New("duplicate email: " + frl.Email)
	}
	return errors.New("duplicate username: " + frl.Username)
}

func Login(loginReq data.LoginRequest, usernameGetter func() (string, error), passwordGetter func(string) (string, error), Db *database.Database) (token string, error error) {
	if username, err := usernameGetter(); err == nil {
		if realPassword, err := passwordGetter(username); err == nil {
			if realPassword == loginReq.Password {
				newToken, err := Db.MakeNewAuth(username, libs.GetRandomString(AuthTokenSize, func(token string) bool {
					if isDuplicate, err := Db.IsThereAuthWithToken(token); err == nil {
						return isDuplicate
					} else {
						error = err
					}
					return false
				}), loginReq.IsFreelancer)
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
