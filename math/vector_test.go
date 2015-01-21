package cmath_test

import (
	"clustering/math"
	"testing"
)

func TestNewVector(t *testing.T) {
	t.Log("test NewVector...")
	m := cmath.RandomVector(10)
	n := cmath.RandomVector(10)
	t.Logf("Vector1:\n %s", m.String())
	t.Logf("Vector1 equal Vector1: %v", m.Eql(m))
	t.Logf("Vector2:\n %s", n.String())
	t.Logf("Vector2 equal Vector2: %v", n.Eql(n))
	t.Logf("Vector1*Vector2:\n %s", m.Mul(n).String())

	t.Logf("Vector1 Normalize:\n %s", m.Normalize().String())
}
