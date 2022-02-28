package main

import "fmt"

func main() {
	a := 10
	b := 9

	if a > b {
		fmt.Println("A is greater than B")
	} else if a == b {
		fmt.Println("A is equal than B")
	} else {
		fmt.Println("A is less than B")
	}

	fmt.Println("-------------")

	x := "hello"
	y := "go"

	if x == "hello" {
		fmt.Println(x + " world!")
	}

	fmt.Println("-------------")

	if x == y {
		fmt.Println(x + " " + y)
	} else {
		fmt.Println("X is not equal Y.")
	}

}
