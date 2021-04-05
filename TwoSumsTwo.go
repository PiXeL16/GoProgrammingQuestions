//Given an array of integers numbers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.
//
//Return the indices of the two numbers (1-indexed) as an integer array answer of size 2, where 1 <= answer[0] < answer[1] <= numbers.length.
//
//You may assume that each input would have exactly one solution and you may not use the same element twice.
//
//
//
//Example 1:
//
//Input: numbers = [2,7,11,15], target = 9
//Output: [1,2]
//Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.

package main

import "fmt"

func main() {

	numbers := []int{2, 3, 4}
	target := 6

	fmt.Println(twoSums(numbers, target))

}

func twoSums(numbers []int, target int) []int {

	result := make([]int, 2)

	for x := 0; x < len(numbers); x++ {
		for y := x + 1; y < len(numbers); y++ {
			if numbers[x]+numbers[y] == target {
				result[0] = x + 1
				result[1] = y + 1
				return result
			}
		}
	}
	return result
}
