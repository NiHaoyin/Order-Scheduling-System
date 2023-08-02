package dao

type ShipmentDatabase struct {
	AllShipments        []*Shipment
	ShipmentsByCategory map[Category][]*Shipment
}

var Shipments = new(ShipmentDatabase)

func (s *ShipmentDatabase) Init() {
	s.AllShipments = make([]*Shipment, 100)
	for i := 0; i < 100; i++ {
		s.AllShipments[i] = NewShipment()
	}

	s.ShipmentsByCategory = make(map[Category][]*Shipment)
	for _, shipment := range s.AllShipments {
		s.ShipmentsByCategory[shipment.Category] = append(s.ShipmentsByCategory[shipment.Category], shipment)
	}
}
