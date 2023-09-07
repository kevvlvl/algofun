package calc

import "fmt"

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
