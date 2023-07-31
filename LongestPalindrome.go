package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("baagg"))
}

func reverseString(s string) string {
	bytes := []byte(s)
	left := 0
	right := len(bytes) - 1
	for left < right {
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left++
		right--
	}
	return string(bytes)
}

func longestPalindrome(s string) string {
	word := ""
	temps := ""
	front := ""
	back := ""
	max := 0

	for i := 0; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			temps = (s[j:i])
			if len(temps)%2 == 0 {
				front = temps[:len(temps)/2]
				back = temps[len(temps)/2:]
			} else {
				front = temps[:len(temps)/2]
				back = temps[len(temps)/2+1:]
			}
			if front == reverseString(back) {
				if len(temps) > max {
					max = len(temps)
					word = temps
				}
			}
		}
	}

	return word
}
