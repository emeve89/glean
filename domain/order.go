package domain

import "errors"

type OrderRepository interface {
	Store(order Order)
	FindById(id int) Order
}

type Order struct {
	Id 			int
	Customer 	Customer
	Items 		[]Item
}

func (order *Order) Add(item Item) error {
	if !item.Available {
		return errors.New("Cannot add unavailable item to order")
	}
	if order.value() + item.Value > 250.00 {
		return errors.New("An order may not exceed a total value of $250.00")
	}

	order.Items = append(order.Items, item)
	return nil
}

func (order *Order) value() float64 {
	sum := 0.0
	for i := range order.Items {
		sum = sum + order.Items[i].Value
	}
	return sum
}