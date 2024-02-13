package main

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var current []int

	backtrack(candidates, target, 0, current, &result)

	return result
}

func backtrack(candidates []int, target, start int, current []int, result *[][]int) {
	if target == 0 {
		*result = append(*result, append([]int{}, current...))
		return
	}

	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			continue
		}

		current = append(current, candidates[i])

		backtrack(candidates, target-candidates[i], i, current, result)

		current = current[:len(current)-1]
	}
}
