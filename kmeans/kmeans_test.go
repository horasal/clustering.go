package kmeans_test

import (
	"clustering/kmeans"
	"clustering/math"
	"testing"
)

func randomPoint(k, n int) cmath.IMatrix {
	m := cmath.NewMatrix((k+1)*n, 2)
	kf := float64(k)
	for i := 0; i < k; i++ {
		x := float64(cmath.RandInt(10*k) - 5*k)
		y := float64(cmath.RandInt(10*k) - 5*k)
		for j := 0; j < n; j++ {
			m.Set(i*n+j, 0, x+cmath.RandFloat()*kf-kf/2)
			m.Set(i*n+j, 1, y+cmath.RandFloat()*kf-kf/2)
		}
	}
	for j := 0; j < n; j++ {
		m.Set(k*n+j, 0, cmath.RandFloat()*10*kf-5*kf)
		m.Set(n*k+j, 1, cmath.RandFloat()*10*kf-5*kf)
	}
	return m
}

func TestKmeans(t *testing.T) {
	kmeans.Threshold = 0.0000001
	kmeans.Variance = 1.0
	kmeans.FuzzyCThreshold = 0.8
	kmeans.FuzzyM = 1.5
	classNum := 4
	m := randomPoint(4, 10)
	t.Logf("test datas:\n %s", m.String())
	t.Log("test clasical kmeans...")
	classical := kmeans.Kmeans(m, classNum)
	for i := 0; i < classical.Len(); i++ {
		t.Logf("class : %d , center: %s", i, classical.Center(i).String())
		for _, v := range classical.Class(i) {
			t.Logf("%s", m.Vector(v).String())
		}
	}
	t.Log("test KMedoids...")
	kmidoids := kmeans.KMedoids(m, classNum)
	for i := 0; i < kmidoids.Len(); i++ {
		t.Logf("class : %d , center: %s", i, kmidoids.Center(i).String())
		for _, v := range kmidoids.Class(i) {
			t.Logf("%s", m.Vector(v).String())
		}
	}
	t.Log("test kmeans with variance limit...")
	variance := kmeans.KmeansCross(m, classNum)
	for i := 0; i < variance.Len(); i++ {
		t.Logf("class : %d , center: %s", i, variance.Center(i).String())
		for _, v := range variance.Class(i) {
			t.Logf("%s", m.Vector(v).String())
		}
	}
	t.Log("test fuzzy kmeans...")
	fuzzy := kmeans.KmeansFuzzy(m, classNum)
	t.Log(fuzzy.String())
}
