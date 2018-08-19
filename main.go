package main

import (
	"flag"
	"fmt"
)

func main() {

	value := flag.Int("val", 0, "Insert Integer")
	binarySearhType := flag.String("type", "loop", "loop or recursive")
	flag.Parse()

	myArrayOfIntegers := []int{1, 3, 12, 16, 34, 45, 67, 71, 80, 81, 87, 99}

	if *binarySearhType == "loop" {
		if LoopSearch(myArrayOfIntegers, *value) {
			fmt.Println("Loop Binary search: Value found")
		} else {
			fmt.Printf("Loop Binary search: Value %d not found\n", *value)
		}
	} else if *binarySearhType == "recursive" {
		if RecursiveSearch(myArrayOfIntegers, 0, len(myArrayOfIntegers)-1, *value) {
			fmt.Println("Recursive Binary search: Value found")
		} else {
			fmt.Printf("Recursive Binary search: Value %d not found\n", *value)
		}
	}

}
