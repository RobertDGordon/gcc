package main

import "fmt"

func main() {
	emails := make(map[string]string)

	//Assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "shar@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	delete(emails, "Sharon")

	fmt.Println(emails)
	fmt.Println(len(emails["Bob"]))

	emails2 := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	emails2["Mike"] = "mike@gmail.com"

	fmt.Println(emails2)
}
