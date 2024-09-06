package main

import (
	"fmt"
	"leetcode_playground/calc"
	"leetcode_playground/lists"
	"leetcode_playground/strs"
)

func main() {

	nums1 := []int{2, 4, 9, 2, 1, 10, 29}

	fmt.Println("------------------------ Find Numbers that sum the target")
	fmt.Println(calc.TwoSum(nums1, 11))
	fmt.Println(calc.TwoSum(nums1, 39))
	fmt.Println(calc.TwoSum(nums1, 13))

	fmt.Println("------------------------ Valid Parenthesis")
	fmt.Println(strs.IsValid("()"))
	fmt.Println(strs.IsValid("(){}[]"))
	fmt.Println(strs.IsValid("((){}[])"))
	fmt.Println(strs.IsValid("(]"))
	fmt.Println(strs.IsValid("({{)"))
	fmt.Println(strs.IsValid("{[]}"))

	fmt.Println("------------------------ Palindromes")
	fmt.Println(calc.IsPalindrome(101))
	fmt.Println(calc.IsPalindrome(50))
	fmt.Println(calc.IsPalindrome(55))
	fmt.Println(calc.IsPalindrome(225522))
	fmt.Println(calc.IsPalindrome(-50))

	fmt.Println("------------------------ Roman numerals converter")
	calc.RomanToInteger("III")
	calc.RomanToInteger("LVIII")
	calc.RomanToInteger("MCMXCIV")

	fmt.Println("------------------------ Longest common prefix")
	fmt.Println("Prefix = ", strs.LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println("Prefix = ", strs.LongestCommonPrefix([]string{"dog", "racecar", "car"}))

	fmt.Println("------------------------ Merge Two Sorted Lists")

	n1 := lists.ListNode{
		Val: 1,
		Next: &lists.ListNode{
			Val:  2,
			Next: &lists.ListNode{Val: 40},
		},
	}

	n2 := lists.ListNode{
		Val: 15,
		Next: &lists.ListNode{
			Val:  40,
			Next: &lists.ListNode{Val: 45},
		},
	}

	merged := lists.MergeTwoLists(&n1, &n2)
	for merged != nil {
		fmt.Println("Var: ", merged.Val)
		merged = merged.Next
	}

	fmt.Println("------------------------ Add Two Numbers")

	a1 := lists.ListNode{Val: 5}
	n := a1.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(0)
	n = n.InsertEnd(1)

	a2 := lists.ListNode{Val: 1}
	a2.InsertEnd(5)

	added := lists.AddTwoNumbers(&a1, &a2)
	for added != nil {
		fmt.Print("Addition: ", added.Val)
		added = added.Next

		if added != nil {
			fmt.Print("->")
		}
	}

	fmt.Println("------------------------ Remove duplicates from Sorted Array")

	numsdups := []int{1, 1, 2, 2}
	size := lists.RemoveDuplicates(numsdups)

	fmt.Println("Unique amount of digits: ", size, " - Array: ", numsdups)

	fmt.Println("------------------------ Remove Element from Array")

	//nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	nums := []int{3, 2, 2, 3}
	//nums := []int{2}
	kept := lists.RemoveElement(nums, 2)

	fmt.Println("Amount of digits kept: ", kept, " - Array: ", nums)

	fmt.Println("------------------------ Find Index of First occurrence of str")

	idx := strs.GetIndexForStr("Good news everyone!", "Good")
	fmt.Println("Found str at index: ", idx)

	fmt.Println("------------------------ Search insert position in array")

	idx = lists.SearchInsert([]int{1, 2, 3, 4}, 3)
	fmt.Println("Ideal index: ", idx)

	fmt.Println("------------------------ Length of Last Word")

	s := "Hello my   darling     vampire     "
	l := strs.LengthOfLastWord(s)
	fmt.Println("Length of last word: ", l)
}
