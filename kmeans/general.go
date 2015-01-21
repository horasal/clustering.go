/*

  Clustering Analysis Package
  Kmeans:GeneralSupport Unit
  Copyright (c) Zhai HongJie 2012

*/
package kmeans

import (
	"clustering/math"
	"math"
)

var (
	Maxloop   int     = 20
	Threshold float64 = 0.001
	Variance  float64 = 0.8
	KMeansPP  bool    = false
)

func checkenvirenment() {
	if Maxloop < 1 {
		Maxloop = 1
	}
	if Threshold < 0 {
		Threshold = 0.001
	}
	if Variance <= 0 {
		Variance = 0.8
	}
}

func j(data cmath.IMatrix, Class classes) float64 {
	f := 0.0
	for _, v := range Class {
		for _, u := range v.Class {
			f += math.Pow(cmath.EuclidDistance(v.Center, data.Vector(u)), 2)
		}
	}
	return f
}

func classmatrix(data cmath.IMatrix, c []int) cmath.IMatrix {
	a := cmath.NewMatrix(len(c), data.DimM())
	for i, v := range c {
		for j := 0; j < data.DimM(); j++ {
			a.Set(i, j, data.At(v, j))
		}
	}
	return a
}

func max(data cmath.IMatrix) float64 {
	m := math.Inf(-1)
	for i := 0; i < data.DimN(); i++ {
		for j := 0; j < data.DimM(); j++ {
			if m < data.At(i, j) {
				m = data.At(i, j)
			}
		}
	}
	return m
}

func min(data cmath.IMatrix) float64 {
	m := math.Inf(1)
	for i := 0; i < data.DimN(); i++ {
		for j := 0; j < data.DimM(); j++ {
			if m > data.At(i, j) {
				m = data.At(i, j)
			}
		}
	}
	return m
}

func trimClass(center classes) classes {
	c := make(classes, 0)
	for i := 0; i < len(center); i++ {
		if len(center[i].Class) != 0 {
			c = append(c, center[i])
		}
	}
	return c
}

func shortestD(center classes, i int, data cmath.IVector) float64 {
	dis := cmath.EuclidDistance(center.Center(0), data)
	for j := 1; j < i; j++ {
		if d := cmath.EuclidDistance(center.Center(j), data); dis > d {
			dis = d
		}
	}
	return math.Pow(dis, 2)
}

func probability(center classes, i int, data cmath.IMatrix, dst cmath.IVector) float64 {
	dt := 0.0
	for j := 0; j < data.DimN(); j++ {
		dt += shortestD(center, i, data.Vector(j))
	}
	return shortestD(center, i, dst) / dt
}

func randomCenter(data cmath.IMatrix, center classes) {
	if KMeansPP {
		kmeansPP(data, center)
		return
	}
	normalKmean(data, center)
}

func normalKmean(data cmath.IMatrix, center classes) {
	selectbox := make([]int, data.DimN())
	for i := 0; i < data.DimN(); i++ {
		selectbox[i] = i
	}
	for i, j := data.DimN()-1, 0; i >= 0 && j < center.Len(); i-- {
		n := cmath.RandInt(i)
		center[j].Center = data.Vector(selectbox[n])
		selectbox[n], selectbox[i] = selectbox[i], selectbox[n]
		j++
	}
}

func kmeansPP(data cmath.IMatrix, center classes) {
	selectbox := make(map[int]int)
	initC := cmath.RandInt(data.DimN())
	center[0].Center = data.Vector(initC)
	selectbox[initC] = 1
	for i := 1; i < center.Len(); i++ {
		Probox := cmath.NewVector(data.DimN() + 1)
		Probox.Set(0, 0)
		sum := 0.0
		for k := 0; k < data.DimN(); k++ {
			if _, ok := selectbox[k]; ok {
				Probox.Set(k+1, sum)
				continue
			}
			sum += probability(center, i, data, data.Vector(k))
			Probox.Set(k+1, sum)
		}
		for k := 0; k < data.DimN(); k++ {
			Probox.Set(k+1, Probox.At(k+1)/sum)
		}
		p := cmath.RandFloat()
		for k := 1; k < data.DimN(); k++ {
			if p < Probox.At(k) && p >= Probox.At(k-1) {
				center[i].Center = data.Vector(k - 1)
				selectbox[k] = 1
				break
			}
		}
	}
}
