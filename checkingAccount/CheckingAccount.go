package checkingAccount

import "github.com/arisvetter/la4/account"

type CheckingAccountInterface interface {
	account.AccountInterface
	Accrue()
}

type CheckingAccount struct {
	account.Account
}

func (c *CheckingAccount) Accrue() {
	// Do nothing here
}
