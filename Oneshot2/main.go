package main

import "fmt"

func main() {
	fmt.Println("Hello World!!!")

	//	Structs and Custom Types : Struct is a blueprint which describes a type of data
	//	 Try to create these within a better file for better code management
	myBill := makeBill("Atmik's Bill")
	//fmt.Println(myBill)

	myBill.addItems("Vada Pav", 18)
	myBill.addItems("Pav Bhaji", 120)
	myBill.addItems("Shwarma", 60)
	myBill.addItems("Anda Bhurji", 20)
	myBill.addItems("Chikcen Platter", 150)
	myBill.addItems("Samosa", 18)

	myBill.updateTip(10)

	fmt.Println(myBill.formatBill())
}
