//Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
//
//Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
//
//
//
//Example 1:
//
//Input: x = 123
//Output: 321
//Example 2:
//
//Input: x = -123
//Output: -321
//Example 3:
//
//Input: x = 120
//Output: 21
//Example 4:
//
//Input: x = 0
//Output: 0

package main

import "fmt"

func main() {

	input := 120
	fmt.Println(reverse(input))

}

func reverse(x int) int {

	if x < 10 && x > 0 {
		return x
	}

	leftOver := x

	if x < 0 {
		leftOver = x * -1
	}

	isNotDone := true
	reversed := 0

	for ok := true; ok; ok = isNotDone {

		lastDigit := lastDigit(leftOver)
		leftOver = removeLastDigit(leftOver)
		reversed = addDigit(reversed, lastDigit)

		if leftOver < 10 {
			isNotDone = false
			reversed = addDigit(reversed, leftOver)
		}
	}

	//check -
	if x < 0 {
		reversed = reversed * -1
	}
	return reversed
}

func lastDigit(x int) int {
	return x % 10
}

func removeLastDigit(x int) int {
	number := x / 10
	return number
}

func addDigit(x int, digit int) int {

	return (x * 10) + digit
}
