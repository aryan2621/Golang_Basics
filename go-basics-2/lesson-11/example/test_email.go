package example

import (
	"errors"
	"regexp"
)

func IsEmailValid(email string) (string, error) {
	r, _ := regexp.Compile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)
	if r.MatchString(email) {
		return email, nil
	} else {
		return "", errors.New("email is not valid")
	}
}
