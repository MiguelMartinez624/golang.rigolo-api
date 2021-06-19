package tests

import (
	"github.com/rigolo-api/services/authentication/domain/data"
	"github.com/rigolo-api/services/authentication/domain/entities"
	"github.com/rigolo-api/services/authentication/domain/repositories"
	"github.com/rigolo-api/services/authentication/domain/usecases"
	"github.com/rigolo-api/services/authentication/domain/usecases/testdata"
	"reflect"
	"testing"
)

type authenticateAccountTestCase struct {
	name        string
	accountRepo repositories.Accounts
	input       data.Credentials
	wantErr     bool
	expErr      error
}

var buildMockAccount = func() *entities.Account {
	acc, err := entities.BuildAccount("test@test.com", "test_password", entities.User{})
	if err != nil {
		panic(err)
	}
	return acc
}

var successAuthenticateCase = authenticateAccountTestCase{
	name:    "it should successfully authenticate with correct credentials",
	wantErr: false,
	input:   data.Credentials{Email: "test@test.com", Password: "test_password"},
	accountRepo: testdata.BuildAccountMockRepo(
		testdata.ReturnMockedAccount(buildMockAccount()),
		testdata.SuccessSaveAccountFunc,
	),
}

func TestAuthenticateAccount_Execute(t *testing.T) {

	tests := []authenticateAccountTestCase{
		successAuthenticateCase,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := usecases.BuildAccountService(tt.accountRepo)

			var err error
			if _, err = service.AuthenticateAccount(tt.input); (err != nil) != tt.wantErr {
				t.Errorf("AccountUseCases() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				if !reflect.DeepEqual(tt.expErr, err) {
					t.Errorf("AccountUseCases() error = %v, expErr %v", err, tt.expErr)
				}
			}
		})
	}
}
