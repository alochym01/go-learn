package pointers

import (
	"errors"
	"fmt"
)

// ErrErrInsufficientFunds is erroerrors display to users maybe overuse
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Bitcoin is a new type of currency which base on INT type
type Bitcoin int

// Wallet represents a Wallet
type Wallet struct {
	balance Bitcoin
}

// Deposit function takes a Bitcoin as parameter
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
	return
}

// Balance function return a balance of Wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw function takes a Bitcoin as parameter
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// String function is format Bitcoin's printing
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
