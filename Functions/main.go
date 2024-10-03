package main

import "fmt"

// Without Arguements
func Greet() {
	fmt.Println("Greetings!!!")
}

// With Arguements
func Greet2(name string) {
	// Similar to Println, just allows you to format the statements better
	fmt.Printf("Hello, %s!\n", name)
}

// With Return values - after the arguements you mention the return types
func add(a int, b int) int {
	return a + b
}

// With Multiple return values - Usually done for Error Handling
func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	} else {
		return a / b, nil
	}
}

// Functions with Named return values
func addition(a, b int) (result int) {
	result = a + b
	return // do not need to return anything since we have already named our return
}

func main() {
	// Calling a Function
	Greet()
	Greet2("Atmik Shetty")
	fmt.Println(add(5, 5))
	res, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	fmt.Println(divide(10, 5))
	fmt.Println(addition(10, 10))
}
