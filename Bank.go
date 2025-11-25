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
	b.Add(account.NewSavingsAccount("01002", 200.0, *c))
	myChan := make(chan float64)
	go b.Accrue(myChan, 0.02)
	interest := <-myChan
	fmt.Printf("Interest from Accruement:%.2f\n", interest)
	// Please note: I am assuming we want to total the interest from the most recent accruement, not all time interest.
	fmt.Printf(b.String())
}

func test() {
	b := Bank{}
	c := customer.NewCustomer("Ann")
	b.Add(account.NewCheckingAccount("01001", 100.0, *c))
	b.Add(account.NewSavingsAccount("01002", 200.0, *c))
	b.Add(account.NewCheckingAccount("01003", 300.0, *c))
	b.Add(account.NewSavingsAccount("01004", 400.0, *c))
	b.Add(account.NewCheckingAccount("01005", 500.0, *c))
	b.Add(account.NewSavingsAccount("01006", 600.0, *c))
	b.Add(account.NewCheckingAccount("01007", 700.0, *c))
	b.Add(account.NewSavingsAccount("01008", 800.0, *c))
	b.Add(account.NewCheckingAccount("01009", 900.0, *c))
	b.Add(account.NewSavingsAccount("01010", 1000.0, *c))
	b.Add(account.NewCheckingAccount("01011", 1100.0, *c))
	b.Add(account.NewSavingsAccount("010012", 1200.0, *c))
	myChan := make(chan float64)
	go b.Accrue(myChan, 0.02)
	interest := <-myChan
	fmt.Printf("Interest from Accruement:%.2f\n", interest)
	// Please note: I am assuming we want to total the interest from the most recent accruement, not all time interest.
	fmt.Printf(b.String())
}
