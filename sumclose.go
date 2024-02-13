package main

import (
	"fmt"
	"math"
	"sort"
)

func ทดสอบ () {
	fmt.Println("Hello, ทดสอบ")
	ทด  := 3
	fmt.Println(ทด)
}

func threeSumClosest(nums []int, target int) int {
	best := nums[0] + nums[1] + nums[2]
	for i := 0; i <= len(nums)-3; i++ {
		for j := i + 1; j <= len(nums)-2; j++ {
			for k := j + 1; k <= len(nums)-1; k++ {
				if (math.Abs(float64(nums[i] + nums[j] + nums[k] - target))) <= math.Abs(float64(best-target)) {
					fmt.Println(nums[i], " ", nums[j], " ", nums[k])
					fmt.Println(math.Abs(float64((nums[i] + nums[j] + nums[k] - target))), "<=", int(math.Abs(float64((best - target)))))
					best = nums[i] + nums[j] + nums[k]
					fmt.Println("best", best)
				}
			}
		}
	}
	return best
}




func threeSumClosestBest(nums []int, target int) int {
	sort.Ints(nums)
	closestSum := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1

		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]
			if abs(currentSum-target) < abs(closestSum-target) {
				closestSum = currentSum
			}
			if currentSum < target {
				left++
			} else if currentSum > target {
				right--
			} else {
				return currentSum
			}
		}
	}

	return closestSum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
