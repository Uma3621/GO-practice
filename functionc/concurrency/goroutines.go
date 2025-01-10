package concurrency

import (
	"fmt"
	"time"
)

// Function that prints numbers from 1 to 5
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second) // Simulating work
	}
}

// Function that prints letters from A to E
func printLetters() {
	for i := 'A'; i <= 'E'; i++ {
		fmt.Printf("%c\n", i)
		time.Sleep(1 * time.Second) // Simulating work
	}
}

func DemonstrateConcurrency() {
	// Start two goroutines to execute functions concurrently
	go printNumbers()
	go printLetters()

	// Sleep for 6 seconds to allow goroutines to complete
	// The main function should wait for these goroutines to finish their work
	time.Sleep(6 * time.Second)
}
