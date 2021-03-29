package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 2
	fmt.Println(binarySearch(nums, target))
}

func binarySearch(nums []int, target int) int {

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
