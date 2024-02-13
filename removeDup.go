package main

import (
	"sort"
)

func removeElement(nums []int, val int) int {
	sort.Ints(nums)
	ans := []int{}
	for _, e := range nums {
		if e != val {
			ans = append(ans, e)
		}
	}
	return len(ans)
}
