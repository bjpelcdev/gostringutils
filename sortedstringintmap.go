// Package main provides sorted map
package gostringutils

import "sort"

type SortedMap struct {
	m map[string]int
	s []string
}

func (sm *SortedMap) Len() int {
	return len(sm.m)
}

func (sm *SortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func (sm *SortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func SortedKeys(m map[string]int) []string {
	sm := new(SortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}
