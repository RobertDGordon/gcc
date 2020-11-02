package main

import "fmt"

func main() {
	var name string = "Robert"
	var age int = 36
	const isCool = true

	lastname, size := "Gordon", 1.2

	fmt.Println(name, lastname, age, isCool)
	fmt.Printf("%T\n", size)
}
