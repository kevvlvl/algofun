package lists

import (
	"fmt"
	"math/big"
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

	l1Number := big.NewInt(0)
	var i int64 = 0
	for l1 != nil {

		e := big.NewInt(0)
		e.Exp(big.NewInt(10), big.NewInt(i), nil)
		m := big.NewInt(0)
		m.Mul(big.NewInt(int64(l1.Val)), e)

		l1Number.Add(l1Number, m)

		fmt.Println("l1 = ", l1.Val, " - i = ", i, " - sum = ", l1Number)
		i++

		l1 = l1.Next
	}

	l2Number := big.NewInt(0)
	var j int64 = 0
	for l2 != nil {

		e := big.NewInt(0)
		e.Exp(big.NewInt(10), big.NewInt(j), nil)
		m := big.NewInt(0)
		m.Mul(big.NewInt(int64(l2.Val)), e)

		l2Number.Add(l2Number, m)

		fmt.Println("l2 = ", l2.Val, " - j = ", j, " - sum = ", l2Number)
		j++

		l2 = l2.Next
	}

	a := big.NewInt(0)
	a.Add(l1Number, l2Number)

	fmt.Println("Resulting number = ", a)

	rem := big.NewInt(0)
	rem.Mod(a, big.NewInt(10))

	rHead := &ListNode{Val: int(rem.Int64())}

	a.Div(a, big.NewInt(10))
	zero := big.NewInt(0)

	if a.Cmp(zero) > 0 {
		r := &ListNode{}
		rHead.Next = r

		for a.Cmp(zero) > 0 {

			rem = big.NewInt(0)
			rem.Mod(a, big.NewInt(10))

			r.Val = int(rem.Int64())
			a.Div(a, big.NewInt(10))

			if a.Cmp(zero) > 0 {
				r.Next = &ListNode{}
				r = r.Next
			}
		}
	}

	return rHead
}

func (l *ListNode) InsertEnd(v int) *ListNode {

	node := &ListNode{Val: v}
	l.Next = node

	return l.Next
}