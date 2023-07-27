func reverse(x int) int {
	const INT_MAX = int(^uint32(0) >> 1)
	const INT_MIN = ^INT_MAX

	temps := int(math.Abs(float64(x)))
	ans := 0

	for temps > 0 {
		remainder := temps % 10
		ans = ans*10 + remainder
		if ans > INT_MAX || ans < INT_MIN {
			return 0
		}
		temps /= 10
	}
	if x < 0 {
		ans = -ans
	}
	return ans
}
