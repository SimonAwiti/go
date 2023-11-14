package main

import "fmt"

func main() {
	//fmt.Print("No new")
	//fmt.Print("in print as compared to println")
	//fmt.Println("prints an output in a new line")

	name := "Simon"
	age := 28
	score := 1.999923

	fmt.Println("mys name is ", name, " and I am ", age)

	// formated string printf
	// %v is a format specifier n it can be any letter

	fmt.Printf("My name is %v and I am %v yeays old \n", name, age)

	//%q prints quoted strings

	fmt.Printf("my name is %q \n", name)

	//printing floats

	fmt.Printf("Your score is %f \n", score)

	//rounding off floats

	fmt.Printf("Your score rounded off to the nearest two is %0.2f", score)

	//saving formated strigs

	var str = fmt.Sprintf("My name is %v and I am %v yeays old \n", name, age)

	fmt.Println(str)

}
