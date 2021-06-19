package entities

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       uint8   `json:"age"`
}

func BuildUser(firstName, lastName string, age uint8) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}
