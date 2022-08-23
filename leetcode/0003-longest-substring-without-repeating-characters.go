package leetcode

import (
	"bytes"
)

// Given that s consists of English letters, digits, symbols and spaces, we assume ISO-8859-1
const universe = 128

type presence [universe]bool

func (p *presence) set(ch byte) (exists bool) {
	exists = p[ch]
	p[ch] = true
	return
}

func (p *presence) clear(ch byte) {
	p[ch] = false
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	// Assumed ISO-8859-1 charset, we can give compiler a hint to ignore possible runes
	bs := []byte(s)

	longest := 0
	current := 0
	singleChars := presence{}

	for i, ch := range bs {
		if singleChars.set(ch) {
			start := i - current
			index := bytes.IndexByte(bs[start:], ch)
			// index is never negative, as we know the character is repeating
			for j := start; j < start+index; j++ {
				singleChars.clear(bs[j])
			}
			if current > longest {
				longest = current
			}
			current -= index
		} else {
			current++
		}
	}
	if current > longest {
		longest = current
	}

	return longest
}
