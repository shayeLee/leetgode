package main

import "fmt"

func twosum(nums []int, target int) []int {
	var num1, num2 int
OuterLoop:
	for i := 0; i < len(nums); i++ {
		for k := 0; k < len(nums); k++ {
			if i != k && nums[i]+nums[k] == target {
				num1 = i
				num2 = k
				break OuterLoop
			}
		}
	}
	return []int{num1, num2}
}

func main() {
	var nums = []int{3, 2, 4}
	fmt.Println(twosum(nums, 6))
}
