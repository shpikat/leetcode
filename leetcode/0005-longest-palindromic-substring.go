package leetcode

func longestPalindrome(s string) string {
	n := len(s)

	if n <= 1 {
		return s
	}

	bs := []byte(s)

	start, end, best := 0, 0, 0

	// Basically implementation of Manacher's algorithm
	odd := make([]int, n)
	left, right := 0, 0
	for c := range bs {
		r := 0
		if c < right {
			r = min(right-c, odd[left+right-c])
		}
		for c-r > 0 && c+r < n-1 && bs[c-r-1] == bs[c+r+1] {
			r++
		}
		odd[c] = r
		if c+r > right {
			left, right = c-r, c+r
		}

		if r > best {
			start, end, best = c-r, c+r, r
		}
	}

	// Manacher's algorithm adapted for substrings of even length
	even := make([]int, n)
	left, right = 0, 0
	for c := range bs {
		r := 0
		if c < right {
			r = min(right-c, even[left+right-c+1])
		}
		for c-r > 0 && c+r < n && bs[c-r-1] == bs[c+r] {
			r++
		}
		even[c] = r
		if c+r-1 > right {
			left, right = c-r, c+r-1
		}

		if r > best {
			start, end, best = c-r, c+r-1, r
		}
	}

	return s[start : end+1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
