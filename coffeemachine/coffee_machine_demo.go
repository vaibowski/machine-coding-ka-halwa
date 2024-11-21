package coffeemachine

import "fmt"

func Run() {
	fmt.Println("zook zook! coffee machine is running")
	coffeeMachine := GetCoffeeMachine()
	coffeeMachine.displayMenu()
	return
}
