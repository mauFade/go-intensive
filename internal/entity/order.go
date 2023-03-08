package entity

import (
	"errors"
)

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}

	err := order.Validate()

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (order *Order) Validate() error {
	if order.ID == "" {
		return errors.New("Order id is required")
	}

	if order.Price <= 0 {
		return errors.New("Order price is required")
	}

	if order.Tax <= 0 {
		return errors.New("Order tax is required")
	}

	return nil
}

func (order *Order) CalculateFinalPrice() error {
	order.FinalPrice = order.Price + order.Tax

	err := order.Validate()

	if err != nil {
		return err
	}

	return nil
}
