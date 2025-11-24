package account

type CheckingAccountInterface interface {
	AccountInterface
	Accrue()
}

type CheckingAccount struct {
	Account
}

func (c *CheckingAccount) Accrue() void {
	// Do nothing here
}