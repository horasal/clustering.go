/*
  Clustering Analysis Package
  Math:Vector Unit
  Copyright (c) Zhai HongJie 2012
*/
package cmath

import (
	"fmt"
	"math"
)

type Vector struct {
	n int
	a []float64
}

func (a *Vector) At(i int) float64 {
	if !(0 <= i && i < a.n) {
		panic("index out of range")
	}
	return a.a[i]
}

func (a *Vector) Dim() int {
	return a.n
}

func (a *Vector) Set(i int, x float64) {
	if !(0 <= i && i < a.n) {
		panic("index out of range")
	}
	a.a[i] = x
}

func (a *Vector) Mul(b IVector) IVector {
	if a.n != b.Dim() {
		panic("illegal Vector multiply")
	}
	c := newVector(a.n)
	for i := 0; i < c.n; i++ {
		for j := 0; j < c.n; j++ {
			c.Set(i, a.At(i)*b.At(j))
		}
	}
	return c
}

func (a *Vector) Normalize() IVector {
	c := newVector(a.n)
	sum := 0.0
	for i := 0; i < a.n; i++ {
		sum += math.Pow(a.At(i), 2)
	}
	sum = math.Sqrt(sum)
	for i := 0; i < a.n; i++ {
		if sum == 0 {
			c.Set(i, 0)
		} else {
			c.Set(i, a.At(i)/sum)
		}
	}
	return c
}

func (a *Vector) Eql(b IVector) bool {
	if a.n != b.Dim() {
		return false
	}
	for i := 0; i < a.n; i++ {
		for j := 0; j < a.n; j++ {
			if a.At(i) != b.At(i) {
				return false
			}
		}
	}
	return true
}

func (a *Vector) String() string {
	s := ""
	for i := 0; i < a.n; i++ {
		s += fmt.Sprintf("%f,", a.At(i))
	}
	s += "\n"
	return s
}

func newVector(n int) *Vector {
	if !(0 <= n) {
		return nil
	}
	a := new(Vector)
	a.n = n
	a.a = make([]float64, n)
	return a
}

func NewVector(n int) IVector {
	return newVector(n)
}
