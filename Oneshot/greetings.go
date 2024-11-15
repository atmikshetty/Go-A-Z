package main

import "fmt"

var values = []int{1, 2, 3, 45, 77, 6, 6, 6, 6}

func sayHello(names []string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func scoring() {
	fmt.Println("Score is", score)
}
