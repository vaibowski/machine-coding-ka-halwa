package coffeemachine

import (
	"errors"
	"fmt"
	"sync"
	"time"
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
		cm.ingredients["water"]:  6,
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

func (cm *CoffeeMachine) makePayment(amount int32) Payment {
	return Payment{amount: amount}
}

func (cm *CoffeeMachine) selectCoffee(name string) *Coffee {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	for _, item := range cm.coffeeMenu {
		if item.name == name {
			return item
		}
	}
	return nil
}

func (cm *CoffeeMachine) dispenseCoffee(coffee *Coffee, payment Payment) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	if coffee == nil {
		return errors.New("invalid choice of coffee")
	}
	if payment.amount < coffee.price {
		return errors.New(fmt.Sprintf("insufficient payment for %s, price: %d, paid: %d\n",
			coffee.name, coffee.price, payment.amount))
	}
	_, err := cm.checkIngredients(coffee)
	if err != nil {
		return err
	} else {
		fmt.Printf("%s has been served. enjoy!!\n", coffee.name)
		cm.updateIngredients(coffee)
		return nil
	}
}

func (cm *CoffeeMachine) checkIngredients(coffee *Coffee) (bool, error) {
	for ingredient, reqQuantity := range coffee.getRecipe() {
		if ingredient.GetQuantity() < reqQuantity {
			return false, errors.New(fmt.Sprintf("can't brew %s, running low on %s", coffee.name, ingredient.name))
		}
	}
	return true, nil
}

func (cm *CoffeeMachine) updateIngredients(coffee *Coffee) {
	for ingredient, reqQuantity := range coffee.getRecipe() {
		ingredient.UpdateQuantity(-reqQuantity)
		if ingredient.GetQuantity() < 3 {
			fmt.Printf("%s running low\n", ingredient.name)
		}
	}
	time.Sleep(1 * time.Second)
}
