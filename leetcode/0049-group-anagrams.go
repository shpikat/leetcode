package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {
	groupedByLetters := map[string][]string{}

	for _, str := range strs {
		bs := Bytes(str)
		sort.Sort(bs)
		groupedByLetters[string(bs)] = append(groupedByLetters[string(bs)], str)
	}

	anagrams := make([][]string, 0, len(groupedByLetters))
	for _, group := range groupedByLetters {
		anagrams = append(anagrams, group)
	}

	return anagrams
}

type Bytes []byte

func (b Bytes) Len() int           { return len(b) }
func (b Bytes) Less(i, j int) bool { return b[i] < b[j] }
func (b Bytes) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
