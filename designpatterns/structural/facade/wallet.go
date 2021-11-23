package facade

type Wallet struct {
	Balance int
}

func NewWallet() *Wallet {
	return &Wallet{
		Balance: 0,
	}
}

func (w *Wallet) Add(amount int) {
	w.Balance += amount
}
