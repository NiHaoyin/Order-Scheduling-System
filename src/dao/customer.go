package dao

type Customer struct {
	Name    string
	ID      int
	Address Address
}

var nextCustomerID = 0

func NewCustomer(name string) *Customer {
	var c = new(Customer)
	c.Name = name
	c.ID = nextCustomerID

	nextCustomerID++
	return c
}

func (c *Customer) updateAddress(a Address) {
	c.Address = a
}

func (c *Customer) String() string {
	return "CustomerName: " + c.Name + "\n"
}
