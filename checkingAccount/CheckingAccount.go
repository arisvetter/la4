package checkingAccount

import "github.com/arisvetter/la4/account"

type CheckingAccountInterface interface {
	account.AccountInterface
}

type CheckingAccount struct {
	account.Account
}

func (c *CheckingAccount) Accrue(rate float64) {
	// Do nothing here, todo check if this is needed or no
}
