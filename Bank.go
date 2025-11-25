package main

import (
	"fmt"

	"github.com/arisvetter/la4/account"
	"github.com/arisvetter/la4/customer"
)

type BankInterface interface {
	Add(a account.AccountInterface)
	Accrue(c chan<- float64, rate float64)
	String() string
	Main()
}

type Bank struct {
	accounts map[account.AccountInterface]struct{}
}

func (b *Bank) Add(a account.AccountInterface) {
	if b.accounts == nil {
		b.accounts = make(map[account.AccountInterface]struct{})
	}
	b.accounts[a] = struct{}{}
}

func (b *Bank) Accrue(c chan<- float64, rate float64) {
	interest := 0.0
	for a := range b.accounts {
		interest += a.Accrue(rate)
	}
	c <- interest
}

func (b *Bank) String() string {
	r := ""
	for account := range b.accounts {
		r += account.String() + "\n"
	}
	return r
}

func main() {
	b := Bank{}
	c := customer.NewCustomer("Ann")
	b.Add(account.NewCheckingAccount("01001", 100.0, *c))
	b.Add(account.NewSavingsAccount("02001", 200.0, *c))
	myChan := make(chan float64)
	go b.Accrue(myChan, 0.02)
	interest := <-myChan
	fmt.Printf("Interest from Accruement:%.2f\n", interest)
	// Please note: I am assuming we want to total the interest from the most recent accruement, not all time interest.
	fmt.Printf(b.String())
}
