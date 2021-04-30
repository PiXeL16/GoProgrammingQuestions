package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildListWithArray(array []int) *ListNode {

	headNode := ListNode{
		Val: array[0],
	}

	array = RemoveIndex(array, 0)

	buildListRecursively(&headNode, array)

	return &headNode
}

func buildListRecursively(node *ListNode, arr []int) {

	if len(arr) == 0 {
		return
	}

	newNode := ListNode{
		Val: arr[0],
	}

	arr = RemoveIndex(arr, 0)

	node.Next = &newNode

	buildListRecursively(node.Next, arr)
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func PrintList(node *ListNode) {

	fmt.Println(node.Val)

	if node.Next != nil {
		PrintList(node.Next)
	}

}
