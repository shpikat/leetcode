package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	// Given that s consists of English letters, digits, symbols and spaces, we assume ISO-8859-1
	const universe = 128
	indices := [universe]int{}

	longest := 0
	current := 0
	substringStart := 0

	// Assumed ISO-8859-1 charset, we can give compiler a hint to ignore possible runes
	bs := []byte(s)
	for i, ch := range bs {
		repeatedIndex := indices[ch] - 1
		indices[ch] = i + 1
		if repeatedIndex < 0 {
			current++
		} else {
			for j := substringStart; j < repeatedIndex; j++ {
				indices[bs[j]] = 0
			}
			if current > longest {
				longest = current
			}
			current -= repeatedIndex - substringStart
			substringStart = repeatedIndex + 1
		}
	}
	if current > longest {
		longest = current
	}

	return longest
}
