package licnecia_errors

import (
	"errors"
	"strings"
)

func MakeLicenciaError(errorStr string) error {
	return errors.New("Licencia: " + errorStr)
}

func IsLicenciaError(error error) bool {
	return strings.Contains(error.Error(), "Licencia: ")
}

func GetErrorStrForRespond(err error) string {
	return err.Error()[10:len(err.Error())]
}
