package coffeemachine

import (
	"fmt"
	"sync"
	"time"
)

func Run() {
	fmt.Println("zook zook! coffee machine is running")
	coffeeMachine := GetCoffeeMachine()
	coffeeMachine.displayMenu()
	coffeeList := []string{"latte", "espresso", "latte", "cappuccino", "espresso"}
	paymentList := []int32{30, 20, 10, 30, 10}
	wg := sync.WaitGroup{}
	for i := range coffeeList {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(i) * time.Second)
			coffee := coffeeMachine.selectCoffee(coffeeList[i])
			payment := coffeeMachine.makePayment(paymentList[i])
			time.Sleep(2 * time.Second)
			err := coffeeMachine.dispenseCoffee(coffee, payment)
			if err != nil {
				fmt.Println(err.Error())
			}
		}()
	}
	wg.Wait()
	fmt.Println("all orders have been served")
	return
}
