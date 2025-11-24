package main

type BankInterface interface {
	Add() void
	Acccrue() void
	String() string
	Main() void
}

type Bank struct {
	AccountList map[account.AccountInterface]set{}
}

func (b *Bank) Add() {
	b.AccountList = make(map[account.AccountInterface]set)
}

func (b *Bank) Acccrue() {
	for account := range b.AccountList {
		account.accrue()
	}
}

func (b *Bank) String() string {
	r = ""
	for account := range b.AccountList {
		r += account.String() + "\n"
	}
	return r
}

func Main() {
	b := Bank{} //todo chekc if this is right?
	c customer.CustomerInterface = Customer.NewCustomer("Alice")
	b.add(checkingAccount.NewCheckingAccount("01001", c, 100.0))
	b.add(savingsAccount.NewSavingsAccount("01002", c, 200.0))
	b.accrue(.02)
	println(b.String())
}