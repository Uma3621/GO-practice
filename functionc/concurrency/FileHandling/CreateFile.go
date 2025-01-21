package FileHandling

import (
	"fmt"
	"os"
)

func Create() {
	//create a file
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("file not create error", err)
		return
	}
	defer file.Close()
	//writing to a file
	_, err = file.WriteString("hello go writing file!\n")
	if err != nil {
		fmt.Println("error wtiting to a file", err)
		return
	}
	fmt.Println("file written successful")

}
