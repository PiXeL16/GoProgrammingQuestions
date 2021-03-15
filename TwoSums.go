package main

//Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//You can return the answer in any order.


import "fmt"

func main() {

	numbers := []int {3,2,4}

	target := 6

	returnValue := twoSum(numbers, target)

	fmt.Print(returnValue)
}

func twoSum(nums []int, target int) []int {


	for i := 0; i < len(nums); i++ {

		for x := i + 1; x < len(nums); x++ {

			if nums[x]+nums[i] == target {

				return []int{i, x}
			}
		}
	}
	return []int{0, 0}
}
