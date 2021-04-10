package main

import "fmt"

//Given an array nums of size n, return the majority element.
//
//The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

//Example 1:
//
//Input: nums = [3,2,3]
//Output: 3
//Example 2:
//
//Input: nums = [2,2,1,1,1,2,2]
//Output: 2

func main() {

	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	m := make(map[int]int)

	for _, value := range nums {
		_, exists := m[value]
		if !exists {
			m[value] = 0
		} else {
			m[value] = m[value] + 1
		}
	}

	topValue := 0
	topKey := 0
	for key, value := range m {
		if value > topValue {
			topValue = value
			topKey = key
		}

	}

	return topKey
}
