package main

import "fmt"

func exam(width int, number int, pos [][]int) int {
	area := make([][]int, width)
	count := 0
	for i := range area {
		area[i] = make([]int, width)
	}

	for i := 0; i <= len(pos)-1; i++ {
		for j := 0; j <= len(area)-1; j++ {
			for k := 0; k <= len(area)-1; k++ {
				area[pos[i][0]][k] = 1
				area[k][pos[i][1]] = 1
			}
		}
	}

	for i := 0; i <= len(area)-1; i++ {
		for j := 0; j <= len(area[0])-1; j++ {
			if area[i][j] == 0 {
				count++
			}
		}
	}
	fmt.Println(area)
	return count
}
