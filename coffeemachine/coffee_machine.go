package coffeemachine

import (
	"fmt"
	"sync"
)

type CoffeeMachine struct {
	coffeeMenu  []*Coffee
	ingredients map[string]*Ingredient
	mu          sync.Mutex
}

var (
	instance *CoffeeMachine
	once     sync.Once
)

func GetCoffeeMachine() *CoffeeMachine {
	once.Do(func() {
		instance = &CoffeeMachine{
			coffeeMenu:  make([]*Coffee, 0),
			ingredients: make(map[string]*Ingredient),
		}
		instance.initIngredients()
		instance.initCoffeeMenu()
	})
	return instance
}

func (cm *CoffeeMachine) initIngredients() {
	cm.ingredients["coffee"] = NewIngredient("coffee", 10)
	cm.ingredients["milk"] = NewIngredient("milk", 10)
	cm.ingredients["water"] = NewIngredient("water", 10)
}

func (cm *CoffeeMachine) initCoffeeMenu() {
	latteRecipe := map[*Ingredient]int{
		cm.ingredients["coffee"]: 2,
		cm.ingredients["milk"]:   3,
		cm.ingredients["water"]:  1,
	}
	cappuccinoRecipe := map[*Ingredient]int{
		cm.ingredients["coffee"]: 2,
		cm.ingredients["milk"]:   2,
		cm.ingredients["water"]:  2,
	}
	espressoRecipe := map[*Ingredient]int{
		cm.ingredients["coffee"]: 3,
		cm.ingredients["milk"]:   0,
		cm.ingredients["water"]:  3,
	}

	latte := NewCoffee("latte", 20, latteRecipe)
	cappuccino := NewCoffee("cappuccino", 30, cappuccinoRecipe)
	espresso := NewCoffee("espresso", 10, espressoRecipe)
	cm.coffeeMenu = append(cm.coffeeMenu, latte, cappuccino, espresso)
}

func (cm *CoffeeMachine) displayMenu() {
	fmt.Println("CoffeeMania menu:")
	for _, item := range cm.coffeeMenu {
		fmt.Printf("%s -> Rs.%d\n", item.getName(), item.getPrice())
	}
}

//func (cm *CoffeeMachine)
