package value_objects

import (
	"strings"
)

type Phone string

// IsPhoneValid checks if the phone provided is valid by regex.
func IsPhoneValid(p string) bool {
	if strings.Trim(p, " ") == "" {
		return false
	}

	if len(p) != 11 {
		return false
	}

	return true
}
