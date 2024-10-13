package main

import "fmt"

// Interfaces is a type which defines a set of methods other types must implement, it does not provide any implementation, i.e. left to the other methods

type football interface {
	Goat() string
}

type Messi struct {
}

type Anthony struct {
}

func (m Messi) Goat() string {
	return "Goat No 2."
}

func (a Anthony) Goat() string {
	return "Goat No 1."
}

func printVals(val interface{}) {
	fmt.Println(val)
}

func main() {
	var footballer football
	var footballer2 football

	footballer = Messi{}
	footballer2 = Anthony{}

	fmt.Println(footballer.Goat())
	fmt.Println(footballer2.Goat())

	//	Empty Interfaces - Can hold values of any types
	printVals(10)
	printVals("Atmik Shetty")
	printVals(99.9999999)
}
