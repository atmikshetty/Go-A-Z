package main

import "fmt"

func main() {

	//	Type Casting - Converting a variable into different type before using it
	// Go does not support implicit casting so has to be done explicitly i.e. manually
	var a int = 10
	var b float32 = 3.14
	sum := a + int(b)
	fmt.Println(sum)

	//	Type Inference means Go understands the type of the variable by itself
	d := 10
	f := "Atmik"
	g := 6.9999
	fmt.Printf("d is of type %T\n", d)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)
}
