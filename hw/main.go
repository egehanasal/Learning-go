package main

import "fmt"

func main() {

	// Creating a slice named sl
	sl := []int{}
	// Adding integers in it from 0 to 10, by looping and using append function
	for i := 0; i <= 10; i++ {
		sl = append(sl, i)
	}
	// Looping through the function to access all the elements in it
	for i := range sl {
		// Checking if the number is even
		if sl[i]%2 == 0 {
			// Printing the number if it is even
			fmt.Println(sl[i], "is even")
		} else {
			// Printing the number if it is odd
			fmt.Println(sl[i], "is odd")
		}
	}
}