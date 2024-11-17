package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

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

	// //Printing and Formatting Strings
	// //all of the methods from fmt package start from a package because when you export a function/method from a go package it has to in capital

	// Print: does not automatically print a new line after, for a new line use \n
	fmt.Print("Atmik Shetty ")
	fmt.Print("Atmik \n")
	fmt.Print("Shetty") // will be printed on a new line

	//	Println: automatically adds a new line
	fmt.Println("Hello Atmik Shetty!!!")
	fmt.Println("Please practice DSA!!!")

	age := 20
	name = "Atmik Shetty"
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

	//	Arrays.. fixed length cannot change
	var ages [5]int = [5]int{1, 2, 3, 4, 5}
	var ages1 = [5]int{1, 2, 3, 4, 5}
	names := [3]string{"Atmik", "Dev", "Vedant"}
	fmt.Println(ages, ages1, names)
	fmt.Printf("Lenght of names is", len(names))

	//Array values can be changed but more elements cannot be added to the array
	names[2] = "Shetty"
	fmt.Println(names)

	//	Slices... uses arrays under the hoods but are more flexible
	var scores = []int{1, 2, 3, 4, 5, 6, 6, 8}
	fmt.Println(scores)

	// slices values can be modified and more values can also be added to it
	scores = append(scores, 100)
	fmt.Println(scores)

	//	Ranges or Slicing
	rangeOne := scores[1:3] // will print values from index 1 to 2, last index - 1 is always printed
	fmt.Println(rangeOne)

	rangeTwo := scores[2:] // everything from 2 to the end
	fmt.Println(rangeTwo)

	rangeThree := scores[:5] // from start to 5
	fmt.Println(rangeThree)

	//	The standard Library
	greetings := "Hello, How are you today!!!"
	fmt.Println(greetings)

	//	Strings package
	fmt.Println(strings.Contains(greetings, "shetty")) // searches within the string

	fmt.Println(strings.ReplaceAll(greetings, "Hello", "Atmik")) // replace occurences

	fmt.Println(strings.ToUpper(greetings)) // converts to upper

	fmt.Println(strings.Index(greetings, "ll")) // return the index

	fmt.Println(strings.Split(greetings, " ")) // returns a slice of elements

	//Sorting numbers
	ages10 := []int{45, 80, 10, 2, 4, 7, 21, 32, 64, 69}
	fmt.Println("Before sorting:", ages)
	sort.Ints(ages10)
	fmt.Println("After sorting:", ages)

	//	Searching for index in the sorted slice
	index := sort.SearchInts(ages10, 4)
	fmt.Println(index)

	//	Sorting Strings
	names3 := []string{"Atmik", "Shetty", "Leo", "Messi"}
	sort.Strings(names3)
	fmt.Println(names3)

	//	Searching for index in the sorted string slice
	fmt.Println(sort.SearchStrings(names3, "Messi"))

	//	Loops
	//	While using for
	x := 0
	for x <= 10 {
		fmt.Println(" x is", x)
		x++
	}

	//	For Loop
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is Even")
		} else {
			fmt.Println(i, "is Odd")
		}
	}

	//	For strings
	names5 := []string{"Atmik", "Shetty", "Leo", "Messi"}
	for i := 0; i < len(names5); i++ {
		fmt.Println(names5[i])
	}

	//	Python like for loop
	for index, value := range names5 {
		fmt.Println(index, value)
	}

	//	to avoid indexes or anything we can use this "_"
	for _, value := range names5 {
		fmt.Println(value)
	}

	//	Booleans and Conditionals

	marks := 36
	fmt.Println(marks <= 50)
	fmt.Println(marks >= 50)
	fmt.Println(marks == 50)
	fmt.Println(marks != 50)

	//	If else
	voting := 21
	if voting > 18 {
		fmt.Println("Can vote")
	} else if voting == 18 {
		fmt.Println("Can vote")
	} else {
		fmt.Println("Cannot vote")
	}

	//	Functions
	hello("Atmik Shetty")
	hello("Leo Messi")
	cycleNames([]string{"Atmik Shetty", "Leo Messi"}, hello) // pass a slice and a function, note that () is not required since cycleNames is already handling the function call
	fmt.Println(circleArea(3))

	//	Functions with Multiple return values
	fmt.Println(getInitials("Atmik Shetty"))

	//	Package Scope: Accessing things from different go files
	// To run do go run main.go greetings.go
	// To access things from main.go make sure to define it out of the main function
	sayHello([]string{"Atmik Shetty", "Leo Messi"})

	for _, v := range values {
		fmt.Println(v)
	}

	scoring()

	//	Maps: like dictionaries in python or objects in javascript
	// keys and values in a single map must be of the same type

	menu := map[string]float64{ // key-string, value-float
		"Shwarma":        60,
		"Platter":        150,
		"Butter Chicken": 300,
	}
	fmt.Println(menu)

	// Can be accessed through the key as well
	fmt.Println("Price for Platter is:", menu["Platter"])
	fmt.Println("Price for Shwarma is:", menu["Shwarma"])

	//Printing a map
	for key, value := range menu {
		fmt.Println(key, value)
	}

	//	Updating within a map
	menu["Butter Chicken"] = 500
	fmt.Println(menu["Butter Chicken"]) // updated to 500

	//	Pass by Value: This means that go makes a copy of the values when we pass it into a function
	namefinal := "Atmik"
	fmt.Println("Original value:", namefinal)

	updateName(name)                                                   // should have updated the name to Shetty
	fmt.Println("Will not update og will update the copy:", namefinal) // but still "Atmik" is printed this is because of the pass by value property, a copy is created rather than the og value being updated

	//	One way to fix this is to make sure the function has a return type and it returns a value
	fmt.Println("Updated the og value:", updateNameReturn(namefinal)) // updates the og value

	menu1 := map[string]float64{
		"Vada Pav": 18,
		"Samosa":   18,
	}
	fmt.Println("Original Menu:", menu1)

	//	Call the update menu function
	updateMenu(menu1)
	fmt.Println("Updated Menu:", menu1) // vada pav price is updated

	//	Notice how in case of maps instead of the copy being updated directly the og value is updated.
	// This is because strings, int, arrays, floats always create a copy and then update it
	// And maps, slices and functions create a copy from the pointer to the og data and then update it hence it directly update the og value, also called as call by "reference"

	// Pointers: a variable that stores the memory address of another variable, basically by using pointers you can access the og variable
	//	In functions, maps, slices, functions go automatically does it for us but custom pointers can also be created

	name10 := "Atmik Shetty"
	fmt.Println(name10)
	fmt.Println("Memory address of name10 is:", &name10) // & is used to get the memory address for any value

	//	Creating a Pointer
	m := &name10                                       // a pointer pointing to the name10
	fmt.Println("Memory address:", m)                  // this pointer has its own memory block but also stores the address for the name10 memory block
	fmt.Println("Value at the memory address is:", *m) // through the " * " also called as "dereferencing the pointer" you can access the value at the pointer

	//	Using the pointers within a function
	// Working:
	// 1. We created name10 variable with its own memory block
	// 2. We then create a pointer to this variable which also has its own address block
	// 3. On passing this pointer to the function another copy is created which also points to the og name10 variable
	// 4. Using dereferencing we update the underlying value at the og memory address which the pointer points to

	updateNamePointer(m)                                                // unlike the updateName function this will directly update the og value
	fmt.Println("Updated value after passing it as a pointer:", name10) // will print updated value, this happened because of the use of pointers
}

func updateNamePointer(x *string) {
	*x = "Atmik Gunapala Shetty" // * is used to dereference the pointer
}

func updateMenu(x map[string]float64) {
	x["Vada Pav"] = 20
}

func updateNameReturn(x string) string {
	x = "Shetty"
	return x
}

func updateName(x string) {
	x = "Shetty"
}

var score = 100

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ") // split the input
	return names[0], names[1]      // return, usually you write a conditional to check whether there are two returns
}

func hello(name string) {
	fmt.Println("Hello", name)
}

func cycleNames(names []string, f func(string)) {
	for _, name := range names {
		f(name)
	}
}

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}
