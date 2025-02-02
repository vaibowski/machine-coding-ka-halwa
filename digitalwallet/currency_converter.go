package digitalwallet

import "sync"

type CurrencyConverter struct {
	exchangeRates map[Currency]int64
	mu            sync.RWMutex
}

var (
	currencyConverter *CurrencyConverter
	once              sync.Once
)

func getCurrencyConverter() *CurrencyConverter {
	once.Do(func() {
		currencyConverter = &CurrencyConverter{
			exchangeRates: make(map[Currency]int64),
		}
		currencyConverter.initializeExchangeRates()
	})
	return currencyConverter
}

func (cc *CurrencyConverter) initializeExchangeRates() {
	cc.mu.Lock()
	defer cc.mu.Unlock()

	cc.exchangeRates[USD] = int64(1)
	cc.exchangeRates[INR] = int64(80)
	cc.exchangeRates[EUR] = int64(2)
}

func (cc *CurrencyConverter) convert(amount int64, sourceCurrency, targetCurrency Currency) int64 {
	cc.mu.Lock()
	defer cc.mu.Unlock()
	return (amount * cc.exchangeRates[targetCurrency]) / cc.exchangeRates[sourceCurrency]
}
