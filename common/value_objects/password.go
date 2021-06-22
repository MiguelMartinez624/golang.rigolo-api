package value_objects

import "fmt"

type SecurePassword string

func BuildPassword(raw string) (SecurePassword, error) {
	return SecurePassword(fmt.Sprintf("%v-secured", raw)), nil
}

func (sp SecurePassword) Equals(strToCompare string) bool {
	return fmt.Sprintf("%v-secured", strToCompare) == string(sp)
}
