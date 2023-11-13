package main

import "fmt"

func main() {
	var age int

	fmt.Println("Enter your age: ")
	fmt.Scanf("%d", &age)

	if age == 28 {
		fmt.Println("success!")
	} else {
		fmt.Println("verification failed")
	}
}
