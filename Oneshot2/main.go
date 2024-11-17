package main

import "fmt"

func main() {
	fmt.Println("Hello World!!!")

	//	Structs and Custom Types : Struct is a blueprint which describes a type of data
	//	 Try to create these within a better file for better code management
	myBill := makeBill("Atmik's Bill")
	fmt.Println(myBill)
}
