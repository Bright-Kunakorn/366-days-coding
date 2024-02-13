package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	least := len(strs[0])
	leastValue := strs[0]
	if len(strs) == 0 || len(strs) == 1 {
		return strs[0]
	}
	for _, k := range strs {
		if len(k) < least {
			least = len(k)
			leastValue = k
		}
	}
	if least == 0 {
		return ""
	}
	if least == 1 {
		for _, l := range strs {
			if l[0:1] != leastValue {
				fmt.Print(l[0:1], "!=", leastValue)
				return ""
			}
		}
		return leastValue
	}

	if strs[0] == "" || strs[0][0] != strs[1][0] {
		return ""
	}

	for j := 1; j <= least; j++ {
		for _, k := range strs {
			if k[0:j] != strs[0][0:j] {
				return strs[0][0 : len(k[0:j])-1]
			}
		}

	}
	return leastValue
}

func longestCommonPrefixVer(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLength := len(strs[0])
	for _, str := range strs {
		if len(str) < minLength {
			minLength = len(str)
		}
	}
	for i := 0; i < minLength; i++ {
		currentChar := strs[0][i]
		for j := 1; j < len(strs); j++ {
			fmt.Println("j", j, "i", i)
			if strs[j][i] != currentChar {
				return strs[0][:i]
			}
		}
	}

	return strs[0][:minLength]
}




