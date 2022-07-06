package main

import "fmt"

func main() {

	counter := 0 // this is variable counter

	increment := func() { //this is anonymous function
		fmt.Println("increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}
