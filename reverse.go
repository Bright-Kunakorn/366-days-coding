package main

import (
	"math"
)

func reverse(x int) int {
	const intMax = int(^uint32(0) >> 1)
	const intMin = ^intMax

	temps := int(math.Abs(float64(x)))
	ans := 0

	for temps > 0 {
		remainder := temps % 10
		ans = ans*10 + remainder
		if ans > intMax || ans < intMin {
			return 0
		}
		temps /= 10
	}
	if x < 0 {
		ans = -ans
	}
	return ans
}
