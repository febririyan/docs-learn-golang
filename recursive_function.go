package main

import "fmt"

func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursiveLoop(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursiveLoop(value-1)
	}
}

func main() {
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursiveLoop(10))
}
