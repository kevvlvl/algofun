package lists

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	head := &ListNode{}
	tail := &ListNode{}

	// Start head and tail with the min value of both list
	if list1.Val <= list2.Val {
		head = list1
		tail = list1

		// Move to next node
		list1 = list1.Next
	} else {
		head = list2
		tail = list2

		// Move to next node
		list2 = list2.Next
	}

	for list1 != nil && list2 != nil {

		if list1.Val <= list2.Val {
			tail.Next = list1
			tail = list1

			// Move to next node
			list1 = list1.Next
		} else {
			tail.Next = list2
			tail = list2

			list2 = list2.Next
		}
	}

	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return head
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// TODO: convert to big int. See https://pkg.go.dev/math/big
	//var l1Big big.Int
	//var l2Big big.Int

	l1Number := float64(0)
	i := 0
	for l1 != nil {

		l1Number += float64(l1.Val) * math.Pow10(i)
		fmt.Println("l1 = ", l1.Val, " - i = ", i, " - sum = ", l1Number)
		i++

		l1 = l1.Next
	}

	l2Number := float64(0)
	j := 0
	for l2 != nil {

		l2Number += float64(l2.Val) * math.Pow10(j)
		fmt.Println("l2 = ", l2.Val, " - j = ", j, " - sum = ", l2Number)
		j++

		l2 = l2.Next
	}

	resultAddition := l1Number + l2Number
	fmt.Println("Resulting number = ", resultAddition)

	rHead := &ListNode{Val: int(math.Mod(resultAddition, 10))}
	resultAddition = float64(int(resultAddition / 10))

	if resultAddition > 0 {
		r := &ListNode{}
		rHead.Next = r

		for resultAddition > 0 {

			r.Val = int(math.Mod(resultAddition, 10))
			resultAddition = float64(int(resultAddition / 10))

			if resultAddition > 0 {
				r.Next = &ListNode{}
				r = r.Next
			}
		}
	}

	return rHead
}
