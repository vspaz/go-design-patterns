package facade

import (
	"fmt"
	"strings"
)

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

func (w *WalletFacade) Deposit(accountId string, securityCode int, amount int) string {
	if !w.Account.IsValid(accountId) {
		return fmt.Sprintf("account '%s' does not exist")
	}
	if !w.SecurityCode.IsAuthorizationCodeCorrect(securityCode) {
		return fmt.Sprintf("security code invalid")
	}
	operation_1 := w.Wallet.Add(amount)
	operation_2 := w.Notification.NotifyDepositTransaction()
	operation_3 := w.Ledger.AddNewEntry(accountId, "deposit", amount)
	return strings.Join([]string{operation_1, operation_2, operation_3}, " ")
}

func (w *WalletFacade) Withdraw(accountId string, securityCode int, amount int) string {
	if !w.Account.IsValid(accountId) {
		return fmt.Sprintf("account '%s' does not exist")
	}
	if !w.SecurityCode.IsAuthorizationCodeCorrect(securityCode) {
		return fmt.Sprintf("security code invalid")
	}
	operation_1 := w.Wallet.Withdraw(amount)
	operation_2 := w.Notification.NotifyWithdrawalTransaction()
	operation_3 := w.Ledger.AddNewEntry(accountId, "withdrawal", amount)
	return strings.Join([]string{operation_1.Error(), operation_2, operation_3}, " ")
}
