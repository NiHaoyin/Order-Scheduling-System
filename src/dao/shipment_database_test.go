package dao

import (
	"fmt"
	"testing"
)

func TestShipmentDatabase(t *testing.T) {
	Shipments.Init()
	fmt.Println(Shipments)
}
