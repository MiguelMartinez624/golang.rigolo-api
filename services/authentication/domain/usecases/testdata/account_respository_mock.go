package testdata

import (
	"github.com/rigolo-api/services/authentication/domain/entities"
)

var testUser = entities.BuildUser("Test", "Runner", 99)

type GetAccountByEmailFunc = func(emailStr string) (*entities.Account, error)
type SaveAccountFunc = func(account *entities.Account) (*entities.Account, error)

type AccountRepositoryMock struct {
	GetAccountByEmailFunc GetAccountByEmailFunc
	SaveAccountFunc       SaveAccountFunc
}

func BuildAccountMockRepo(
	getAccountByEmailFunc GetAccountByEmailFunc,
	saveAccountFunc SaveAccountFunc) *AccountRepositoryMock {
	return &AccountRepositoryMock{getAccountByEmailFunc, saveAccountFunc}
}

func BuildAccountRepositoryMock() *AccountRepositoryMock {
	return &AccountRepositoryMock{}
}

func (r *AccountRepositoryMock) GetAccountByEmail(emailStr string) (*entities.Account, error) {
	return r.GetAccountByEmailFunc(emailStr)
}

func (r *AccountRepositoryMock) Save(account *entities.Account) (*entities.Account, error) {
	return r.SaveAccountFunc(account)
}

// Mock functions this are use to mock the repository behavior

func SuccessSaveAccountFunc(account *entities.Account) (*entities.Account, error) {
	return account, nil
}

func SuccessGetAccountByEmailFunc(emailStr string) (*entities.Account, error) {
	return entities.BuildAccount(emailStr, "mock", *testUser)
}

func NoFoundGetAccountByEmailFunc(emailStr string) (*entities.Account, error) {
	return nil, nil
}

func ReturnMockedAccount(mockedAccount  *entities.Account) GetAccountByEmailFunc {
	return func(emailStr string) (*entities.Account, error) {
		return mockedAccount, nil
	}
}
