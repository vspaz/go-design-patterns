package facade

type WalletFacade struct {
	Account      *Account
	Wallet       *Wallet
	SecurityCode *SecurityCode
	Notification *Notification
	Ledger       *Ledger
}

func NewWalletFacade(accountId string, code int) *WalletFacade {
	return &WalletFacade{
		Account:      NewAccount(accountId),
		SecurityCode: NewSecurityCode(code),
		Wallet:       NewWallet(),
		Notification: NewNotification(),
		Ledger:       NewLedger(),
	}
}
