package leetcode

import "unicode"

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		l := rune(s[left])
		for !unicode.IsLetter(l) && !unicode.IsDigit(l) {
			left++
			// no overflow from the left side
			if left == right {
				return true
			}
			l = rune(s[left])
		}

		r := rune(s[right])
		for !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			right--
			// no additional overflow check, as left is guaranteed to exist
			r = rune(s[right])
		}

		if unicode.ToLower(l) == unicode.ToLower(r) {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}
