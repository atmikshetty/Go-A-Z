package main

import "fmt"

func main() {

	//	1. Arrays - Fixed Size
	var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numbers)

	// using shorthand
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	//	without specifying the size
	nums1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(nums1)

	//	2. Slice - Dynamic
	var names = []string{"Atmik", "Shetty"}
	fmt.Println(names)

	//	using shorthand
	names2 := []int{45, 45, 45}
	fmt.Println(names2)

	//	3. Maps - Kind off like Dictionaries
	var people = map[string]string{"name": "Atmik", "surname": "Shetty"}
	fmt.Println(people)
	fmt.Println(people["name"])
	fmt.Println(people["surname"])

	//	4. Zero or Null values in GO
	var a int    // By Default 0
	var b bool   // false
	var c string // ""
	//	nil for maps, arrays, slices, etc.
	fmt.Println(a, b, c)
}
