package main

import "fmt"

//Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.
//
//
//
//Example 1:
//
//Input: num = 38
//Output: 2
//Explanation: The process is
//38 --> 3 + 8 --> 11
//11 --> 1 + 1 --> 2
//Since 2 has only one digit, return it.

func main() {

	num := 19
	fmt.Println(addDigits(num))
}

func addDigits(num int) int {

	if num < 10 {

		return num
	}

	nums := splitNumber(num)
	t := true
	sum := 0
	if len(nums) > 1 {
		for ok := true; ok; ok = t {

			sum = reduce(nums)

			nums = splitNumber(sum)

			if len(nums) == 1 {
				t = false
			}
		}
	}
	return sum
}

func reduce(nums []int) int {

	sum := 0
	for x := 0; x < len(nums); x++ {

		sum = sum + nums[x]
	}

	return sum

}

func splitNumber(num int) []int {
	var nums []int
	splittedNumber := num

	t := true

	for ok := true; ok; ok = t {

		if splittedNumber >= 10 {

			lastDigit := splittedNumber % 10
			nums = append(nums, lastDigit)
			splittedNumber = splittedNumber / 10

		} else {
			nums = append(nums, splittedNumber)
			t = false
		}
	}

	return nums
}
