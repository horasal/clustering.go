/*

  Clustering Analysis Package
  Math:Interface Unit
  Copyright (c) Zhai HongJie 2012

*/

package cmath

type IMatrix interface {
	At(i, j int) float64
	Set(i, j int, x float64)
	Mul(b IMatrix) IMatrix
	Eql(b IMatrix) bool
	DimN() int
	DimM() int
	String() string
	Vector(i int) IVector
	VectorT(i int) IVector
	NormalizeN() IMatrix
	NormalizeM() IMatrix
	MatrixT() IMatrix
	Transposition() IMatrix
}

type IVector interface {
	At(i int) float64
	Set(i int, x float64)
	Mul(b IVector) IVector
	Eql(b IVector) bool
	Dim() int
	String() string
	Normalize() IVector
}
