package main

import "fmt"

func main() {

	//	1. IF ELSE - Opening brace should always be on the same line as the if/else
	age := 20
	if age >= 18 {
		fmt.Println("Eligible to vote!!!")
	} else {
		fmt.Println("Not eligible to vote!!!")
	}

	//	Variables can be initialized with the IF statement itself
	if deadlift := 750; deadlift > 405 {
		fmt.Println("Strong!!!")
	} else {
		fmt.Println("Very Weak!!!")
	}

	//	2. Switch - No need for break statement, go breaks out after matching a case
	day := 5
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Weekends")
	}

	//	Multiple cases can be done at once
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekday")
	default:
		fmt.Println("Weekends")
	}

	//	Can also be used as Multiple IF-ELSE statements by eliminating expression
	switch {
	case age < 13:
		fmt.Println("Child")
	case age < 18:
		fmt.Println("Teenager")
	default:
		fmt.Println("Adult")
	}

	//	3. For Loops
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//	Can also be used as a While Loop
	m := 10
	for m <= 20 {
		fmt.Println(m)
		m++
	}

	//	4. Iterating over arrays,slices, maps, etc. - Use range keyword
	// By default returns both the index and the value
	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for ind, val := range nums {
		fmt.Println(ind, val)
	}

	//	 To omit Indexes you can use '_'
	for _, val := range nums {
		fmt.Println(val)
	}

	//	5. defer - schedules a function call to be executed only when the surrounding function has finished it's executiom
	fmt.Println("Atmik")
	defer fmt.Println("Will be printed last!!!")
	fmt.Println("Shetty")

}
