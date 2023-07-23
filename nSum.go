package main

import "fmt"

func nSum(nums []int, target int, n int) [][]int {
	var result [][]int
	if n < 2 || len(nums) < n {
		return result
	}
	if n == 2 {
		numMap := make(map[int]int)

		for i, num := range nums {
			complement := target - num
			if index, found := numMap[complement]; found {
				result = append(result, []int{nums[index], num})
			}
			numMap[num] = i
		}
	} else {
		for i := 0; i < len(nums)-n+1; i++ {
			subResult := nSum(nums[i+1:], target-nums[i], n-1)
			for _, sr := range subResult {
				combination := append([]int{nums[i]}, sr...)
				result = append(result, combination)
			}
		}
	}
	return result
}

func main() {
	nums := []int{2, 7, 4, 0, 9, 5, 1, 3}
	target := 14
	n := 4
	result := nSum(nums, target, n)
	fmt.Println(result)
}
