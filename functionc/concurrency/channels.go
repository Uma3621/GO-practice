package concurrency

import "fmt"

func squareAnumber(num int, ch chan int) {
	result := num * num
	ch <- result //send data to channel
}

func cubeAnumber(num int, ch chan int) {
	result2 := num * num * num
	ch <- result2 //send data to channel
}

func Funcmain() {
	//create channels for squar and cube
	squarechannel := make(chan int)
	cubechannel := make(chan int)

	//create goroutines for these channels
	go squareAnumber(2, squarechannel)
	go cubeAnumber(2, cubechannel)

	//receive result from channels
	squareresult := <-squarechannel
	cuberesult := <-cubechannel

	//print output
	fmt.Println("channels started here\n", squareresult)
	fmt.Println(cuberesult)
}
