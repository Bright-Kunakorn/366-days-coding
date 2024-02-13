package main

func permute(nums []int) [][]int {
	var result [][]int
	generatePermutations(nums, 0, &result)
	return result
}

func generatePermutations(nums []int, index int, result *[][]int) {
	if index == len(nums)-1 {
		perm := make([]int, len(nums))
		copy(perm, nums)
		*result = append(*result, perm)
		return
	}

	for i := index; i < len(nums); i++ {
		nums[index], nums[i] = nums[i], nums[index]
		generatePermutations(nums, index+1, result)
		nums[index], nums[i] = nums[i], nums[index]
	}
}
