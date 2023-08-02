package dao

import "strconv"

type Order struct {
	ID          int
	Customer    *Customer
	Shipment    *Shipment
	Destination Address
	Status      Status
}

var nextOrderID = 0

func NewOrder(c *Customer, s *Shipment) *Order {
	var order = new(Order)
	order.ID = nextOrderID
	order.Customer = c
	order.Shipment = s
	order.Status = Placed

	nextOrderID++
	return order
}

func (o *Order) String() string {
	return "OrderID is: " + strconv.Itoa(o.ID) + "\n"
}
