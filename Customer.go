package customer

type CustomerInterface interface {
	String() string
}

struct Customer struct {
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