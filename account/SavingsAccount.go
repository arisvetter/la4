package savingsAccount

type SavingsAccount interface{
	account.AccountInterface
	Accrue()
}

type SavingsAccount struct{
	account.Account
	Interest double
}

func (s *SavingsAccount) Accrue() {
	s.balance = s.balance * s.Interest
}