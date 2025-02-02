package digitalwallet

import "fmt"

func Run() {
	digitalWallet := NewWalletService()

	user1 := NewUser("U01", "user1", "user1@gmail.com", "pass")
	user2 := NewUser("U02", "user2", "user2@gmail.com", "pass")

	digitalWallet.CreateUser(user1)
	digitalWallet.CreateUser(user2)

	account1 := NewAccount("a1", user1, "a1", USD)
	account2 := NewAccount("a2", user2, "a2", EUR)
	account3 := NewAccount("a3", user1, "a3", INR)

	digitalWallet.CreateAccount(account1)
	digitalWallet.CreateAccount(account2)
	digitalWallet.CreateAccount(account3)

	account1.Deposit(1000)
	account2.Deposit(2000)
	account3.Deposit(10000)

	err := digitalWallet.TransferFunds(account1, account3, 10, USD)
	if err != nil {
		fmt.Printf("Transfer failed with err: %s\n", err.Error())
	}

	fmt.Println("Transaction history for Account 1:")
	for _, transaction := range digitalWallet.GetTransactionHistory(account1) {
		fmt.Printf("Transaction ID: %s\n", transaction.ID)
		fmt.Printf("Amount: %v %s\n", transaction.Amount, transaction.Currency)
		fmt.Printf("Timestamp: %v\n\n", transaction.Timestamp)
	}

	fmt.Println("Transaction History for Account 3:")
	for _, transaction := range digitalWallet.GetTransactionHistory(account3) {
		fmt.Printf("Transaction ID: %s\n", transaction.ID)
		fmt.Printf("Amount: %v %s\n", transaction.Amount, transaction.Currency)
		fmt.Printf("Timestamp: %v\n\n", transaction.Timestamp)
	}

}
