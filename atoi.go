package main

import (
	"strings"
)


func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	sign := 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	result := 0
	for _, char := range s {
		if char < '0' || char > '9' {
			break
		}
		digit := int(char - '0')
		result = result*10 + digit
		if result*sign > (1<<31 - 1) {
			return 1<<31 - 1
		} else if result*sign < -(1 << 31) {
			return -(1 << 31)
		}
	}

	return result * sign
}
