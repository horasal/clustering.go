package search_test

import (
	"clustering/math"
	"clustering/search"
	"fmt"
	"testing"
)

type testint struct {
	i int
}

func (i testint) Equal(a search.IBinarySearchableData) int {
	return a.(testint).i - i.i

}

func newInt(i int) search.IBinarySearchableData {
	return testint{i}
}

func TestBinarySearch(t *testing.T) {
	t.Log("test BinarySearch...")
	v := make([]search.IBinarySearchableData, 100)
	for i := -50; i < 50; i++ {
		v[i+50] = newInt(i)
	}
	t.Logf("test data:\n %v", v)
	t.Logf("search:55, result:%d", search.BinarySearch(v, testint{55}))
	t.Logf("search:3, result:%d", search.BinarySearch(v, testint{5}))
	t.Logf("search:-15, result:%d", search.BinarySearch(v, testint{-15}))
	t.Logf("search:0, result:%d", search.BinarySearch(v, testint{0}))
	t.Logf("search:-50, result:%d", search.BinarySearch(v, testint{-50}))
	t.Logf("search:49, result:%d", search.BinarySearch(v, testint{49}))
}

func TestNewBinaryData(t *testing.T) {
	t.Log("test BinaryData...")
	d := search.NewBinaryData()
	s := "Original Data:"
	for i := 0; i < 50; i++ {
		a := cmath.RandInt(100) - 50
		d.Add(newInt(a))
		s += fmt.Sprintf("%d,", a)
	}
	t.Log(s)
	s = "DataInTable:"
	for i := 0; i < d.Len(); i++ {
		s += fmt.Sprintf("%d,", d.Value(i).(testint).i)
	}
	t.Log(s)
	for i := 0; i < 50; i++ {
		a := cmath.RandInt(100) - 50
		t.Logf("search:%d,result:%d", a, d.Search(newInt(a)))
	}

}
