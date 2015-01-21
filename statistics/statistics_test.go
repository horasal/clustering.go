package statistics_test

import (
	"clustering/math"
	"clustering/statistics"
	"testing"
)

func TestCosine(t *testing.T) {
	t.Log("test cosine...")
	a := cmath.RandomVector(10)
	b := cmath.RandomVector(10)
	t.Logf("Vector:\n%s\n%s", a.String(), b.String())
	t.Logf("Cosine value:%f", statistics.Cosine(a, b))
}

func TestAverage(t *testing.T) {
	t.Log("test Average...")
	a := cmath.RandomMatrix(15, 15)
	t.Logf("Matrix:\n%s", a.String())
	t.Logf("Average:\n%s", statistics.Average(a).String())
}

func TestVariance(t *testing.T) {
	t.Log("test Average...")
	a := cmath.RandomMatrix(15, 15)
	t.Logf("Matrix:\n%s", a.String())
	t.Logf("Variance:%f", statistics.Variance(a))
}
