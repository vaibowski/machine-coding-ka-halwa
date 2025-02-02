package digitalwallet

import (
	"fmt"
	"sync"
	"time"
)

type WalletService struct {
	users    map[string]*User
	accounts map[string]*Account
	mu       sync.RWMutex
}

var (
	walletService *WalletService
	walletOnce    sync.Once
)

func NewWalletService() *WalletService {
	walletOnce.Do(func() {
		walletService = &WalletService{
			users:    make(map[string]*User),
			accounts: make(map[string]*Account),
		}
	})
	return walletService
}

func (ws *WalletService) CreateUser(user *User) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.users[user.ID] = user
}

func (ws *WalletService) CreateAccount(account *Account) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.accounts[account.ID] = account
	account.user.AddAccount(account)
}

func (ws *WalletService) TransferFunds(senderAccount, receiverAccount *Account, amount int64, currency Currency) error {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	converter := getCurrencyConverter()

	senderAmount := amount
	if senderAccount.Currency != currency {
		senderAmount = converter.convert(senderAmount, currency, senderAccount.Currency)
	}
	err := senderAccount.Withdraw(senderAmount)
	if err != nil {
		return err
	}

	receiverAmount := amount
	if receiverAccount.Currency != currency {
		receiverAmount = converter.convert(receiverAmount, currency, receiverAccount.Currency)
	}
	receiverAccount.Deposit(receiverAmount)

	txnID := fmt.Sprintf("T%d", time.Now().UnixNano())
	transaction := NewTransaction(txnID, senderAccount, receiverAccount, amount, currency)
	senderAccount.AddTransaction(transaction)
	receiverAccount.AddTransaction(transaction)

	return nil
}

func (ws *WalletService) GetTransactionHistory(account *Account) []*Transaction {
	return account.GetTransactions()
}
