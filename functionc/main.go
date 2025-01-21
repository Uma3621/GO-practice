package main

import (
	"fmt"
	"main/concurrency"
	"main/concurrency/FileHandling"
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
	s = "R"
	n = 43
	return
}
func main() {

	concurrency.SelectStatement()
	fmt.Println("File handling starts here...........")
	FileHandling.Create()
	FileHandling.Read()
	/*concurrency.Funcmain()
	fmt.Println("starting goroutines example...........")
	concurrency.DemonstrateConcurrency() //goroutines exp
	fmt.Println("ending goroutines example...........")*/

	/*//interfcaes
	dog := interfaces.Dog{Name: "Buddy"}
	cat := interfaces.Cat{Name: "Kitty"}

	interfaces.MakeAnimalSpeak(dog) // Output: Woof!
	interfaces.MakeAnimalSpeak(cat) // Output: Meow!
	structs.UsingStructs()

	modules.Vals()
	rangeover.Range()
	//multipleReturnValues() //callig multiplereturnvalues.go
	const name = "prathyusha"
	const num = 21 //we can declare and initialize conct without mentioning any type
	fmt.Println("this is a fucntion")
	fmt.Println(num, " ", name)
	fmt.Println(appname)
	fmt.Println(DATE_FORMAT)
	withparam("hi", 66)

	age := retuningASingleValue()
	fmt.Println("Age:", age)

	a, b := retuningMultiValues()
	fmt.Println(a, b)
	a, b = retuningMultiValues2()
	fmt.Println(a, b)*/
}
