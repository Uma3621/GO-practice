package rangeover

import "fmt"

func Range() {

	// Here we use `range` to sum the numbers in a slice.
	// Arrays work like this too.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {

		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range on map iterates over key/value pairs.
	//Initializes a map named kvs with string keys and string values.

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
