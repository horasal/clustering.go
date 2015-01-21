package cmath

import (
	"math/rand"
	"time"
)

func init() {
	randomize()
}

func randomize() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RandInt(i int) int {
	return rand.Intn(i)
}

func RandFloat() float64 {
	return rand.Float64()
}

func RandomMatrix(n, m int) IMatrix {
	matrix := NewMatrix(n, m)
	for i := 0; i < matrix.DimN(); i++ {
		for j := 0; j < matrix.DimM(); j++ {
			matrix.Set(i, j, rand.Float64())
		}
	}
	return matrix
}

func RandomVector(n int) IVector {
	m := NewVector(n)
	for i := 0; i < m.Dim(); i++ {
		m.Set(i, rand.Float64())
	}
	return m
}

func RandomSummvMatrix(n, m int) IMatrix {
	matrix := NewMatrix(n, m)
	for i := 0; i < matrix.DimN(); i++ {
		for j := 0; j < matrix.DimM(); j++ {
			if i > j {
				matrix.Set(i, j, matrix.At(j, i))
			} else {
				matrix.Set(i, j, rand.Float64())
			}
		}
	}
	return matrix
}

func RandomDiagMatrix(n int) IMatrix {
	m := NewMatrix(n, n)
	for i := 0; i < m.DimN(); i++ {
		for j := 0; j < m.DimM(); j++ {
			if i == j {
				m.Set(i, j, rand.Float64())
			}
		}
	}
	return m
}
