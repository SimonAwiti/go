package main

import "fmt"

// this is a single line comment

/* multi line
comments */
// anyname := simon you cannot declare a var out a function like this

var anyname string = "simon" // when declared out the function it can be unused

func main() {

	var name string
	var age int

	fmt.Print("enter your name: ")
	fmt.Print("enter your age: ")
	fmt.Scanf("%s %d", name, age)
	if age != 28 {
		fmt.Println("verification failed")
	} else {
		fmt.Println("Hello simon")
	}
}
