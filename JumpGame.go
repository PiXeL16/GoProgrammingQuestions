//Given an array of non-negative integers nums, you are initially positioned at the first index of the array.
//
//Each element in the array represents your maximum jump length at that position.
//
//Determine if you are able to reach the last index.
//
//
//
//Example 1:
//
//Input: nums = [2,3,1,1,4]
//Output: true
//Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
//Example 2:
//
//Input: nums = [3,2,1,0,4]
//Output: false
//Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.

package main

import "fmt"

func main() {

	input := []int{3, 2, 1, 0, 4}

	fmt.Println(canJump(input))
}

func canJump(nums []int) bool {

	if len(nums) == 0 {
		return true
	}

	if len(nums) == 1 {
		return true
	}

	for x := 0; x < len(nums)-1; x++ {

		value := nums[x]

		if value == 0 {
			return false
		}

		if (len(nums)-1)-(value+x) <= 0 {
			return true
		}

	}

	return false
}
