package main

import "fmt"

func endApp() {
	message := recover()
	fmt.Println("Error with message", message)

	fmt.Println("Application Finished")
}

func startApp(error bool) {
	defer endApp()
	if error {
		panic("Application Error")
	}

	fmt.Println("Application Starting")
}

func main() {
	startApp(true)
}

//panic func adlaah func yang kita bisa gunakan untuk memberhentikan program
//panic func biasanya di panggil ketika terjadi error pada saat program kita berjalan

//recover adlaah func yang kita bisa gunakan untuk menangkap data panic
//dengan recover,proses panic kita akan berhenti,sehingga program akan tetap berjalan
