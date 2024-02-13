package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	var res [][]int
	sort.Ints(nums)
	if nums[0] == nums[len(nums)-1] {
		if nums[0]+nums[1]+nums[2]+nums[3] == target {
			res = append(res, []int{nums[0], nums[1], nums[2], nums[3]})
		} else {
			res = [][]int{}
		}
		return res
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				for l := k + 1; l < len(nums); l++ {
					if nums[i]+nums[j]+nums[k]+nums[l] == target {
						temp := []int{nums[i], nums[j], nums[k], nums[l]}
						sort.Ints(temp)
						if stringInSlice(temp, res) == false {
							res = append(res, temp)
						}
					}
				}
			}
		}
	}

	return res
}

func stringInSlice(a []int, list [][]int) bool {
	for _, b := range list {
		if len(a) == len(b) {
			equal := true
			for i, v := range a {
				if v != b[i] {
					equal = false
					break
				}
			}
			if equal {
				return true
			}
		}
	}
	return false
}

func fourSum2(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left, right := j+1, len(nums)-1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})

					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}

					left++
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return result
}
