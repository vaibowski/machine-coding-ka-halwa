package digitalwallet

type CreditCard struct {
	BasePaymentMethod
	CardNumber string
	Expiry     string
	CVV        string
}

func NewCreditCard(id string, user *User, cardNumber, expiry, cvv string) *CreditCard {
	return &CreditCard{
		BasePaymentMethod: BasePaymentMethod{
			ID:   id,
			User: user,
		},
		CardNumber: cardNumber,
		Expiry:     expiry,
		CVV:        cvv,
	}
}

func (cc *CreditCard) ProcessPayment(amount float32, currency Currency) (bool, error) {
	return true, nil
}
