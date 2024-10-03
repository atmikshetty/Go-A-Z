package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello World!!!!")

	// 1. Explicit Declaration of Variables

	// Strings
	var name string = "Atmik Shetty"
	fmt.Println(name)

	//	Integers
	var age int = 20
	fmt.Println(age)

	//	Float - 2 types
	//	float32 - stores a 32-bit number, useful for faster operations

	var numFloat32 float32 = 64.7654345
	fmt.Println(numFloat32)

	//	float64 - stores a 64-bit number, slower but accurately represents numbers
	var numFloat64 float64 = 64.8888888887346374676464
	fmt.Println(numFloat64)

	// 2.	Type Inferencing - Go compiler can infer the type of any variable without you explicitly declaring them
	var surname = "Shetty"
	fmt.Println(surname)
	fmt.Println(reflect.TypeOf(surname)) // string

	var year = 2003
	fmt.Println(year)
	fmt.Println(reflect.TypeOf(year)) // int

	var f64 = 64.8888888887346374676464
	fmt.Println(f64)
	fmt.Println(reflect.TypeOf(f64)) // float64

	//	3. Shorthand for declaring variables (Make sure this is always used within a function)
	goat := "Messi"
	ballondor := 8
	fmt.Println(goat, ballondor)
	fmt.Println(reflect.TypeOf(goat), reflect.TypeOf(ballondor)) // string, int

	//	4. Constants: Immutable i.e. cannot be changed once defined, all the above 3 steps will also work for constants
	const pi = 3.1415926
	const leapYear int = 366
	fmt.Println(pi, leapYear)

	//	5. Types of Integers in Go
	//	int - int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64
	//	float - float32, float64
	//	boolean - bool
	//	strings - string
	//		More on uint i.e. Unsigned Integers later

}
