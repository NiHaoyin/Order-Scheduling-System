package service

import (
	"Order-Scheduling-System/src/dao"
	"errors"
	"strconv"
)

func PlaceOrder(customer *dao.Customer, shipmentID int) error {
	shipmentExisted := false
	for _, shipment := range dao.Shipments.AllShipments {
		if shipmentID == shipment.ID {
			shipmentExisted = true
			order := dao.NewOrder(customer, shipment)
			dao.OrderDatabaseInstance.AddOrder(order)
			break
		}
	}
	if !shipmentExisted {
		return errors.New("ShipmentID not exist:" + strconv.Itoa(shipmentID))
	} else {
		return nil
	}
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
