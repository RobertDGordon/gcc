package main

import "fmt"

func main() {
	x := 15
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	color := "orange"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is not blue or red")
	}

	//Switch
	switch color {
	case "red":
		fmt.Println("red")
	case "blue":
		fmt.Println("blue")
	default:
		fmt.Println("neither blue or red")
	}
}
