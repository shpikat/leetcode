package leetcode

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	counts := make([]int, len(s)+1)
	counts[0] = 1
	if s[0] != '0' {
		counts[1] = 1
	}
	for i := 2; i <= len(s); i++ {
		if s[i-1] != '0' {
			counts[i] += counts[i-1]
		}
		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6') {
			counts[i] += counts[i-2]
		}
	}
	return counts[len(s)]
}
