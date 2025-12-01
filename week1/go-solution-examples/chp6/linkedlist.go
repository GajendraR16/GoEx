package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Helper: Create linked list from slice
func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	curr := head

	for i := 1; i < len(vals); i++ {
		curr.Next = &ListNode{Val: vals[i]}
		curr = curr.Next
	}

	return head
}

// Helper: Print linked list
func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Print(curr.Val)
		if curr.Next != nil {
			fmt.Print(" → ")
		}
		curr = curr.Next
	}
	fmt.Println()
}

func reversed(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func recReversed(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := recReversed(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
		list1 = list1.Next
		tail = tail.Next
	}

	if list2 != nil {
		tail.Next = list2
		list2 = list2.Next
		tail = tail.Next
	}

	return dummy.Next
}

func recMergeTwoList(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	// Recursive case
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func main() {
	list1 := createList([]int{1, 2, 4})
	fmt.Print("Original l1: ")
	printList(list1)

	list2 := createList([]int{1, 3, 5})
	fmt.Print("Original l2: ")
	printList(list2)

	// Reverse
	list1A := createList([]int{1, 2, 3, 4, 5})
	reversed := reversed(list1A)
	fmt.Print("Reversed: ")
	printList(reversed)

	// Recursive Reverse
	list1B := createList([]int{5, 4, 3, 2, 1})
	recreversed := recReversed(list1B)
	fmt.Print("Recrusive Reversed: ")
	printList(recreversed)

	// Merged
	merged := mergeTwoLists(list1, list2)
	fmt.Print("Merged: ")
	printList(merged)

	// Recursive Merged
	l1 := createList([]int{1, 2, 4})
	l2 := createList([]int{1, 3, 5})
	recmerged := recMergeTwoList(l1, l2)
	fmt.Print("Recursive Merged: ")
	printList(recmerged)
}
