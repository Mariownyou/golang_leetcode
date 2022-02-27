package main

import "fmt"

func moveZeroes(A []int) {
	l := 0
	for r := 0; r < len(A); r++ {
		if A[r] > 0 {
			temp := A[r]
			A[r] = A[l]
			A[l] = temp
			l++
		}
	}
	fmt.Println(A)
}

func main() {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	moveZeroes([]int{0, 1, 0})
	moveZeroes([]int{0, 0, 1})
}
