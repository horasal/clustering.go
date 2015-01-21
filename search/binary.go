package search

import (
	"sort"
)

type binaryData struct {
	content []IBinarySearchableData
}

func (d *binaryData) Add(t IBinarySearchableData) {
	if d.Search(t) == -1 {
		d.content = append(d.content, t)
		sort.Sort(d)
	}
}
func (d binaryData) Search(t IBinarySearchableData) int { return BinarySearch(d.content, t) }
func (d binaryData) Value(i int) IBinarySearchableData  { return d.content[i] }
func (d binaryData) Len() int                           { return len(d.content) }
func (d binaryData) Less(a, b int) bool                 { return d.content[a].Equal(d.content[b]) > 0 }
func (d *binaryData) Swap(a, b int)                     { d.content[a], d.content[b] = d.content[b], d.content[a] }

func NewBinaryData() IBinarySearch {
	b := new(binaryData)
	b.content = make([]IBinarySearchableData, 0)
	return b
}

func BinarySearch(src []IBinarySearchableData, dst IBinarySearchableData) int {
	if len(src) == 0 {
		return -1
	}
	le, re := 0, len(src)-1
	for i := (re + le) / 2; re >= le; i = (re + le) / 2 {
		if c := src[i].Equal(dst); c == 0 {
			return i
		} else if c < 0 {
			re = i - 1
		} else {
			le = i + 1
		}
	}
	return -1
}
