package dao

import (
	"errors"
	"strconv"
)

type OrderDatabase struct {
	AllOrders      []*Order
	OrdersByStatus map[Status][]*Order
}

var OrderDatabaseInstance = new(OrderDatabase)

func (odb *OrderDatabase) Init() {
	odb.AllOrders = make([]*Order, 0)
	odb.OrdersByStatus = make(map[Status][]*Order)
}

func (odb *OrderDatabase) AddOrder(newOrder *Order) {
	odb.AllOrders = append(odb.AllOrders, newOrder)
	odb.OrdersByStatus[newOrder.Status] = append(odb.OrdersByStatus[newOrder.Status], newOrder)
}

func (odb *OrderDatabase) GetOrder(orderID int) (*Order, error) {
	for _, order := range odb.AllOrders {
		if orderID == order.ID {
			return order, nil
		}
	}
	return nil, errors.New("Cannot find orderID:" + strconv.Itoa(orderID))
}

// UpdateStatus
// 1. Delete order from the old entry
// 2. Insert order into the new entry
// 3. Change order status
func (odb *OrderDatabase) UpdateStatus(order *Order, newStatus Status) error {
	orders := odb.OrdersByStatus[order.Status]
	isOrderExist := false
	for i, o := range orders {
		if o.ID == order.ID {
			orders = append(orders[:i], orders[i+1:]...)
			odb.OrdersByStatus[order.Status] = orders
			isOrderExist = true
			break
		}
	}
	if !isOrderExist {
		return errors.New("Order not found: " + order.String())
	}

	orders = odb.OrdersByStatus[newStatus]
	odb.OrdersByStatus[newStatus] = append(orders, order)

	order.Status = newStatus

	return nil
}
