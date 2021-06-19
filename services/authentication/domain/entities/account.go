package entities

import (
	"github.com/rigolo-api/common/value_objects"
)

type Account struct {
	ID       value_objects.Identifier     `json:"id" gorm:"primarykey""`
	Email    value_objects.Email          `json:"email"`
	Password value_objects.SecurePassword `json:"password"`
	Profile  User                         `gorm:"embedded"`

}

// Getters

func BuildAccount(email string, password string, user User) (*Account, error) {
	securedPassword, err := value_objects.BuildPassword(password)
	if err != nil {
		return nil, err
	}

	validEmail, err := value_objects.NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &Account{
		ID:       value_objects.NewID(),
		Password: securedPassword,
		Email:    validEmail,
		Profile:  user,
	}, nil
}
