package cmath

import (
	"math"
)

func EuclidDistance(a, b IVector) float64 {
	if a.Dim() != b.Dim() {
		return 0
	}
	f := 0.0
	for i := 0; i < a.Dim(); i++ {
		f += math.Pow(a.At(i)-b.At(i), 2)
	}
	return math.Sqrt(f)
}

func VectorDet(a IVector) float64 {
	sum := 0.0
	for i := 0; i < a.Dim(); i++ {
		sum += math.Pow(a.At(i), 2)
	}
	return math.Sqrt(sum)
}

func VectorDot(a, b IVector) float64 {
	if a.Dim() != b.Dim() {
		return 0
	}
	c := 0.0
	for i := 0; i < a.Dim(); i++ {
		c += a.At(i) * b.At(i)
	}
	return c
}
