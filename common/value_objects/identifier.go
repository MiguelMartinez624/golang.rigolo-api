package value_objects

import "github.com/google/uuid"

type Identifier string

func (i Identifier) IsEmpty() bool {
	return string(i) == ""
}

func NewID() Identifier {
	return Identifier(uuid.New().String())
}
