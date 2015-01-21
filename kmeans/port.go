/*

  Clustering Analysis Package
  Kmeans:Interface Unit
  Copyright (c) Zhai HongJie 2012

*/

package kmeans

import "clustering/math"

type IClasses interface {
	Len() int
	Center(i int) cmath.IVector
	Class(i int) []int
}

type cluster struct {
	Center cmath.IVector
	Class  []int
}

type classes []*cluster

func (c classes) Center(i int) cmath.IVector {
	return c[i].Center
}

func (c classes) Class(i int) []int {
	return c[i].Class
}

func (c classes) Len() int {
	return len(c)
}

func newCluster() *cluster {
	return new(cluster)
}
