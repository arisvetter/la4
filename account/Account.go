package account

import (
	"strconv"

	"github.com/arisvetter/la4/customer"
)

type AccountInterface interface {
	Balance() float64
	Deposit(amount float64)
	Withdraw(amount float64)
	String() string
	Accrue(rate float64) float64
}

type Account struct {
	number   string
	balance  float64
	customer customer.Customer
}

func (a *Account) Accrue(rate float64) float64 {
	return 0.00
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
}

func (a *Account) Withdraw(amount float64) {
	a.balance -= amount
}

func (a *Account) Balance() float64 {
	return a.balance
}

func (a *Account) String() string {
	return a.number + ":" + a.customer.String() + ":" + strconv.FormatFloat(a.balance, 'f', 2, 64)
}
