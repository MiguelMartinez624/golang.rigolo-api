package common

import "github.com/google/uuid"

type Identifier string

func  NewIdentifier() Identifier {
	return Identifier(uuid.New().String())
}


type Email string

type Phone string
