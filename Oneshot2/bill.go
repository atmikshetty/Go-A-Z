package main

import (
	"fmt"
	"os"
)

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

// This is called a receiver function which is associated with a type like a struct
// formatting the bill, simple functions can be called by anyone we want our function to be specifically used for the bill struct
func (b bill) formatBill() string { // adding (b bill) makes sure that the function follows with our custom struct
	fs := "Bill Breakdown: \n"
	var totalValue float64 = 0

	//	List the items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ... $%v \n", k+":", v) // -25 give the spacing, between the values
		totalValue += v
	}

	// Display Tip
	fs += fmt.Sprintf("%-25v ... $%v \n", "tip:", b.tip)

	// Display total
	fs += fmt.Sprintf("%-25v ... $%0.2f", "total:", totalValue+b.tip)
	return fs
}

// Function to add tips... simply calling the vals won't work so use it as a pointer
// When working with structs you do not need to derefernce it, go automatically does it for you
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// Function to add items to the menu
func (b *bill) addItems(name string, price float64) {
	b.items[name] = price
}

func (b *bill) saveBill() {
	//	saving data as a slice
	data := []byte(b.formatBill())

	//	saving to a file
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("file saved successfully!!!")
}
