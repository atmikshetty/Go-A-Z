package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

func divide2(a, b int) int {
	if b == 0 {
		panic("Cannot divide by Zero")
	} else {
		return a / b
	}
}

func divide3(a, b int) int {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Trying to recover from the panic", r) // will execute once it faces a panic
		}
	}()

	if b == 0 {
		panic("Again cannot divide by Zero!!!!!")
	}
	return a / b
}

func main() {

	//	Go provides an "errors" package to perform Error Handling
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	//	Panic - A runtime error that might stop the execution of the program
	//	Generally used when a program cannot recover
	fmt.Println("Starting program...")
	ans := divide2(4, 2) // If b == 0, This will cause a panic which will immediately stop the program
	fmt.Println("Result:", ans)
	fmt.Println("End of program") // this will never be executed in case of panic

	// Recover - to regain control from a panicking program
	//	Can only be called inside a deferred function
	//	 Deferred function is a function which only executes when it's surrounding function finishes its execution

	fmt.Println("Starting the Divide 3 Function")
	final := divide3(10, 0) // causes a panic
	fmt.Println("Result:", final)
	fmt.Println("Program has been recovered, further operations can now execute successfully!!!") // executed after execution
}
