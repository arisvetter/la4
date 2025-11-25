package account

import "github.com/arisvetter/la4/customer"

type SavingsAccountInterface interface {
	AccountInterface
}

type SavingsAccount struct {
	Account
	interest float64
}

func (a *SavingsAccount) Accrue(rate float64) float64 {
	additionalInterest := a.Account.Balance() * rate
	a.interest += additionalInterest
	a.Account.Deposit(additionalInterest)
	return additionalInterest
}

func NewSavingsAccount(n string, b float64, c customer.Customer) *SavingsAccount {
	return &SavingsAccount{
		Account: Account{
			number:   n,
			balance:  b,
			customer: c,
		},
		interest: 0.0,
	}
}
