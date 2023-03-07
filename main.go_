package main

import "fmt"

type Order struct {
	ID       string
	Price    float64
	Quantity int
}

func (order *Order) getTotal() float64 {
	return order.Price * float64(order.Quantity)
}

func (order Order) setPrice(price float64) {
	order.Price = price

	fmt.Println("Price: ", order.Price)
}

func main() {
	order := Order{
		ID:       "123",
		Price:    10.5,
		Quantity: 5,
	}

	order.setPrice(74.89)
	fmt.Println(order.getTotal())
}
