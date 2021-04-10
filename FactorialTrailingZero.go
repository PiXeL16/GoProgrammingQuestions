//Given an integer n, return the number of trailing zeroes in n!.
//
//Follow up: Could you write a solution that works in logarithmic time complexity?
//
//
//
//Example 1:
//
//Input: n = 3
//Output: 0
//Explanation: 3! = 6, no trailing zero.
//Example 2:
//
//Input: n = 5
//Output: 1
//Explanation: 5! = 120, one trailing zero.
//Example 3:
//
//Input: n = 0
//Output: 0

package main

import (
	"fmt"
	"math/big"
)

func main() {

	number := 30
	fmt.Println(trailingZeroes(number))
}

func trailingZeroes(n int) int {

	var fact = new(big.Int)
	fact.MulRange(1, int64(n))

	trailingCeros := 0

	hasCero := false

	for ok := true; ok; ok = hasCero {
		if lastDigit(factorial) == 0 {
			trailingCeros++
			hasCero = true
			factorial = removeLastDigit(factorial)

		} else {
			hasCero = false
		}

	}

	return trailingCeros

}

func lastDigit(n big.Int) uint64 {
	var lastDigit big.Int
	lastDigit.Mod(big.NewInt(n), big.NewInt(10))
	return lastDigit
}

func removeLastDigit(n uint64) uint64 {
	removedDigit := n / 10
	return removedDigit
}

//func calculateFactorial(n int) uint64 {
//
//	var factorial uint64
//	factorial = 1
//	for x :=1; x<=n; x++ {
//		factorial = factorial * uint64(x)
//	}
//
//	return factorial
//}
