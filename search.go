package main

func searchInsert(nums []int, target int) int {
	if target == 0 && nums[0] >= 0 {
		return 0
	}
	
	if len(nums) == 1 {
		if nums[0] >= target {
			return 0
		} else {
			return 1
		}
	}
	if nums[0] >= target {
		return 0
	}
	for i := 0; i < len(nums)-1; i++ {
		if target > nums[i] && target <= nums[i+1] {
			return i + 1
		}
	}
	return len(nums)
}
