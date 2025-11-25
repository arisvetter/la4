package account

import (
	"github.com/arisvetter/la4/customer"
)

type CheckingAccountInterface interface {
	AccountInterface
}

type CheckingAccount struct {
	Account
}

func NewCheckingAccount(n string, b float64, c customer.Customer) *CheckingAccount {
	return &CheckingAccount{
		Account: Account{
			number:   n,
			balance:  b,
			customer: c,
		},
	}
}
