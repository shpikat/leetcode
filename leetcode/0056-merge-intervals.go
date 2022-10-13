package leetcode

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Sort(Intervals(intervals))
	merged := make([][]int, 0, len(intervals))
	current := intervals[0]
	merged = append(merged, current)
	for _, interval := range intervals[1:] {
		if current[1] >= interval[0] {
			if current[1] < interval[1] {
				current[1] = interval[1]
			}
		} else {
			current = interval
			merged = append(merged, current)
		}
	}
	return merged
}

type Intervals [][]int

func (is Intervals) Len() int           { return len(is) }
func (is Intervals) Less(i, j int) bool { return is[i][0] < is[j][0] }
func (is Intervals) Swap(i, j int)      { is[i], is[j] = is[j], is[i] }
