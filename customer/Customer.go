package customer

type CustomerInterface interface {
	String() string
}

type Customer struct {
	name  string
}

func NewCustomer(name string) *Customer {
	return &Customer{
		name: name,
	}
}

func (c *Customer) String() string {
	return c.name
}