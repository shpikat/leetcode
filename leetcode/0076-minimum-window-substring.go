package leetcode

import "math"

func minWindow(s string, t string) string {
	// Uppercase and lowercase English letters fit into 7-bit ASCII
	letters := [128]int{}
	for _, ch := range t {
		letters[ch]++
	}
	required := len(t)
	minStart, minLength := 0, math.MaxInt32

	for slow, fast := 0, 0; fast < len(s); fast++ {
		ch1 := s[fast]
		letters[ch1]--
		if letters[ch1] >= 0 {
			required--
		}
		for required == 0 {
			length := fast + 1 - slow
			if length < minLength {
				minStart, minLength = slow, length
			}
			ch2 := s[slow]
			letters[ch2]++
			if letters[ch2] > 0 {
				required++
			}
			slow++
		}
	}
	if minLength == math.MaxInt32 {
		return ""
	}
	return s[minStart : minStart+minLength]
}
