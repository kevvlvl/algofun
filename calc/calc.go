package calc

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// TwoSum returns the indexes of the two numbers that give us the target summed value
func TwoSum(nums []int, target int) []int {

	for i, vali := range nums {
		for j, valj := range nums {

			if i != j {

				if vali+valj == target {
					fmt.Printf("Sum of %v and %v gives target %v\n", vali, valj, target)
					return []int{i, j}
				}
			}
		}
	}

	return []int{}
}

// IsPalindrome returns true if x reads as a palindrome
func IsPalindrome(x int) bool {

	s := strconv.Itoa(x)
	l := len(s) - 1

	sb := []byte(s)
	sbr := []byte(s)

	for i := 0; i <= l; i++ {
		sbr[i] = sb[l-i]
	}

	return bytes.Equal(sb, sbr)
}

// RomanToInteger return the int value of the roman numeral string
func RomanToInteger(r string) int {

	value := 0

	for i := 0; i < len(r); i++ {
		if string(r[i]) == "M" {
			value += 1000
		}

		if string(r[i]) == "D" {
			value += 500
		}

		if string(r[i]) == "C" {
			value += 100
		}

		if string(r[i]) == "L" {
			value += 50
		}

		if string(r[i]) == "X" {
			value += 10
		}

		if string(r[i]) == "V" {
			value += 5
		}

		if string(r[i]) == "I" {
			value += 1
		}
	}

	if strings.Index(r, "IV") > -1 {
		value -= 2
	}

	if strings.Index(r, "IX") > -1 {
		value -= 2
	}

	if strings.Index(r, "XL") > -1 {
		value -= 20
	}

	if strings.Index(r, "XC") > -1 {
		value -= 20
	}

	if strings.Index(r, "CD") > -1 {
		value -= 200
	}

	if strings.Index(r, "CM") > -1 {
		value -= 200
	}

	fmt.Printf("Value of %v is %v\n", r, value)

	return value
}
