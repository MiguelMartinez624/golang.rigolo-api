package domain

import (
	"github.com/rigolo-api/services/authentication/domain/data"
	"github.com/rigolo-api/services/authentication/domain/repositories"
	"github.com/rigolo-api/services/authentication/domain/usecases"
)

type Context struct {
	accountUseCases *usecases.AccountUseCases
}

func BuildContext(accountsRepo repositories.Accounts) *Context {
	context := &Context{
		accountUseCases: usecases.BuildAccountService(accountsRepo),
	}

	return context
}

func (ctx *Context) RegisterAccount(account data.NewAccount) error {
	return ctx.accountUseCases.RegisterAccount(account)
}

func (ctx *Context) AuthenticateAccount(credentials data.Credentials) (*data.PublicAccount, error) {
	return ctx.accountUseCases.AuthenticateAccount(credentials)
}
