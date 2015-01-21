package csort

import (
	"sort"
)

type sortableIndex struct {
	indexs []int
	value  SortableSource
}

func (q sortableIndex) Len() int { return len(q.indexs) }

func (q sortableIndex) Less(i, j int) bool {
	return q.value.At(q.indexs[i]) > q.value.At(q.indexs[j])
}

func (q *sortableIndex) Swap(i, j int) {
	q.indexs[i], q.indexs[j] = q.indexs[j], q.indexs[i]
}

func newSortableIndex(value SortableSource) *sortableIndex {
	s := new(sortableIndex)
	s.indexs = make([]int, value.Dim())
	s.value = value
	for i := 0; i < len(s.indexs); i++ {
		s.indexs[i] = i
	}
	return s
}

func Sort(v SortableSource) []int {
	s := newSortableIndex(v)
	sort.Sort(s)
	return s.indexs
}

func MaxN(v SortableSource, n int) []int {
	return Sort(v)[:n]
}
