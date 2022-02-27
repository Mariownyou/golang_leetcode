package main

import "fmt"

func search(A []int, target int) int {
	l := 0
	r := len(A) - 1

	for l <= r {
		m := l + ((r - l) / 2)
		val := A[m]
		if val == target {
			return m
		}
		if val > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

func main() {
	res := search([]int{-1, 0, 3, 5, 9, 12}, 9)
	fmt.Println(res)

	test_1 := search([]int{-1, 0, 3, 5, 9, 12}, 2)
	fmt.Println(test_1)
}
