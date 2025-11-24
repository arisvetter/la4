package savingsAccount

import "github.com/arisvetter/la4/account"

type SavingsAccountInterface interface {
	account.AccountInterface
	Accrue()
}

type SavingsAccount struct {
	account.Account
	Interest float64
}

func (s *SavingsAccount) Accrue() {
	s.Account.Deposit(s.Account.Balance() * s.Interest)
}
