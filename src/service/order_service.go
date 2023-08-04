package service

import (
	"Order-Scheduling-System/src/dao"
	"errors"
	"fmt"
	"strconv"
)

func PlaceOrder(customer *dao.Customer, shipmentID int) error {
	shipmentExisted := false
	for _, shipment := range dao.Shipments.AllShipments {
		if shipmentID == shipment.ID && shipment.StockCount > 0 {
			shipmentExisted = true
			order := dao.NewOrder(customer, shipment)
			shipment.StockCount--

			dao.OrderDatabaseInstance.AddOrder(order)
			break
		}
	}
	if !shipmentExisted {
		return errors.New("ShipmentID not existed or out of stock:" + strconv.Itoa(shipmentID))
	} else {
		return nil
	}
}

func GatherOrder(orderID int) error {
	orderExisted := false
	for _, order := range dao.OrderDatabaseInstance.OrdersByStatus[dao.Placed] {
		if orderID == order.ID {
			orderExisted = true
			err := dao.OrderDatabaseInstance.UpdateStatus(order, dao.Gathering)
			if err != nil {
				return err
			}
		}
	}
	if !orderExisted {
		return errors.New("OrderID not exist:" + strconv.Itoa(orderID))
	}

	return nil
}

func DeliverOrder(orderID int) error {
	orderExisted := false
	for _, order := range dao.OrderDatabaseInstance.OrdersByStatus[dao.Gathering] {
		if orderID == order.ID {
			orderExisted = true
			err := dao.OrderDatabaseInstance.UpdateStatus(order, dao.OutForDelivery)
			if err != nil {
				return err
			}
		}
	}
	if !orderExisted {
		return errors.New("OrderID not exist:" + strconv.Itoa(orderID))
	}
	return nil
}

func CompleteOrder(orderID int) error {
	orderExisted := false
	for _, order := range dao.OrderDatabaseInstance.OrdersByStatus[dao.OutForDelivery] {
		if orderID == order.ID {
			orderExisted = true
			err := dao.OrderDatabaseInstance.UpdateStatus(order, dao.Delivered)
			if err != nil {
				return err
			}
		}
	}
	if !orderExisted {
		return errors.New("OrderID not exist:" + strconv.Itoa(orderID))
	}
	return nil
}

func CancelOrder(orderID int) error {
	orderExisted := false
	for _, order := range dao.OrderDatabaseInstance.OrdersByStatus[dao.Placed] {
		if orderID == order.ID {
			orderExisted = true
			err := dao.OrderDatabaseInstance.UpdateStatus(order, dao.Cancelled)
			if err != nil {
				return err
			}
			order.Shipment.StockCount++
			return nil
		}
	}
	for _, order := range dao.OrderDatabaseInstance.OrdersByStatus[dao.Gathering] {
		if orderID == order.ID {
			orderExisted = true
			err := dao.OrderDatabaseInstance.UpdateStatus(order, dao.Cancelled)
			if err != nil {
				return err
			}
			order.Shipment.StockCount++
			return nil
		}
	}
	for _, order := range dao.OrderDatabaseInstance.OrdersByStatus[dao.OutForDelivery] {
		if orderID == order.ID {
			orderExisted = true
			return errors.New(fmt.Sprintf("Order ID: %d is OutForDelivery. Cannot cancel it.", orderID))
		}
	}
	for _, order := range dao.OrderDatabaseInstance.OrdersByStatus[dao.Delivered] {
		if orderID == order.ID {
			orderExisted = true
			return errors.New(fmt.Sprintf("Order ID: %d is OutForDelivery. Cannot cancel it.", orderID))
		}
	}
	if !orderExisted {
		return errors.New(fmt.Sprintf("Order ID: %d doesn't exist.", orderID))
	} else {
		return nil
	}
}
