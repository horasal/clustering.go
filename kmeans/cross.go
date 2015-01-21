/*

  Clustering Analysis Package
  Kmeans:CrossKmeans Unit
  Copyright (c) Zhai HongJie 2012

*/
package kmeans

import (
	"clustering/math"
	"clustering/statistics"
	"math"
)

func avgdis(data cmath.IMatrix, c cluster) float64 {
	r := 0.0
	for _, v := range c.Class {
		r += cmath.EuclidDistance(data.Vector(v), c.Center)
	}
	return r / float64(len(c.Class))
}

func classifyCross(data cmath.IMatrix, Class classes) classes {
	for _, v := range Class {
		v.Class = []int{}
	}
	for i := 0; i < data.DimN(); i++ {
		min := 0
		dis := math.Inf(1)
		for j, v := range Class {
			if t := cmath.EuclidDistance(v.Center, data.Vector(i)); dis > t {
				dis = t
				min = j
			}
		}

		//if avgdis(data, Class[min]) + math.Sqrt(cstatistics.Variance(classmatrix(data, Class[min].Class))) > cmath.EuclidDistance(Class[min].Center,data.Vector(j))
		if statistics.Variance(classmatrix(data, append(Class[min].Class, i))) > Variance {
			c := newCluster()
			c.Center = data.Vector(i)
			c.Class = []int{i}
			Class = append(Class, c)
		} else {
			Class[min].Class = append(Class[min].Class, i)
		}
	}
	return Class
}

func centerCross(data cmath.IMatrix, Class classes) {
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

func KmeansCross(data cmath.IMatrix, k int) IClasses {
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
		resc = classifyCross(data, resc)
		f1 = j(data, resc)
		trimClass(resc)
		centerCross(data, resc)
		count++
		if count > Maxloop {
			break
		}
	}
	return trimClass(resc)
}
