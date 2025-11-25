package savingsAccount

import "github.com/arisvetter/la4/account"

type SavingsAccountInterface interface {
	account.AccountInterface
}

type SavingsAccount struct {
	account.Account
	Interest float64
}

func (a *SavingsAccount) Accrue(rate float64) {
	a.Interest += a.Account.Balance() * rate
	a.Account.Deposit(a.Account.Balance() * rate)
}
