package main

import "fmt"

// Menggunakan ini harus hati hati karena dapat merubah data dari counter nya

func main() {
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	increment()

	fmt.Println(counter)
}
