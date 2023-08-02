package dao

import (
	"fmt"
	"testing"
)

func TestOrderDatabase(t *testing.T) {
	OrderDatabaseInstance.Init()
	_, err := OrderDatabaseInstance.GetOrder(100)
	fmt.Println(err)

	OrderDatabaseInstance.AddOrder(NewOrder(NewCustomer("Giao"), NewShipment()))
	_, err = OrderDatabaseInstance.GetOrder(100)
	fmt.Println(err)

	order, err := OrderDatabaseInstance.GetOrder(0)
	fmt.Println(order)
}

func TestOrderDatabaseUpdateOrderStatus(t *testing.T) {
	OrderDatabaseInstance.Init()

	order := NewOrder(NewCustomer("Giao"), NewShipment())
	OrderDatabaseInstance.AddOrder(order)

	newStatus := OutForDelivery
	err := OrderDatabaseInstance.UpdateStatus(order, newStatus)
	if err != nil {
		fmt.Println(err)
	}
	if len(OrderDatabaseInstance.OrdersByStatus[newStatus]) != 1 {
		t.Errorf("Wrong")
	}
	if len(OrderDatabaseInstance.OrdersByStatus[Placed]) != 0 {
		t.Errorf("Wrong")
	}
}
