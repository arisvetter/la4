package main

import (
	"github.com/arisvetter/la4/account"
	"github.com/arisvetter/la4/checkingAccount"
	"github.com/arisvetter/la4/customer"
	"github.com/arisvetter/la4/savingsAccount"
)

type BankInterface interface {
	Add()
	Accrue()
	String() string
	Main()
}

type Bank struct {
	AccountList map[account.AccountInterface]struct{}
}

func (b *Bank) Add(a account.AccountInterface) {
	if b.AccountList == nil {
		b.AccountList = make(map[account.AccountInterface]struct{})
	}
	b.AccountList[a] = struct{}{}
}

func (b *Bank) Accrue() {
	for account := range b.AccountList {
		account.Accrue()
	}
}

func (b *Bank) String() string {
	r := ""
	for account := range b.AccountList {
		r += account.String() + "\n"
	}
	return r
}

func Main() {
	b := Bank{}
	c := customer.NewCustomer("Ann")
	b.Add(&checkingAccount.CheckingAccount{Account: account.Account{Number: "01001", BalanceField: 100.0, Customer: *c}})
	b.Add(&savingsAccount.SavingsAccount{Account: account.Account{Number:"01002", Customer: *c, BalanceField: 200.0}, Interest: 0.00})
	b.Accrue()
	println(b.String())
}
