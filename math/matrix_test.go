package cmath_test

import (
	"clustering/math"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	t.Log("test NewMatrix...")
	m := cmath.RandomMatrix(10, 11)
	n := cmath.RandomMatrix(11, 12)
	t.Logf("Matrix1:\n %s", m.String())
	t.Logf("Matrix1 equal Matrix1: %v", m.Eql(m))
	t.Logf("Matrix2:\n %s", n.String())
	t.Logf("Matrix2 equal Matrix2: %v", n.Eql(n))
	t.Logf("Matrix1*Matrix2:\n %s", m.Mul(n).String())

	t.Logf("Matrix1 NormalizeM:\n %s", m.NormalizeM().String())
	t.Logf("Matrix1 NormalizeN:\n %s", m.NormalizeN().String())
}
