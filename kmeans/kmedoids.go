/*

  Clustering Analysis Package
  Kmeans:KMedoids Unit
  Copyright (c) Zhai HongJie 2012

*/
package kmeans

import (
	"clustering/math"
	"math"
)

func centerMedoids(data cmath.IMatrix, Class classes) {
	for _, v := range Class {
		dis, min := math.Inf(1), 0
		for _, u := range v.Class {
			if k := cmath.EuclidDistance(v.Center, data.Vector(u)); dis > k {
				dis, min = k, u
			}
		}
		v.Center = data.Vector(min)
	}
}

func KMedoids(data cmath.IMatrix, k int) IClasses {
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
		centerMedoids(data, resc)
		count++
		if count > Maxloop {
			break
		}
	}
	return trimClass(resc)
}
