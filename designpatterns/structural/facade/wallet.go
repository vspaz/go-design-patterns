package facade

import (
	"errors"
	"fmt"
)

type Wallet struct {
	Balance int
}

func NewWallet() *Wallet {
	return &Wallet{
		Balance: 0,
	}
}

func (w *Wallet) Add(amount int) string {
	w.Balance += amount
	return fmt.Sprintf("'%d' added to account", amount)
}

func (w *Wallet) Withdraw(amount int) error {
	if w.Balance < amount {
		return errors.New("insufficient funds")
	}
	w.Balance -= amount
	return nil
}
