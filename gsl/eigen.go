/*

A interface of GSL for google go language
GSL-Go:Eigen Unit
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
import "C"
import (
	"clustering/math"
	"unsafe"
)

func EigenSymmv(m cmath.IMatrix) (cmath.IVector, cmath.IMatrix) {
	tempm := make([]float64, m.DimM()*m.DimN())
	for i := 0; i < m.DimN(); i++ {
		for j := 0; j < m.DimM(); j++ {
			tempm[i*m.DimN()+j] = m.At(i, j)
		}
	}
	gsl_matrix_view := C.gsl_matrix_view_array((*C.double)(unsafe.Pointer(&tempm[0])), C.size_t(m.DimN()), C.size_t(m.DimM()))
	w := C.gsl_eigen_symmv_alloc(C.size_t(m.DimM()))
	eval := C.gsl_vector_alloc(C.size_t(m.DimM()))
	evec := C.gsl_matrix_alloc(C.size_t(m.DimN()), C.size_t(m.DimM()))
	C.gsl_eigen_symmv(&gsl_matrix_view.matrix, eval, evec, w)
	C.gsl_eigen_symmv_free(w)
	C.gsl_eigen_symmv_sort(eval, evec, C.GSL_EIGEN_SORT_VAL_DESC)
	eigenvalue := cmath.NewVector(m.DimN())
	eigenvector := cmath.NewMatrix(m.DimN(), m.DimM())
	for i := 0; i < m.DimN(); i++ {
		eigenvalue.Set(i, float64(C.gsl_vector_get(eval, C.size_t(i))))
		evec_i := C.gsl_matrix_column(evec, C.size_t(i))
		for j := 0; j < m.DimM(); j++ {
			eigenvector.Set(i, j, float64(C.gsl_vector_get(&evec_i.vector, C.size_t(j))))
		}
	}
	C.gsl_vector_free(eval)
	C.gsl_matrix_free(evec)
	return eigenvalue, eigenvector
}
