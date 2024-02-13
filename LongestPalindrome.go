package main

func preProcess(s string) string {
	n := len(s)
	if n == 0 {
		return "^$"
	}

	result := "^"
	for i := 0; i < n; i++ {
		result += "#" + string(s[i])
	}
	result += "#$"
	return result
}

func longestPalindrome2(s string) string {
	T := preProcess(s)
	n := len(T)
	P := make([]int, n)
	C, R := 0, 0

	for i := 1; i < n-1; i++ {
		mirror := 2*C - i
		if R > i {
			P[i] = min(R-i, P[mirror])
		} else {
			P[i] = 0
		}

		for T[i+1+P[i]] == T[i-1-P[i]] {
			P[i]++
		}

		if i+P[i] > R {
			C = i
			R = i + P[i]
		}
	}

	maxLen, centerIndex := 0, 0
	for i := 1; i < n-1; i++ {
		if P[i] > maxLen {
			maxLen = P[i]
			centerIndex = i
		}
	}

	start := (centerIndex - 1 - maxLen) / 2
	return s[start : start+maxLen]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
