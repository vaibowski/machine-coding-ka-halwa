package digitalwallet

type PaymentMethod interface {
	GetID() string
	GetUser() *User
	ProcessPayment(amount float32, currency Currency) (bool, error)
}

type BasePaymentMethod struct {
	ID   string
	User *User
}

func (b *BasePaymentMethod) GetID() string {
	return b.ID
}

func (b *BasePaymentMethod) GetUser() *User {
	return b.User
}
