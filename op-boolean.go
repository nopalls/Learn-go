package main

import "fmt"

func main() {
	var ujian = 80
	var absen = 75

	fmt.Println(ujian >= 70 && absen >= 60)
}

//operasi boolean
// && = dan
// || = atau
// | = kebalikan

//operasi &&
//(nilai1) true && (nilai2)true = true
//true && false = false
//false && true = false
//false && false = false

//opeasi || / or
//(nilai1) true || true = true
//true || false = true
//false || true = true
//false || false = false

//operasi !
//(operator !)true = false
//operator !)false = true
