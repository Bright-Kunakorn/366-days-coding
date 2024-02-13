package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	lastSpace := strings.LastIndex(s, " ")

	if lastSpace == -1 {
		return len(s)
	}

	return len(s) - lastSpace - 1
}
