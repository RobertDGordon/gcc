package main

import "fmt"

func main() {
	// var fruitArr [2]string
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	fruitArr := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(fruitArr)
	fmt.Println(len(fruitArr))
	fmt.Println(fruitArr[1:3])
}
