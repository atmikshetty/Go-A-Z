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

	////	Variables, Strings and Numbers
	//
	//// Strings should always be within double quotes " " , single quotes are not allowed.
	////Method 1
	//var name string = "Atmik"
	//fmt.Println(name)
	//
	////Method 2
	//var surname = "Shetty"
	//fmt.Println(surname)
	//
	////Method 3... this is the most widely used method...cannot be used outside of a function
	//fullname := "Atmik Shetty"
	//fmt.Println(fullname)
	//
	////Method 4 ... declare first use later
	//var skills string // empty intially
	//skills = "Coding Gym"
	//fmt.Println(skills)
	//
	//// All of the above 4 methods can be used for declaring any types of variables, recommend using the 3rd one
	//
	//// The values can be updated as well since they are not constant
	//fullname = "Atmik Shetty Atmik Shetty"
	//fmt.Println(fullname)
	//
	//skills = "Sleeping Coding Gym"
	//fmt.Println(skills)
	//
	////	Numbers
	//num := 1000000
	//fmt.Println(num)
	//
	////	Golang also provides access to integer with bits so that we can declare as per the memory required
	//var num1 int8 = 25 // can hold upto 8 bits i.e -128 to 127
	//var num2 int16 = 180
	//num3 := 700
	//
	//fmt.Println(num1, num2, num3)
	//
	//////	Using out of range values for a 8 sized integer variable
	////var numTest int8 = 128 // does not run will throw an error
	////fmt.Println(numTest)
	//
	////	Refer the documentation for further use:
	//// https://pkg.go.dev/builtin#uint16
	//
	////	Uint types cannot have a negative number
	////	var unum uint8 = -99 // will give an erorr
	//var unum uint8 = 200 // range here is bigger than int since we are not using negative so more values can be accessed here
	//fmt.Println(unum)
	//
	////	Floats, all the methods mentioned for other above will work for float as well
	//floatnum := 99.99999
	//fmt.Println(floatnum)

	//	Printing and Formatting Strings
	//	all of the methods from fmt package start from a package because when you export a function/method from a go package it has to in capital

	//// Print: does not automatically print a new line after, for a new line use \n
	//fmt.Print("Atmik Shetty ")
	//fmt.Print("Atmik \n")
	//fmt.Print("Shetty") // will be printed on a new line
	//
	////	Println: automatically adds a new line
	//fmt.Println("Hello Atmik Shetty!!!")
	//fmt.Println("Please practice DSA!!!")

	age := 20
	name := "Atmik Shetty"
	salary := 99.12007516
	//No need for a separate space go compiler does that automatically
	fmt.Println("Age is", age)
	fmt.Println("Name is", name)
	fmt.Println("My name is", name, "and my age is", age)

	//	Formatted strings: Printf... also does not automatically add a new line later
	// the "%" followed by something listed below are called format specifiers
	// %v to display any type variable - I prefer using this
	// %s for strings
	// %d for integer
	// %f for floats
	// %q for strings but gives "" quotes around the variable
	// %T gives the type of the variable
	fmt.Printf("My name is %v and my age is %v \n", name, age)
	fmt.Printf("My name is %s and my age is %d \n", name, age)
	fmt.Printf("My salary is %f \n", salary)
	fmt.Printf("My salary is %.2f \n", salary) // can specify upto how many values do you want
	fmt.Printf("My name is %q \n", name)
	fmt.Printf("my name is %T \n", name)

	//	Sprintf: can save any formatted string
	var saving string = fmt.Sprintf("My name is %v and my age is %v \n", name, age)
	fmt.Println(saving)

}
