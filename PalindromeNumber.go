package main

import (
	"fmt"
	"strconv"
)

//Given an integer x, return true if x is palindrome integer.
//
//An integer is a palindrome when it reads the same backward as forward. For example, 121 is palindrome while 123 is not.




func main() {


	input := 121

	returnValue := isPalindrome(input)

	fmt.Print(returnValue)
}


func isPalindrome(x int) bool {

	stringValue := strconv.Itoa(x)
	stringValue = reverse(stringValue)
	reversedInt, _ := strconv.Atoi(stringValue)

	return x == reversedInt
}


func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
