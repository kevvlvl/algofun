package main

import (
	"fmt"
	"leetcode_playground/calc"
	"leetcode_playground/strs"
)

func main() {

	nums1 := []int{2, 4, 9, 2, 1, 10, 29}

	fmt.Println(calc.TwoSum(nums1, 11))
	fmt.Println(calc.TwoSum(nums1, 39))
	fmt.Println(calc.TwoSum(nums1, 13))

	fmt.Println(strs.IsValid("()"))
	fmt.Println(strs.IsValid("(){}[]"))
	fmt.Println(strs.IsValid("((){}[])"))
	fmt.Println(strs.IsValid("(]"))
	fmt.Println(strs.IsValid("({{)"))
	fmt.Println(strs.IsValid("{[]}"))
}
