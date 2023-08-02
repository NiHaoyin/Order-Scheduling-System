package dao

type Status int

const (
	Placed         Status = 0
	Gathering      Status = 1
	OutForDelivery Status = 2
	Delivered      Status = 3
	Cancelled      Status = -1
)
