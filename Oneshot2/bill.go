package main

import "fmt"

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
		items: map[string]float64{"vada pav": 18, "tea": 10},
		tip:   0,
	}
	return b
}

// formatting the bill, simple functions can be called by anyone we want our function to be specifically used for the bill struct
func (b bill) formatBill() string { // adding (b bill) makes sure that the function follows with our custom struct
	fs := "Bill Breakdown: \n"
	var totalValue float64 = 0

	//	List the items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ... %v \n", k+":", v) // -25 give the spacing, between the values
		totalValue += v
	}
	fs += fmt.Sprintf("%-25v ... %v.2f", "total:", totalValue)
	return fs
}
