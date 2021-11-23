package facade

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFacadeOk(t *testing.T) {
	walletFacade := NewWalletFacade("noname", 12345)
	transactionResult := walletFacade.Deposit("noname", 12345, 500)
	assert.Equal(t,
		"'500' added to account deposit transaction took place new entry for account id 'noname'; transaction type 'deposit'; amount 500",
		transactionResult)

	transactionResult = walletFacade.Withdraw("noname", 12345, 200)
	assert.Equal(t,
		"transaction successful; withdrawal transaction took place new entry for account id 'noname'; transaction type 'withdrawal'; amount 200",
		transactionResult)
}
