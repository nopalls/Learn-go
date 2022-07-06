package main

import "fmt"

func getHello(name string) string {

	return "Welcome " + name

}

func main() {

	sayHello := getHello
	fmt.Println(sayHello("Naufal"))
}

//func value adalah first class citizen
//func juga merupakan tipe data,dan bisa di simpan dalam variable
