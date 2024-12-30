package main

import (
	"fmt"
	"main/modules"
)

const (
	DATE_FORMAT = "DD/MM/YY" // it's exported as starting with upper case
	appname     = "uma"      // it's only inside this package as it's starting with lower case
)

func withparam(p1 string, p2 int) { //with multiple parameter
	fmt.Println("with parameters ::"+p1, p2)
}

// age
func retuningASingleValue() int {
	return 43
}

// name and age
func retuningMultiValues() (string, int) {
	return "Rao", 43
}

// name and age
func retuningMultiValues2() (s string, n int) {
	s = "Rao"
	n = 43
	return
}
func main() {
	modules.Vals()
	multipleReturnValues() //callig multiplereturnvalues.go
	const name = "prathyusha"
	const num = 21 //we can declare and initialize conct without mentioning any type
	fmt.Println("this is a fucntion")
	fmt.Println(appname)
	fmt.Println(DATE_FORMAT)
	withparam("hi", 66)

	age := retuningASingleValue()
	fmt.Println("Age:", age)

	a, b := retuningMultiValues()
	fmt.Println(a, b)
	a, b = retuningMultiValues2()
	fmt.Println(a, b)
}
