/*

A interface of GSL for google go language
GSL-Go:Inverse Unit
copyright (c) zhai hongjie 2012

GNU Scientific Library (GSL)
Copyright © 1996, 1997, 1998, 1999, 2000,
2001, 2002, 2003, 2004, 2005, 2006, 2007,
2008, 2009, 2010, 2011 The GSL Team.

*/
package gslgo

//#cgo LDFLAGS: -lgsl -lgslcblas
//#include <gsl/gsl_eigen.h>
//#include <gsl/gsl_math.h>
//#include <gsl/gsl_linalg.h>
//#include <gsl/gsl_permutation.h>
import "C"
import "unsafe"
import "math"
import "clustering/math"

func Inverse(m cmath.IMatrix) cmath.IMatrix {
	tempm := make([]float64, m.DimM()*m.DimN())
	for i := 0; i < m.DimN(); i++ {
		for j := 0; j < m.DimM(); j++ {
			tempm[i*m.DimN()+j] = m.At(i, j)
		}
	}
	gsl_matrix_view := C.gsl_matrix_view_array((*C.double)(unsafe.Pointer(&tempm[0])), C.size_t(m.DimN()), C.size_t(m.DimM()))
	evec := C.gsl_matrix_alloc(C.size_t(m.DimN()), C.size_t(m.DimM()))
	p := C.gsl_permutation_alloc(C.size_t(m.DimM()))
	s := 0
	C.gsl_linalg_LU_decomp(&gsl_matrix_view.matrix, p, (*C.int)(unsafe.Pointer(&s)))
	C.gsl_linalg_LU_invert(&gsl_matrix_view.matrix, p, evec)
	r := cmath.NewMatrix(m.DimN(), m.DimM())
	for j := 0; j < m.DimM(); j++ {
		evec_i := C.gsl_matrix_column(evec, C.size_t(j))
		for i := 0; i < m.DimN(); i++ {
			r.Set(i, j, float64(C.gsl_vector_get(&evec_i.vector, C.size_t(i))))
		}
	}
	C.gsl_matrix_free(evec)
	return r
}

func InverseDiag(m cmath.IMatrix) cmath.IMatrix {
	r := cmath.NewMatrix(m.DimN(), m.DimM())
	for i := 0; i < m.DimN() && i < m.DimM(); i++ {
		r.Set(i, i, 1/m.At(i, i))
	}
	return r
}

func InverseDiagSqrt(m cmath.IMatrix) cmath.IMatrix {
	r := cmath.NewMatrix(m.DimN(), m.DimM())
	for i := 0; i < m.DimN() && i < m.DimM(); i++ {
		r.Set(i, i, math.Pow(1/m.At(i, i), 0.5))
	}
	return r
}
