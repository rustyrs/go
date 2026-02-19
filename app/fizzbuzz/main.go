package main

import "fmt"

func main() {
	for n := 1; n <= 20; n++ {
		if n % 15 == 0 {
			fmt.Println("FizzBuzz")
		} else if n % 5 == 0 {
			fmt.Println("Buzz")
		} else if n % 3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(n)
		}
	}

	for n := 1; n <= 20; n++ {
		switch {
		case n % 15 == 0:
			fmt.Println("FizzBuzz")
		case n % 5 == 0:
			fmt.Println("Buzz")
		case n % 3 == 0:
			fmt.Println("Fizz")
		default:
		    fmt.Println(n)
		}
	}

	for n := 1; n <= 20; n++ {
		result := ""

		if n%3 == 0 {
			result += "Fizz"
		}
		if n%5 == 0 {
			result += "Buzz"
		}

		if result == "" {
			fmt.Println(n)
		} else {
			fmt.Println(result)
		}
	}
}