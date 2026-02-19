package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "apple", "orange", "banana", "apple", "grape"}
	fruitsMap:= map[string]int{}

	for _, fruit := range fruits {
		fruitsMap[fruit] += 1
	}

	for fruit, count := range fruitsMap {
		fmt.Println(fruit, count)
	}
}