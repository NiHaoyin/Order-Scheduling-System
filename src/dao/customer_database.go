package dao

import "strconv"

type CustomerDatabase struct {
	Customers map[int]*Customer
}

var Customers = new(CustomerDatabase)

func (c *CustomerDatabase) Init() {
	c.Customers = make(map[int]*Customer)
	for i := 0; i < 100; i++ {
		customer := NewCustomer("CustomerName" + strconv.Itoa(i))
		c.Customers[customer.ID] = customer
	}

}
