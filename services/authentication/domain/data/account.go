package data

import (
	"github.com/rigolo-api/common/value_objects"
	"github.com/rigolo-api/services/authentication/domain/entities"
)

type NewAccount struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Age       uint8  `json:"age"`
}

type PublicAccount struct {
	ID    value_objects.Identifier `json:"id"`
	Email value_objects.Email                   `json:"email"`
	entities.User
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
