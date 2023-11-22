package main

import (
	"fmt"
	"strings"
)

func main() {
	ingredients := []string{"meat", "cooking oil", "flour"}

	newstatement := "hello chef"

	chef := "Simon"

	fmt.Println(ingredients, chef)

	ingredients = append(ingredients, "onions")

	fmt.Println(strings.Contains(newstatement, "onions"))
	fmt.Println(strings.ReplaceAll(newstatement, "hello", "hi"))

	statement := fmt.Sprintf("Chef %v used the %v ingredients to cook: %v ", chef, len(ingredients), ingredients)

	fmt.Println(statement)

}
