package main

import "fmt"

func main() {
	fmt.Println("Hello Golang")

	fmt.Println("Sum", Sum(4))
}

func Sum(number int) int {
	sum := 0
	for i := 0; i < number; i++ {
		sum += i
	}

	return sum
}
