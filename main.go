package main

import "fmt"

func main() {
	//arrays
	var onlinestores [3]string = [3]string{"aws", "target", "alibaba"}
	var scores [3]int = [3]int{23, 22, 21}

	fmt.Println(scores)

	//another way to declare an array

	var ages = [3]int{12, 15, 17}

	fmt.Println(ages, len(ages), onlinestores)

	//to shorten it all

	trainers := [4]string{"Josh", "Simon", "ted", "Keziah"}

	fmt.Println(trainers, len(trainers))
}
