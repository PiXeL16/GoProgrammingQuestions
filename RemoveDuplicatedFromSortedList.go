package main

import (
	"GoProgrammingQuestions/utils"
	"fmt"
)

//
//Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func main() {

	nodes := buildList()

	// Original
	utils.printList(nodes)

	notDuplicated := deleteDuplicates(nodes)

	fmt.Println("Not Duplicated nodes")
	utils.printList(notDuplicated)

}

func deleteDuplicates(head *utils.ListNode) *utils.ListNode {

	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	notDuplicatedListNode := utils.ListNode{
		Val: head.Val,
	}

	deleteDuplicatesRec(head.Next, &notDuplicatedListNode)

	return &notDuplicatedListNode
}

func deleteDuplicatesRec(node *utils.ListNode, notDuplicatedNode *utils.ListNode) {

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
