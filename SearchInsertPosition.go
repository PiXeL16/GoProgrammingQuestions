package main

//Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

import (
	"fmt"
)

func main() {

	nums := []int{1, 3, 5, 6}
	target := 7

	fmt.Println(searchInsertBinary(nums, target))
}

func searchInsert(nums []int, target int) int {

	for x := 0; x < len(nums); x++ {

		if nums[x] == target {
			return x
		} else if nums[x] > target {
			return x
		}
	}
	return len(nums)
}

func searchInsertBinary(nums []int, target int) int {

	low := 0
	high := len(nums)

	for low < high {

		mid := (low + high) / 2

		if nums[mid] >= target {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}
