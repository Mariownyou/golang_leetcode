package main

import "fmt"

func main() {
	res := firstBadVersion(1)
	fmt.Println(res)
}

func firstBadVersion(n int) int {
	l := 0
	r := n
	for l <= r {
		mid := l + ((r - l) / 2)
		if l == r {
			return l
		}
		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return -1
}

func isBadVersion(version int) bool {
	if version >= 1 {
		return true
	}
	return false
}
