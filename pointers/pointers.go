package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficentFunds = errors.New("you can't withdraw more than you have")

type Stringer interface {
	String() string
}
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(value Bitcoin) {
	w.balance += value
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(value Bitcoin) error {
	if value > w.balance {
		return ErrInsufficentFunds
	}
	w.balance -= value
	return nil
}
