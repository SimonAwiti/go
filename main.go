package main

import "fmt"

func main() {
	//arrays together with formated variables

	onlinestores := [3]string{"amazon", "target", "alibaba"}

	fmt.Printf("they like shopping at these %v online stores: %v \n", len(onlinestores), onlinestores)

	statement := fmt.Sprintf("they like shopping at these %v online stores: %v \n", len(onlinestores), onlinestores)

	fmt.Printf("The sister said that %v \n", statement)

}
