package main

import (
	"fmt"
)

const (
	DATE_FORMAT = "DD/MM/YY" // it's exported as starting with upper case
	appname     = "uma"      // it's only inside this package as it's starting with lower case
)

func withparam(p1 string, p2 int) {
	fmt.Println("with parameters ::"+p1, p2)
}
func main() {
	const name = "prathyusha"
	const num = 21 //we can declare and initialize conct without mentioning any type
	fmt.Println("this is a fucntion")
	fmt.Println(appname)
	fmt.Println(DATE_FORMAT)
	withparam("hi", 66)
}
