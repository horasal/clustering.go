/*

A interface of GSL for google go language
GSL-Go:SVDDecomposit Unit
copyright (c) zhai hongjie 2012

GNU Scientific Library (GSL)
Copyright © 1996, 1997, 1998, 1999, 2000,
2001, 2002, 2003, 2004, 2005, 2006, 2007,
2008, 2009, 2010, 2011 The GSL Team.

*/
package gslgo

//#cgo LDFLAGS: -lgsl -lgslcblas
//#include <gsl/gsl_linalg.h>
//#include <gsl/gsl_math.h>
import "C"
import (
	"clustering/math"
)

func SVDDecompostion(m cmath.IMatrix) (cmath.IMatrix, cmath.IMatrix, cmath.IMatrix) {
	gsl_matrix := C.gsl_matrix_alloc(C.size_t(m.DimN()), C.size_t(m.DimM()))
	for i := 0; i < m.DimN(); i++ {
		for j := 0; j < m.DimM(); j++ {
			C.gsl_matrix_set(gsl_matrix, C.size_t(i), C.size_t(j), C.double(m.At(i, j)))
		}
	}
	evec := C.gsl_matrix_alloc(C.size_t(m.DimM()), C.size_t(m.DimM()))
	evel := C.gsl_vector_alloc(C.size_t(m.DimM()))
	w := C.gsl_vector_alloc(C.size_t(m.DimM()))

	C.gsl_linalg_SV_decomp(gsl_matrix, evec, evel, w)
	u := cmath.NewMatrix(m.DimN(), m.DimM())
	for j := 0; j < m.DimM(); j++ {
		evec_i := C.gsl_matrix_column(gsl_matrix, C.size_t(j))
		for i := 0; i < m.DimN(); i++ {
			u.Set(i, j, float64(C.gsl_vector_get(&evec_i.vector, C.size_t(i))))
		}
	}

	v := cmath.NewMatrix(m.DimM(), m.DimM())
	for i := 0; i < m.DimM(); i++ {
		evec_i := C.gsl_matrix_column(evec, C.size_t(i))
		for j := 0; j < m.DimM(); j++ {
			v.Set(j, i, float64(C.gsl_vector_get(&evec_i.vector, C.size_t(j))))
		}
	}
	s := cmath.NewMatrix(m.DimM(), m.DimM())
	for i := 0; i < m.DimM(); i++ {
		s.Set(i, i, float64(C.gsl_vector_get(evel, C.size_t(i))))
	}
	C.gsl_matrix_free(evec)
	C.gsl_matrix_free(gsl_matrix)
	C.gsl_vector_free(evel)
	C.gsl_vector_free(w)
	return u, s, v
}

func SVDDecompostionJacobi(m cmath.IMatrix) (cmath.IMatrix, cmath.IMatrix, cmath.IMatrix) {
	gsl_matrix := C.gsl_matrix_alloc(C.size_t(m.DimN()), C.size_t(m.DimM()))
	for i := 0; i < m.DimN(); i++ {
		for j := 0; j < m.DimM(); j++ {
			C.gsl_matrix_set(gsl_matrix, C.size_t(i), C.size_t(j), C.double(m.At(i, j)))
		}
	}
	evec := C.gsl_matrix_alloc(C.size_t(m.DimM()), C.size_t(m.DimM()))
	evel := C.gsl_vector_alloc(C.size_t(m.DimM()))

	C.gsl_linalg_SV_decomp_jacobi(gsl_matrix, evec, evel)

	u := cmath.NewMatrix(m.DimN(), m.DimM())
	for j := 0; j < m.DimM(); j++ {
		evec_i := C.gsl_matrix_column(gsl_matrix, C.size_t(j))
		for i := 0; i < m.DimN(); i++ {
			u.Set(i, j, float64(C.gsl_vector_get(&evec_i.vector, C.size_t(i))))
		}
	}

	v := cmath.NewMatrix(m.DimM(), m.DimM())
	for i := 0; i < m.DimM(); i++ {
		evec_i := C.gsl_matrix_column(evec, C.size_t(i))
		for j := 0; j < m.DimM(); j++ {
			v.Set(j, i, float64(C.gsl_vector_get(&evec_i.vector, C.size_t(j))))
		}
	}

	s := cmath.NewMatrix(m.DimM(), m.DimM())
	for i := 0; i < m.DimM(); i++ {
		s.Set(i, i, float64(C.gsl_vector_get(evel, C.size_t(i))))
	}
	C.gsl_matrix_free(evec)
	C.gsl_matrix_free(gsl_matrix)
	C.gsl_vector_free(evel)
	return u, s, v
}
