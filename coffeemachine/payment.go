package coffeemachine

type Payment struct {
	amount int32
}

func NewPayment(amount int32) *Payment {
	return &Payment{amount: amount}
}
