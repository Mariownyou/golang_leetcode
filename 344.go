func reverseString(s []byte)  {
	left := 0
	right := len(s) - 1
	for {
		if left >= right {
			break
		}
		l_el := s[left]
		s[left] = s[right]
		s[right] = l_el
		left++
		right--
	}
}