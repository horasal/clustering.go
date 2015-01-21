package statistics

import (
	"clustering/math"
	"math"
)

func Cosine(a, b cmath.IVector) float64 {
	if cmath.VectorDot(a, b) == 0 {
		return 0
	}
	return cmath.VectorDot(a, b) / (cmath.VectorDet(a) * cmath.VectorDet(b))
}

func Average(data cmath.IMatrix) cmath.IVector {
	c := cmath.NewVector(data.DimM())
	for i := 0; i < data.DimN(); i++ {
		for j := 0; j < data.DimM(); j++ {
			c.Set(j, c.At(j)+data.At(i, j))
		}
	}
	for i := 0; i < c.Dim(); i++ {
		c.Set(i, c.At(i)/float64(data.DimN()))
	}
	return c
}

func Variance(data cmath.IMatrix) float64 {
	if data.DimN() == 0 {
		return 0
	}
	center := Average(data)
	v := 0.0
	for i := 0; i < data.DimN(); i++ {
		v += math.Pow(cmath.EuclidDistance(center, data.Vector(i)), 2)
	}
	return v / float64(data.DimN())
}
