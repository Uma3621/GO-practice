package FileHandling

import (
	"fmt"
	"os"
)

func Read() {

	// Open the file in read-only mode
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a buffer to hold the file content
	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the content
	fmt.Println("File content:", string(buf[:n]))
}
