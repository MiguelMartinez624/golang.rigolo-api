package tests

import (
	"github.com/rigolo-api/services/authentication/domain/data"
	"github.com/rigolo-api/services/authentication/domain/errors"
	"github.com/rigolo-api/services/authentication/domain/repositories"
	"github.com/rigolo-api/services/authentication/domain/usecases"
	"github.com/rigolo-api/services/authentication/domain/usecases/testdata"
	"reflect"
	"testing"
)

type registerAccountTestCase struct {
	name        string
	accountRepo repositories.Accounts
	input       data.NewAccount
	expErr      error
}

var successCase = registerAccountTestCase{
	name:  "it should successfully created account",
	input: data.NewAccount{Password: "test"},
	accountRepo: testdata.BuildAccountMockRepo(
		testdata.NoFoundGetAccountByEmailFunc,
		testdata.SuccessSaveAccountFunc,
	),
	expErr: nil,
}

var alreadyUsedEmailCase = registerAccountTestCase{
	name:  "it should return already used email",
	input: data.NewAccount{Password: "test"},
	accountRepo: testdata.BuildAccountMockRepo(
		testdata.SuccessGetAccountByEmailFunc,
		testdata.SuccessSaveAccountFunc,
	),
	expErr: errors.EmailAlreadyUsed,
}

func TestRegisterAccount_Execute(t *testing.T) {

	tests := []registerAccountTestCase{
		successCase, alreadyUsedEmailCase,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := usecases.BuildAccountService(tt.accountRepo)
			err := service.RegisterAccount(tt.input)

			if !reflect.DeepEqual(tt.expErr, err) {
				t.Errorf("RegisterAccount() error = %v, expErr %v", err, tt.expErr)
			}

		})
	}
}
