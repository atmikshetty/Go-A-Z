package main

import "fmt"

func main() {

	//	Pointer hold the memory address of a variable
	//	Declared using the * symbol
	//	& symbol gives the address

	x := 20
	y := &x         // pointer to the variable x
	fmt.Println(x)  // Actual value
	fmt.Println(y)  // Memory address
	fmt.Println(*y) // Value at the memory address, also called dereferencing

	//	Changing the values directly within a function using Pointers
	a := 10
	fmt.Println(a) // 10
	change(&a)     // address of the variable is passed hence value changes to 100
	fmt.Println(a) // 100
}

func change(a *int) {
	*a = 100 // directly changes the og value
}
