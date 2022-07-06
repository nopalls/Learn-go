package main

import "fmt"

func startLog() {
	fmt.Println("Program Finished")
}

func ProccessLog() {
	defer startLog()
	fmt.Println("Program Starting")
}

func main() {
	ProccessLog()
}
