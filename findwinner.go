package main

func Count(n, k int) int {
	Arr := make([]int, n+1)

	for i := 0; i < n; i++ {
		Arr[i] = i + 1
	}

	curr := k - 1

	for n > 1 {
		k = Arr[curr]

		for i := curr; i < n; i++ {
			Arr[i] = Arr[i+1]
		}

		curr = ((curr - 1) + k) % (n - 1)
		n--
	}

	return Arr[0]
}



