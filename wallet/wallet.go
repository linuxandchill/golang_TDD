package wallet

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance Bitcoin
}

// how type is printed when used with the %s format string in prints
type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amt Bitcoin) {
	w.balance += amt
}

func (w *Wallet) GetBalance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt > w.balance {
		return errors.New("insufficient funds")
	}
	w.balance -= amt
	return nil
}
