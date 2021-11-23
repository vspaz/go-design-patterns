package facade

type Account struct {
	Name string
}

func NewAccount(accountName string) *Account {
	return &Account{
		Name: accountName,
	}
}

func (a *Account) IsValid(accountName string) bool {
	return a.Name == accountName
}
