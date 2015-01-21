package gslgo_test

import (
	"clustering/gsl"
	"clustering/math"
	"testing"
)

func TestEigenSymmv(t *testing.T) {
	m := cmath.RandomSummvMatrix(10, 10)
	t.Logf("test matrix:\n %s", m.String())
	t.Log("testing EigenSymmv...")
	v, m := gslgo.EigenSymmv(m)
	t.Logf("EigenValue: %s", v.String())
	t.Logf("EigenVectors:\n %s", m.String())
}
