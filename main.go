package main

import "fmt"

type Coster interface {
	Cost() float64
}

func displayCost(c Coster) {
	fmt.Println("cost: ", c.Cost())
}

type Item struct {
	name string
	price float64
	quantity int
}

func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

type DiscountItem struct {
	Item
	discountRate float64
}

func (t DiscountItem) Cost() float64 {
	return t.Item.Cost() * (1.0 - t.discountRate/100)
}

func main() {
	shoes := Item{"shoes", 30000, 2}
	eventShoes := DiscountItem{
		Item{"eventShoes", 50000, 3},
		10.00,
	}
	displayCost(shoes)
	displayCost(eventShoes)
}