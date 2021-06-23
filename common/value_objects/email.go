package value_objects

import (
	"errors"
	"regexp"
	"strings"
)

type Email string

func NewEmail(raw string) (Email, error) {
	if !IsEmailValid(raw) {
		return "", errors.New("invalid email format")
	}
	return Email(raw), nil
}

// isEmailValid checks if the email provided is valid by regex.
func IsEmailValid(e string) bool {
	if strings.Trim(e, " ") == ""{
		return false
	}

	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}
