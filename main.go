package main

import (
	"fmt"
	"leetcode_playground/calc"
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
}
