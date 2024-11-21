package coffeemachine

type Coffee struct {
	name   string
	price  int32
	recipe map[*Ingredient]int
}

func NewCoffee(name string, price int32, recipe map[*Ingredient]int) *Coffee {
	return &Coffee{
		name:   name,
		price:  price,
		recipe: recipe,
	}
}

func (c *Coffee) getName() string {
	return c.name
}

func (c *Coffee) getPrice() int32 {
	return c.price
}

func (c *Coffee) getRecipe() map[*Ingredient]int {
	return c.recipe
}
