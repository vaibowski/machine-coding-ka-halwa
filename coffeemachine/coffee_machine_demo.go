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
	wg := sync.WaitGroup{}
	for i, coffeeOrder := range coffeeList {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(i) * time.Second)
			coffee := coffeeMachine.selectCoffee(coffeeOrder)
			payment := coffeeMachine.makePayment(30)
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
