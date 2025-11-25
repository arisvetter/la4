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
	Accrue(rate float64)
}

type Account struct {
	Number       string
	BalanceField float64
	Customer     customer.Customer
}

func (a *Account) Accrue(rate float64) {}

func (a *Account) Deposit(amount float64) {
	a.BalanceField += amount
}

func (a *Account) Withdraw(amount float64) {
	a.BalanceField -= amount
}

func (a *Account) Balance() float64 {
	return a.BalanceField
}

func (a *Account) String() string {
	return a.Number + " " + a.Customer.String() + strconv.FormatFloat(a.BalanceField, 'f', 2, 64)
}
