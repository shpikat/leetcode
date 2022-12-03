package leetcode

func longestConsecutive(nums []int) int {
	m := make(map[int]bool, len(nums))
	for _, n := range nums {
		m[n] = false
	}

	max := 0
	for n := range m {
		// ensure it's the start of the streak
		_, exists := m[n-1]
		if !exists {
			streak := 0
			hasMore := true
			for hasMore {
				streak++
				_, hasMore = m[n+streak]
			}

			if streak > max {
				max = streak
			}
		}
	}

	return max
}
