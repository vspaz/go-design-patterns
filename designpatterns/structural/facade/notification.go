package facade

type Notification struct {
}

func NewNotification() *Notification {
	return &Notification{}
}

func (n *Notification) NotifyDepositTransaction() string {
	return "deposit transaction took place"
}

func (n *Notification) NotifyWithdrawalTransaction() string {
	return "withdrawal transaction took place"
}
