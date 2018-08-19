package main

// LoopSearch - Binary search of value in an array by using a loop
func LoopSearch(integers []int, value int) bool {
	limInf := 0
	limSup := len(integers)

	for limInf < limSup {
		midIndex := limInf + (limSup-limInf)/2
		if value == integers[midIndex] {
			return true
		} else if value < integers[midIndex] {
			limSup = midIndex - 1
		} else {
			limInf = midIndex + 1
		}
	}
	return false
}

// RecursiveSearch - Binary search of value in an array by using recursion
func RecursiveSearch(integers []int, limInf int, limSup int, value int) bool {

	midIndex := (limSup + limInf) / 2

	if limInf > limSup {
		return false
	}

	if integers[midIndex] == value {
		return true
	} else if value < integers[midIndex] {
		return RecursiveSearch(integers, limInf, midIndex-1, value)
	} else {
		return RecursiveSearch(integers, midIndex+1, limSup, value)
	}
}
