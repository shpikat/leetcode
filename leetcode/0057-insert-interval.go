package leetcode

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {
	// intervals are given to be in order for starti, and as none of the intervals overlap--for endi, too
	// so we can use binary search to find the insertion point
	start := sort.Search(len(intervals), func(i int) bool { return newInterval[0] < intervals[i][0] })
	end := sort.Search(len(intervals), func(i int) bool { return newInterval[1] < intervals[i][1] })
	var newLength int
	if start == end && (start == 0 || intervals[start-1][1] < newInterval[0]) && (end == len(intervals) || intervals[end][0] > newInterval[1]) {
		// newInterval doesn't overlap with anything, can't be merged, adding as a new interval
		intervals = append(intervals, []int{})
		newLength = len(intervals)
	} else {
		// merge the intervals around the insertion point
		if start > 0 && intervals[start-1][1] >= newInterval[0] {
			start--
			newInterval[0] = intervals[start][0]
		}
		if end < len(intervals) && intervals[end][0] <= newInterval[1] {
			newInterval[1] = intervals[end][1]
			end++
		}
		newLength = start + 1 + len(intervals) - end
	}

	// shift the intervals left (if several intervals were merged) or right (if a new one was added)
	copy(intervals[start+1:], intervals[end:])
	intervals[start] = newInterval
	intervals = intervals[:newLength]

	return intervals
}
