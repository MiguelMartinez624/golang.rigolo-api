package domain

import "github.com/rigolo-api/common"

type Client struct {
	ID     common.Identifier `json:"id" gorm:"primarykey"`
	Phone  common.Phone      `json:"phone"`
	Email  common.Email      `json:"email"`
	Name   string            `json:"name"`
	Region string            `json:"region"`
}

type NewClient struct {
	Name   string `json:"name"`
	Phone  common.Phone `json:"phone"`
	Email  common.Email `json:"email"`
	Region string `json:"region"`
}

func (nc NewClient) IsValid() *ClientError {
	return nil
}
