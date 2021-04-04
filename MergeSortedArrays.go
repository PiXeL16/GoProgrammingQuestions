package main

import (
	"fmt"
	"sort"
)

//Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
//
//The number of elements initialized in nums1 and nums2 are m and n respectively. You may assume that nums1 has a size equal to m + n such that it has enough space to hold additional elements from nums2.

func main() {

	nums1 := []int{0, 0, 0, 0, 0}
	m := 0
	nums2 := []int{1, 2, 3, 4, 5}
	n := 5

	merge(nums1, m, nums2, n)

	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	if n == 0 {
		return
	}

	y := 0
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[y]
		y++
	}

	sort.Ints(nums1)
}
