package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	//	Switch case provide user to access options
	switch option {
	case "a":
		name, _ := getInput("Name of the Item:", reader)
		price, _ := getInput("Price of the Item:", reader)

		//Parsing Floats
		p, err := strconv.ParseFloat(price, 64) // converts a string to float
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		} else {
			b.addItems(name, p)
		}

		fmt.Println("Item added to the Menu!!")
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter the tip amount:", reader)

		//Parsing floats
		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Print("The tip must be a number!!!")
			promptOptions(b)
		} else {
			b.updateTip(t)
		}

		fmt.Println("You chose to give:", tip)
		promptOptions(b)
	case "s":
		fmt.Println(b.formatBill())
		fmt.Println("Save the tip.")
	default:
		fmt.Println("Invalid option")
		promptOptions(b) // fires of the func and asks the user to choose something again
	}
}

func main() {

	//	Structs and Custom Types : Struct is a blueprint which describes a type of data
	//	 Try to create these within a better file for better code management
	myBill := createBill()
	promptOptions(myBill)

}
