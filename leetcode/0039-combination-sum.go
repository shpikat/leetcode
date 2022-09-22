package leetcode

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	stack := CombinationSumStack{}
	stack.Push(CombinationSumEntry{[]int{}, 0, 0})

	for !stack.IsEmpty() {
		entry := stack.Pop()
		if entry.sum == target {
			result = append(result, entry.combination)
		} else {
			for i := entry.index; i < len(candidates); i++ {
				sum := entry.sum + candidates[i]
				if sum <= target {
					// We could have avoided multiple array copies by counting the indices,
					// but the problem states there are not going to be that many combinations
					// to make that difference substantial.
					var s []int
					s = append(s, entry.combination...)
					stack.Push(CombinationSumEntry{
						append(s, candidates[i]),
						sum,
						i,
					})
				}
			}
		}
	}

	return result
}

type CombinationSumStack []CombinationSumEntry

type CombinationSumEntry struct {
	combination []int
	sum         int
	index       int
}

func (s *CombinationSumStack) Push(v CombinationSumEntry) {
	*s = append(*s, v)
}

func (s *CombinationSumStack) Pop() CombinationSumEntry {
	last := len(*s) - 1
	t := (*s)[last]
	*s = (*s)[:last]
	return t
}

func (s CombinationSumStack) IsEmpty() bool {
	return len(s) == 0
}
