package facade

import "fmt"

type Ledger struct {
}

func (l *Ledger) AddNewEntry(accountId string, transactionType string, amount int) string {
	return fmt.Sprintf("new entry for account id '%s'; transaction type '%s'; amount %d",
		accountId, transactionType, amount,
	)
}
