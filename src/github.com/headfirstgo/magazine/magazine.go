package magazine

type Subscriber struct{
	name string
	rate float64
	active bool
	Address
}
type Employee struct {
	Name string
	Salary float64
	Address
}

type Address struct {
	Street string
	City string
	State string
	PostalCode string
}

