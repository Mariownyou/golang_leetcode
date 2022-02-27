package main

import "fmt"

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	last := m + n - 1

	for m > 0 && n > 0 {
		val1 := nums1[m-1]
		val2 := nums2[n-1]
		if val1 > val2 {
			nums1[last] = val1
			m--
		} else {
			nums1[last] = val2
			n--
		}
		last--
	}

	for n > 0 {
		nums1[last] = nums2[n-1]
		n--
		last--
	}
	fmt.Println(nums1)
}
