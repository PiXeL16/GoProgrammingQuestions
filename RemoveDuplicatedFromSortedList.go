package main

import "fmt"

//
//Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	nodes := buildList()

	// Original
	printTree(nodes)

	notDuplicated := deleteDuplicates(nodes)

	fmt.Println("Not Duplicated nodes")
	printTree(notDuplicated)

}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	notDuplicatedListNode := ListNode{
		Val: head.Val,
	}

	deleteDuplicatesRec(head.Next, &notDuplicatedListNode)

	return &notDuplicatedListNode
}

func deleteDuplicatesRec(node *ListNode, notDuplicatedNode *ListNode) {

	if notDuplicatedNode.Val != node.Val {
		notDuplicatedNode.Next = node
	} else {
		notDuplicatedNode.Next = nil
	}

	if node.Next != nil {

		if notDuplicatedNode.Next != nil {
			deleteDuplicatesRec(node.Next, notDuplicatedNode.Next)
		} else {
			deleteDuplicatesRec(node.Next, notDuplicatedNode)
		}
	}
}

func buildList() *ListNode {

	exampleArr := []int{1, 2, 3, 4, 5, 5, 6}

	headNode := ListNode{
		Val: 1,
	}

	buildTreeRecursively(&headNode, exampleArr, 0)

	return &headNode

}

func buildTreeRecursively(node *ListNode, arr []int, currentPosition int) {

	if currentPosition > len(arr)-1 {
		return
	}

	newNode := ListNode{
		Val: arr[currentPosition],
	}

	node.Next = &newNode

	buildTreeRecursively(node.Next, arr, currentPosition+1)
}

func printTree(node *ListNode) {

	fmt.Println(node.Val)

	if node.Next != nil {
		printTree(node.Next)
	}

}
