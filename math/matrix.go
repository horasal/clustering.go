/*
  Clustering Analysis Package
  Math:Matrix Unit
  Copyright (c) Zhai HongJie 2012
*/

package cmath

import (
	"fmt"
	"math"
)

type Matrix struct {
	n, m int
	a    []float64
}

func (a *Matrix) Vector(i int) IVector {
	v := newVector(a.m)
	for j := 0; j < a.m; j++ {
		v.Set(j, a.At(i, j))
	}
	return v
}

func (a *Matrix) VectorT(i int) IVector {
	v := newVector(a.m)
	for j := 0; j < a.m; j++ {
		v.Set(j, a.At(j, i))
	}
	return v
}

func (a *Matrix) At(i, j int) float64 {
	if !(0 <= i && i < a.n && 0 <= j && j < a.m) {
		panic("index out of range")
	}
	return a.a[i*a.m+j]
}

func (a *Matrix) DimN() int {
	return a.n
}

func (a *Matrix) DimM() int {
	return a.m
}

func (a *Matrix) Set(i, j int, x float64) {
	if !(0 <= i && i < a.n && 0 <= j && j < a.m) {
		panic("index out of range")
	}
	a.a[i*a.m+j] = x
}

func (a *Matrix) Transposition() IMatrix{
	m := newMatrix(a.DimM(),a.DimN())
	for i:=0;i<a.DimN();i++{
		for j:=0;j<a.DimM();j++{
			m.Set(j,i,a.At(i,j))
		}
	}
	return m
}

func (a *Matrix) NormalizeN() IMatrix {
	c := newMatrix(a.n, a.m)
	for i := 0; i < c.n; i++ {
		sum := 0.0
		for j := 0; j < a.m; j++ {
			sum += math.Pow(a.At(i, j), 2)
		}
		sum = math.Sqrt(sum)
		for j := 0; j < a.m; j++ {
			if sum == 0 {
				c.Set(i, j, 0)
			} else {
				c.Set(i, j, a.At(i, j)/sum)
			}
		}
	}
	return c
}
func (a *Matrix) NormalizeM() IMatrix {
	c := newMatrix(a.n, a.m)
	for i := 0; i < c.m; i++ {
		sum := 0.0
		for j := 0; j < a.n; j++ {
			sum += math.Pow(a.At(j, i), 2)
		}
		sum = math.Sqrt(sum)
		for j := 0; j < a.n; j++ {
			if sum == 0 {
				c.Set(j, i, 0)
			} else {
				c.Set(j, i, a.At(j, i)/sum)
			}
		}
	}
	return c
}

func (a *Matrix) MatrixT() IMatrix {
	c := newMatrix(a.m, a.n)
	for i := 0; i < c.m; i++ {
		for j := 0; j < a.n; j++ {
			c.Set(j, i, a.At(i, j))
		}
	}
	return c
}

func (a *Matrix) Mul(b IMatrix) IMatrix {
	if a.m != b.DimN() {
		panic("illegal Matrix multiply")
	}
	c := newMatrix(a.n, b.DimM())
	for i := 0; i < c.n; i++ {
		for j := 0; j < c.m; j++ {
			x := 0.0
			for k := 0; k < a.m; k++ {
				x += a.At(i, k) * b.At(k, j)
			}
			c.Set(i, j, x)
		}
	}
	return c
}

func (a *Matrix) Eql(b IMatrix) bool {
	if a.n != b.DimN() || a.m != b.DimM() {
		return false
	}
	for i := 0; i < a.n; i++ {
		for j := 0; j < a.m; j++ {
			if a.At(i, j) != b.At(i, j) {
				return false
			}
		}
	}
	return true
}

func (a *Matrix) String() string {
	s := ""
	for i := 0; i < a.n; i++ {
		for j := 0; j < a.m; j++ {
			s += fmt.Sprintf("%f,", a.At(i, j))
		}
		s += "\n"
	}
	return s
}

func newMatrix(n, m int) *Matrix {
	if !(0 <= n && 0 <= m) {
		return nil
	}
	a := new(Matrix)
	a.n = n
	a.m = m
	a.a = make([]float64, n*m)
	return a
}

func newUnit(n int) *Matrix {
	a := newMatrix(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x := 0.0
			if i == j {
				x = 1
			}
			a.Set(i, j, x)
		}
	}
	return a
}

func NewMatrix(n, m int) IMatrix {
	return newMatrix(n, m)
}

func NewUnit(n int) IMatrix {
	return newUnit(n)
}
