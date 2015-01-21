/*

  Clustering Analysis Package
  Kmeans:FuzzyKmeans Unit
  Copyright (c) Zhai HongJie 2012

*/

package kmeans

import (
	"clustering/math"
	"math"
)

var (
	FuzzyM          float64 = 5
	FuzzyCThreshold float64 = 1
)

func jFuzzy(data cmath.IMatrix, Class classes, fuzzy cmath.IMatrix) float64 {
	f := 0.0
	for i := 0; i < data.DimN(); i++ {
		for j, _ := range Class {
			f += math.Pow(fuzzy.At(i, j), FuzzyM) * math.Pow(dis(i, j, data, Class), 2)
		}
	}
	return f
}

func dis(i, j int, data cmath.IMatrix, Class classes) float64 {
	return cmath.EuclidDistance(data.Vector(i), Class[j].Center)
}

func classifyFuzzy(data cmath.IMatrix, Class classes, fuzzy cmath.IMatrix) {
	for i := 0; i < data.DimN(); i++ {
		for j, _ := range Class {
			c := 0.0
			for k, _ := range Class {
				c += math.Pow(dis(i, j, data, Class)/dis(i, k, data, Class), 2/(FuzzyM-1))
			}
			fuzzy.Set(i, j, 1/c)
		}
	}
}

func centerFuzzy(data cmath.IMatrix, Class classes, fuzzy cmath.IMatrix) {
	for j, v := range Class {
		sum := 0.0
		for i := 0; i < data.DimN(); i++ {
			sum += math.Pow(fuzzy.At(i, j), FuzzyM)
		}
		for k := 0; k < v.Center.Dim(); k++ {
			val := 0.0
			for i := 0; i < data.DimN(); i++ {
				val += math.Pow(fuzzy.At(i, j), FuzzyM) * data.Vector(i).At(k)
			}
			v.Center.Set(k, val/sum)
		}
	}
}

func KmeansFuzzy(data cmath.IMatrix, k int) cmath.IMatrix {
	fuzzy := cmath.NewMatrix(data.DimN(), k)
	if k > data.DimN() {
		return nil
	}
	resc := make(classes, k)
	for i := 0; i < k; i++ {
		resc[i] = new(cluster)
		resc[i].Class = []int{i}
	}
	checkenvirenment()
	for i := 0; i < k; i++ {
		resc[i].Center = cmath.NewVector(data.DimM())
	}
	randomCenter(data, resc)
	count := 0
	for i := 0; i < fuzzy.DimN(); i++ {
		for j := 0; j < fuzzy.DimM(); j++ {
			fuzzy.Set(i, j, cmath.RandFloat())
		}
	}
	if k == data.DimN() {
		return fuzzy
	}
	f1 := math.Inf(1)
	for math.Abs(f1-jFuzzy(data, resc, fuzzy)) > Threshold {
		centerFuzzy(data, resc, fuzzy)
		f1 = jFuzzy(data, resc, fuzzy)
		classifyFuzzy(data, resc, fuzzy)
		count++
		if count > Maxloop {
			break
		}
	}
	return fuzzy
}
