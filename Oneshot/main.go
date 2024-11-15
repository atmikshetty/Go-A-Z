package main

// Imports
import "fmt"

// Theory: Read through this once

// Entry Point of any go application
// Always there will be only one main functions
// You can always create other functions within other files and invoke them within this main function
// You must always use a declared variable in Go, if left unused the compiler will give an error
// Similarly any imported package must also be used or else the go compiler will through an error

func main() {

	// Printing something to the console
	fmt.Println("Hello World!!!")

	//	Variables, Strings and Numbers

	// Strings should always be within double quotes " " , single quotes are not allowed.
	//Method 1
	var name string = "Atmik"
	fmt.Println(name)

	//Method 2
	var surname = "Shetty"
	fmt.Println(surname)

	//Method 3... this is the most widely used method...cannot be used outside of a function
	fullname := "Atmik Shetty"
	fmt.Println(fullname)

	//Method 4 ... declare first use later
	var skills string // empty intially
	skills = "Coding Gym"
	fmt.Println(skills)

	// All of the above 4 methods can be used for declaring any types of variables, recommend using the 3rd one

	// The values can be updated as well since they are not constant
	fullname = "Atmik Shetty Atmik Shetty"
	fmt.Println(fullname)

	skills = "Sleeping Coding Gym"
	fmt.Println(skills)

	//	Numbers
	num := 1000000
	fmt.Println(num)

	//	Golang also provides access to integer with bits so that we can declare as per the memory required
	var num1 int8 = 25 // can hold upto 8 bits i.e -128 to 127
	var num2 int16 = 180
	num3 := 700

	fmt.Println(num1, num2, num3)

	////	Using out of range values for a 8 sized integer variable
	//var numTest int8 = 128 // does not run will throw an error
	//fmt.Println(numTest)

	//	Refer the documentation for further use:
	// https://pkg.go.dev/builtin#uint16

	//	Uint types cannot have a negative number
	//	var unum uint8 = -99 // will give an erorr
	var unum uint8 = 200 // range here is bigger than int since we are not using negative so more values can be accessed here
	fmt.Println(unum)

	//	Floats, all the methods mentioned for other above will work for float as well
	floatnum := 99.99999
	fmt.Println(floatnum)
}
