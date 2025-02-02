package digitalwallet

type BankAccount struct {
	BasePaymentMethod
	AccountNumber string
}

func NewBankAccount(id string, user *User, accountNumber string) *BankAccount {
	return &BankAccount{
		BasePaymentMethod: BasePaymentMethod{
			ID:   id,
			User: user,
		},
		AccountNumber: accountNumber,
	}
}

func (ba *BankAccount) ProcessPayment(amount int64, currency Currency) (bool, error) {
	return true, nil
}
