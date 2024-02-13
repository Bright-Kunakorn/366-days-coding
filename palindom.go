package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	temp := x
	y := 0
	for temp > 0 {
		y = (y + (temp % 10)) * 10
		temp = temp / 10

	}
	y = y / 10
	if y == x {
		return true
	} else {
		return false
	}
}


