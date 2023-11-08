package main

import "fmt"

// this is a single line comment

/* multi line
comments */
// anyname := simon you cannot declare a var out a function like this

var anyname string = "simon" // when declared out the function it can be unused

func main() {

	var num1 int = 1 // variables can be declared by variable name type = value

	var firstname string = "simon"

	var secondname string = " " //undeclared variables are set to null

	//var b, f, s = true, 2.33, "lucky" //variables can also be declared in this way

	//othername := "Antonny" // with this you can declare your variables in a shorten name
	firstname, lastname := "simon", "awiti"

	//x, y, z := 1, 2, 3

	//sum := x + y + z

	const PI = 3.142 //once set can not be used anywhere else

	if num1 != 1 {
		fmt.Println("not true")
	} else if firstname != "simon" {
		fmt.Println("not true")

	} else if secondname != " " {
		fmt.Println("not true")

	} else if lastname != "awiti" {
		fmt.Println("not true")
	} else {
		fmt.Println("youre not the person werelooking for")
	}

	//fmt.Println("Getting more familia")
	//fmt.Println(num1)
	//fmt.Println(firstname)
	//fmt.Println(secondname)
	//fmt.Println(b, f, s)
	//fmt.Println(othername)
	//fmt.Println(sum)
	//fmt.Println(firstname + lastname)

}
