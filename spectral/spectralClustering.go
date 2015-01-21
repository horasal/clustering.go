/*

  Clustering Analysis Package
  SpectralClustering:Spectral Unit
  Copyright (c) Zhai HongJie 2012

*/

package spectral

import (
	"clustering/gsl"
	"clustering/math"
	"clustering/sort"
	"math"
)

func GaussianKernel(a cmath.IMatrix) cmath.IMatrix {
	m := cmath.NewMatrix(a.DimN(), a.DimM())
	for i := 0; i < a.DimN(); i++ {
		for j := 0; j < a.DimM(); j++ {
			m.Set(i, j, math.Exp(-math.Pow(a.At(i, j), 2)/math.Pow(Sigma, 2)))
		}
	}
	return m
}

func LinearKernel(a cmath.IMatrix) cmath.IMatrix {
	m := cmath.NewMatrix(a.DimN(), a.DimM())
	for i := 0; i < a.DimN(); i++ {
		for j := 0; j < a.DimM(); j++ {
			m.Set(i, j, a.At(i, j))
		}
	}
	return m
}

func PolynomialKernel(a cmath.IMatrix) cmath.IMatrix {
	m := cmath.NewMatrix(a.DimN(), a.DimM())
	for i := 0; i < a.DimN(); i++ {
		for j := 0; j < a.DimM(); j++ {
			m.Set(i, j, math.Pow(a.At(i, j)+1, 2))
		}
	}
	return m
}

func NASGaussianKernel(a cmath.IMatrix) cmath.IMatrix {
	v := buildSigmaMatrix(a)
	m := cmath.NewMatrix(a.DimN(), a.DimM())
	for i := 0; i < a.DimN(); i++ {
		for j := 0; j < a.DimM(); j++ {
			m.Set(i, j, math.Exp(-math.Pow(a.At(i, j), 2)/v.At(i)*v.At(j)))
		}
	}
	return m
}

func buildSigmaMatrix(matrix cmath.IMatrix) cmath.IVector {
	m := cmath.NewVector(matrix.DimN())
	for i := 0; i < matrix.DimN(); i++ {
		va := matrix.Vector(i)
		c := csort.MaxN(va, NAS)
		count := 0.0
		for _, v := range c {
			count += va.At(v)
		}
		m.Set(i, math.Abs(count/float64(NAS)))
	}
	return m
}

func matrixA(matrix cmath.IMatrix) cmath.IMatrix {
	m := cmath.NewMatrix(matrix.DimN(), matrix.DimM())
	for i := 0; i < matrix.DimN(); i++ {
		for j := 0; j < matrix.DimM(); j++ {
			if i == j {
				m.Set(i, j, 0)
			} else {
				m.Set(i, j, matrix.At(i,j))
			}
		}
	}
	return m
}

func matrixD(matrix cmath.IMatrix) cmath.IMatrix {
	m := cmath.NewMatrix(matrix.DimN(), matrix.DimM())
	if matrix.DimN() < matrix.DimM() {
		return m
	}
	for i := 0; i < matrix.DimN(); i++ {
		s := 0.0
		for j := 0; j < matrix.DimM(); j++ {
			s += matrix.At(i, j)
		}
		m.Set(i, i, s)
	}
	return m
}

func Laplacian(matrix cmath.IMatrix) cmath.IMatrix {
	A := matrixA(matrix)
	DI := gslgo.InverseDiagSqrt(matrixD(matrix))
	return DI.Mul(A).Mul(DI)
}

func LpEigen(matrix cmath.IMatrix) (cmath.IVector, cmath.IMatrix) {
	return gslgo.EigenSymmv(matrix)
}
