package persistency

import (
	"github.com/rigolo-api/services/authentication/domain/entities"
	"gorm.io/gorm"
)

type AccountRepositoryPostgres struct {
	connection *gorm.DB
}

func BuildAccountRepositoryPostgres(connection *gorm.DB) *AccountRepositoryPostgres {
	return &AccountRepositoryPostgres{connection}
}

func (a AccountRepositoryPostgres) GetAccountByEmail(emailStr string) (*entities.Account, error) {
	var entity entities.Account
	result := a.connection.Where("email = ?", emailStr).Find(&entity)
	if entity.ID.IsEmpty() {
		return nil, result.Error
	}

	return &entity, result.Error

}

func (a AccountRepositoryPostgres) Save(account *entities.Account) (*entities.Account, error) {
	result := a.connection.Create(account)
	return account, result.Error
}
