package concurrency

import (
	"fmt"
	"time"
)

func SelectStatement() {
	//creating two channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	//launching goroutines
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "channel 1 is sent msg"
	}() //this is used for call anonymous functions

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "channel 2 is sent msg"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received from c1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from c2:", msg2)
	}

}
