package main

import (
	"fmt"
)

const (
	DATE_FORMAT = "DD/MM/YY" // it's exported as starting with upper case
	appname     = "uma"      // it's only inside this package as it's starting with lower case
)

func main() {
	const name = "prathyusha"
	const num = 21 //we can declare and initialize conct without mentioning any type
	fmt.Println("this is a fucntion")
	fmt.Println(appname)
	fmt.Println(DATE_FORMAT)
}
