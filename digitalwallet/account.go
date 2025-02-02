package digitalwallet

import (
	"errors"
	"sync"
)

const insufficientBalanceError = "insufficient funds in the account"

type Account struct {
	ID            string
	user          *User
	AccountNumber string
	Currency      Currency
	balance       int64
	transactions  []*Transaction
	mu            sync.RWMutex
}

func NewAccount(id string, user *User, accountNumber string, currency Currency) *Account {
	return &Account{
		ID:            id,
		user:          user,
		AccountNumber: accountNumber,
		Currency:      currency,
		balance:       0,
		transactions:  []*Transaction{},
	}
}

func (a *Account) Deposit(amount int64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

func (a *Account) Withdraw(amount int64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance < amount {
		return errors.New(insufficientBalanceError)
	} else {
		a.balance -= amount
		return nil
	}
}

func (a *Account) AddTransaction(transaction *Transaction) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.transactions = append(a.transactions, transaction)
}

func (a *Account) GetTransactions() []*Transaction {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return append([]*Transaction{}, a.transactions...)
}

func (a *Account) GetBalance() int64 {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.balance
}
