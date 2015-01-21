package gslgo_test

import (
	"clustering/gsl"
	"clustering/math"
	"testing"
)

func TestInverse(t *testing.T) {
	m := cmath.RandomMatrix(10, 10)
	t.Logf("test matrix:\n %s", m.String())
	t.Log("testing Inverse...")
	inv := gslgo.Inverse(m)
	t.Logf("Inverse Matrix:\n %s", inv.String())
	t.Logf("m*m^{-1}:\n %s", m.Mul(inv).String())
}

func TestInverseDiag(t *testing.T) {
	m := cmath.RandomDiagMatrix(10)
	t.Logf("test matrix:\n %s", m.String())
	t.Log("testing InverseDiag...")
	inv := gslgo.InverseDiag(m)
	t.Logf("Inverse Matrix by gsl:\n %s", inv.String())
	t.Logf("Inverse Matrix:\n %s", inv.String())
	t.Logf("m*m^{-1}:\n %s", m.Mul(inv).String())
}

func TestInverseDiagSqrt(t *testing.T) {
	m := cmath.RandomDiagMatrix(10)
	t.Logf("test matrix:\n %s", m.String())
	t.Log("testing InverseDiagSqrt...")
	inv := gslgo.InverseDiagSqrt(m)
	t.Logf("Inverse Matrix by gsl:\n %s", inv.String())
	t.Logf("Inverse Matrix:\n %s", inv.String())
	t.Logf("m*Sqrt(m)^{-2}:\n %s", m.Mul(inv.Mul(inv)).String())
}
