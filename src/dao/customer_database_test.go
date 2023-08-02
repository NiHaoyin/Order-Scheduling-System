package dao

import (
	"fmt"
	"testing"
)

func TestCustomerDatabase(t *testing.T) {
	Customers.Init()
	fmt.Println(Customers)
}
