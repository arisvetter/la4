package account

import "customer"

type AccountInterface interface {
	Balance() float64
	Deposit()
	Withdraw()
	String() string
}

type Account struct {
	number   string
	balance  float64
	customer customer.Customer
}

func Deposit(a *Account, amount float64) {
	a.balance += amount
}

func Withdraw(a *Account, amount float64) {
	a.balance -= amount
}

func (a *Account) Balance() float64 {
	return a.balance
}

func (a *Account) String() string {
	return a.number + " " + a.customer.String() + a.balance
}
