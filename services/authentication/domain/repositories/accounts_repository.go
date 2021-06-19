package repositories

import "github.com/rigolo-api/services/authentication/domain/entities"

type Accounts interface {
	GetAccountByEmail(emailStr string) (*entities.Account, error)

	Save(account *entities.Account) (*entities.Account, error)
}
