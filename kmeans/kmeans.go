/*

  Clustering Analysis Package
  Kmeans:Kmeans Unit
  Copyright (c) Zhai HongJie 2012

*/
package kmeans

import (
	"clustering/math"
	"math"
)

func classify(data cmath.IMatrix, Class classes) {
	for _, v := range Class {
		v.Class = []int{}
	}
	for i := 0; i < data.DimN(); i++ {
		min, dis := 0, math.Inf(1)
		for j, v := range Class {
			if t := cmath.EuclidDistance(v.Center, data.Vector(i)); dis > t {
				min, dis = j, t
			}
		}
		Class[min].Class = append(Class[min].Class, i)
	}
}

func center(data cmath.IMatrix, Class classes) {
	for _, v := range Class {
		v.Center = cmath.NewVector(data.DimM())
		for _, u := range v.Class {
			for k := 0; k < data.Vector(u).Dim(); k++ {
				v.Center.Set(k, v.Center.At(k)+data.Vector(u).At(k))
			}
		}
		for k := 0; k < v.Center.Dim(); k++ {
			v.Center.Set(k, v.Center.At(k)/float64(len(v.Class)))
		}
	}
}

func Kmeans(data cmath.IMatrix, k int) IClasses {
	if k > data.DimN() {
		return nil
	}
	resc := make(classes, k)
	for i := 0; i < k; i++ {
		resc[i] = new(cluster)
		resc[i].Class = []int{i}
	}
	if k == data.DimN() {
		return resc
	}
	checkenvirenment()
	for i := 0; i < k; i++ {
		resc[i].Center = cmath.NewVector(data.DimM())
	}
	randomCenter(data, resc)
	count := 0
	f1 := math.Inf(1)
	for math.Abs(f1-j(data, resc)) > Threshold {
		classify(data, resc)
		f1 = j(data, resc)
		center(data, resc)
		count++
		if count > Maxloop {
			break
		}
	}
	return trimClass(resc)
}
