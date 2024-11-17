package main

// Custom type for bills
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// func to make new bills
func makeBill(name string) bill { // returns a bill type object
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}
