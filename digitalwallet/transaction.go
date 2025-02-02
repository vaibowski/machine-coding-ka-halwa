package digitalwallet

import "time"

type Transaction struct {
	ID              string
	senderAccount   *Account
	receiverAccount *Account
	Amount          int64
	Currency        Currency
	Timestamp       time.Time
}

func NewTransaction(id string, senderAccount, receiverAccount *Account, amount int64, currency Currency) *Transaction {
	return &Transaction{
		ID:              id,
		senderAccount:   senderAccount,
		receiverAccount: receiverAccount,
		Amount:          amount,
		Currency:        currency,
		Timestamp:       time.Now(),
	}
}
