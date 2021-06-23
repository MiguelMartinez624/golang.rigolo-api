package domain

import (
	"github.com/rigolo-api/common/value_objects"
	"strings"
)

type Client struct {
	ID     value_objects.Identifier `json:"id" gorm:"primarykey"`
	Phone  value_objects.Phone      `json:"phone"`
	Email  value_objects.Email      `json:"email"`
	Name   string                   `json:"name"`
	Region string                   `json:"region"`
}

type NewClient struct {
	Name   string              `json:"name"`
	Phone  value_objects.Phone `json:"phone"`
	Email  value_objects.Email `json:"email"`
	Region string              `json:"region"`
}

func (nc *NewClient) IsValid() *ClientError {
	if   !value_objects.IsEmailValid(string(nc.Email)) {
		return NewVerificationLayerError("email")
	}

	if strings.Trim(nc.Name, " ") == "" {
		return NewVerificationLayerError("name")
	}

	if !value_objects.IsPhoneValid(string(nc.Phone)) {
		return NewVerificationLayerError("phone")
	}

	return nil
}
