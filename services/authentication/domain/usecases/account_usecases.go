package usecases

import (
	"github.com/rigolo-api/services/authentication/domain/data"
	"github.com/rigolo-api/services/authentication/domain/entities"
	"github.com/rigolo-api/services/authentication/domain/errors"
	"github.com/rigolo-api/services/authentication/domain/repositories"
	"log"
)

type AccountUseCases struct {
	accountsRepo repositories.Accounts
}

func BuildAccountService(accountsRepo repositories.Accounts) *AccountUseCases {
	return &AccountUseCases{accountsRepo}
}

func (s *AccountUseCases) AuthenticateAccount(credentials data.Credentials) (*data.PublicAccount, error) {
	account, err := s.accountsRepo.GetAccountByEmail(credentials.Email)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, errors.AccountNotFound
	}

	if match := account.Password.Equals(credentials.Password); !match {
		return nil, errors.InvalidCredentials
	}

	result := &data.PublicAccount{
		ID:    account.ID,
		Email: account.Email,
		User:  account.Profile,
	}

	return result, nil

}

func (s *AccountUseCases) RegisterAccount(newAccount data.NewAccount) error {
	log.Printf("[BuildRegisterAccount] input %+v \n", newAccount)
	exist, err := s.accountsRepo.GetAccountByEmail(string(newAccount.Email))
	if err != nil {
		return err
	}
	log.Printf("[BuildRegisterAccount] %+v \n", exist)
	if exist != nil {
		return errors.EmailAlreadyUsed
	}

	user := entities.BuildUser(newAccount.FirstName, newAccount.LastName, newAccount.Age)
	account, err := entities.BuildAccount(newAccount.Email, newAccount.Password, *user)
	if err != nil {
		return err
	}

	_, err = s.accountsRepo.Save(account)
	return err

}
