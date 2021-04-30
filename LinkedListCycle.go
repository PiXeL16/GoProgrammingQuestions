package main

import (
	"GoProgrammingQuestions/utils"
	"fmt"
)

func main() {

	exampleArr := []int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5}

	list := utils.BuildListWithArray(exampleArr)

	utils.PrintList(list)

	fmt.Println(hasCycle(list))

}

func hasCycle(head *utils.ListNode) bool {

	var memoryArr []int

	if head == nil {
		return false
	}

	notDone := true

	var node *utils.ListNode
	node = head

	for ok := true; ok; ok = notDone {

		if !contains(node.Val, memoryArr) {

			if node.Next == nil {
				notDone = false
			}

			memoryArr = append(memoryArr, node.Val)

			if node.Next == nil {
				return false
			} else {
				node = node.Next
			}
		} else {
			return true
		}

	}

	return false
}

func contains(value int, arr []int) bool {

	for _, number := range arr {

		if number == value {

			fmt.Printf("Duplicate is %v\n", number)
			return true
		}

	}
	return false

}
