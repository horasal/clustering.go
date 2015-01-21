package gslgo_test

import (
	"clustering/gsl"
	"clustering/math"
	"testing"
)

var (
	m = cmath.RandomMatrix(15, 10)
)

func TestSVDDecompostion(t *testing.T) {
	t.Log("test svd decomposition...")
	t.Logf("Raw Matrix:\n%s", m.String())
	u, s, v := gslgo.SVDDecompostion(m)
	t.Logf("U:\n%s", u.String())
	t.Logf("S:\n%s", s.String())
	t.Logf("V:\n%s", v.String())
	t.Logf("USV^T:\n%s", u.Mul(s).Mul(v.MatrixT()).String())
}

func TestSVDDecompostionJacobi(t *testing.T) {
	t.Log("test jacobi svd decomposition...")
	t.Logf("Raw Matrix:\n%s", m.String())
	u, s, v := gslgo.SVDDecompostionJacobi(m)
	t.Logf("U:\n%s", u.String())
	t.Logf("S:\n%s", s.String())
	t.Logf("V:\n%s", v.String())
	vt := cmath.NewMatrix(10, 10)
	for i := 0; i < vt.DimN(); i++ {
		for j := 0; j < vt.DimM(); j++ {
			vt.Set(i, j, v.At(j, i))
		}
	}
	t.Logf("USV^T:\n%s", u.Mul(s).Mul(v.MatrixT()).String())
}
