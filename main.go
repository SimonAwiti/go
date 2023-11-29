package main

import (
	"fmt"
)

func main() {
	nbateams := []string{"Boston", "Lakers", "Suns", "GSW", "Dallas"}

	for i := 0; i < len(nbateams); i++ {
		fmt.Println("The team is : ", nbateams[i])
	}

	for index, value := range nbateams {
		fmt.Printf("The team in index %v is : %v \n", index, value)
	}
}
