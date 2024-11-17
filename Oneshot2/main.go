package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n') // will read everything after the new line i.e after the user presses enter

	return strings.TrimSpace(input), err
}

func createBill() bill {
	//	taking input from the user
	reader := bufio.NewReader(os.Stdin)

	//Method 1
	//fmt.Print("Create a New Bill: ")
	//name, _ := reader.ReadString('\n') // will read everything after the new line i.e after the user presses enter
	//name = strings.TrimSpace(name)

	//Method 2
	name, _ := getInput("Create a new Bill: ", reader)

	//	Call the function to create the bill
	b := makeBill(name)
	fmt.Println("Bill Created for:", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	option, _ := getInput("Choose a option(a - add item, s - save bill, t - add tip): ", reader)
	fmt.Println(option)
}

func main() {

	//	Structs and Custom Types : Struct is a blueprint which describes a type of data
	//	 Try to create these within a better file for better code management
	myBill := createBill()
	promptOptions(myBill)

	//myBill.addItems("Vada Pav", 18)
	//myBill.addItems("Pav Bhaji", 120)
	//myBill.addItems("Shwarma", 60)
	//myBill.addItems("Anda Bhurji", 20)
	//myBill.addItems("Chikcen Platter", 150)
	//myBill.addItems("Samosa", 18)
	//
	//myBill.updateTip(10)
	//
	//fmt.Println(myBill.formatBill())
}
