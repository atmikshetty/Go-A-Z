package main

import "fmt"

type Coder struct {
	Name   string
	domain string
	YOE    int
	skills string
}

func (c Coder) Greet() {
	fmt.Println("Hello", c.Name)
}

func (c *Coder) addYOE() {
	c.YOE++
}

func main() {

	//	Structs: Collection of fields and entities which allows us to group related data together or create complex Data Types
	coder1 := Coder{Name: "Atmik Shetty", domain: "Data Science", YOE: 2, skills: "PyTorch, CUDA, Parallel Computing"}

	// Entire struct
	fmt.Println(coder1)

	//	Accessing fields
	fmt.Println(coder1.Name)
	fmt.Println(coder1.domain)
	fmt.Println(coder1.YOE)
	fmt.Println(coder1.skills)

	//	Also these fields can be modified
	fmt.Printf("Previous domain is %s\n", coder1.domain)
	coder1.domain = "Deep Learning"
	fmt.Printf("New domain is %s\n", coder1.domain)

	//	Methods - Just functions but can be used directly on a Struct
	coder1.Greet()

	//	Fields of a Struct can also be modified using Methods
	fmt.Printf("YOE is %d\n", coder1.YOE)
	coder1.addYOE()
	fmt.Printf("New YOE is %d\n", coder1.YOE)
}
