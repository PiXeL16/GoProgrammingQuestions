package main

import "fmt"

func main() {

	sequence := []int{1,4,4}
	array := []int{1,2,3,4}

	valid := IsValidSubsequence(array, sequence)


	fmt.Print(valid)
}

func IsValidSubsequence(array []int, sequence []int) bool {

	foundCount := 0
	mark := 0
	for i:= 0 ; i< len(sequence); i++ {
		numberFromSequence := sequence[i]
		for y:= mark; y < len(array); y++ {
			numberFromArray := array[y]
			if numberFromSequence == numberFromArray {
				foundCount ++
				y ++
				mark = y
				break
			}
		}
	}

	if foundCount == len(sequence) {
		return true
	} else {
		return false
	}

}



