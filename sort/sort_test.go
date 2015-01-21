package csort_test

import (
	"clustering/math"
	"clustering/sort"
	"fmt"
	"testing"
)

func TestNewVector(t *testing.T) {
	t.Log("test Sort...")
	m := cmath.RandomVector(50)
	t.Logf("Vector:\n %s", m.String())
	sorted := csort.Sort(m)
	s := ""
	for i := 0; i < len(sorted); i++ {
		s += fmt.Sprintf("%f,", m.At(sorted[i]))
	}
	t.Logf("Sorted: \n %s", s)

}
