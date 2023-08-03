package dao

import (
	"math/rand"
	"strconv"
)

type Shipment struct {
	ID         int
	Name       string
	Category   Category
	Price      int
	StockCount int
}

var nextID = 0

func (s *Shipment) String() string {
	res := "ID" + strconv.Itoa(s.ID) + "\n"
	return res
}

func NewShipment() *Shipment {
	var s = new(Shipment)
	s.Name = "Shipment " + strconv.Itoa(nextID)
	s.ID = nextID
	s.Price = rand.Intn(100) + 100
	s.Category = Category(rand.Intn(4))
	s.StockCount = rand.Intn(10) + 5

	nextID++
	return s
}
