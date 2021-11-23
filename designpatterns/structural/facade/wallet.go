package facade

import (
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

func (w *Wallet) Withdraw(amount int) string {
	if w.Balance < amount {
		return "insufficient funds"
	}
	w.Balance -= amount
	return "transaction successful;"
}
