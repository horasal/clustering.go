package spectral_test

import (
	"clustering/kmeans"
	"clustering/math"
	"clustering/sort"
	"clustering/spectral"
	"testing"
)

func TestLaplacian(t *testing.T) {
	t.Log("test laplacian...")
	m := cmath.RandomSummvMatrix(10, 10)
	t.Logf("Raw Matrix:\n %s", m.String())
	lap := spectral.Laplacian(m)
	t.Logf("Laplacian Matrix:\n %s", lap.String())
	eigenval, eigenvector := spectral.LpEigen(lap)
	t.Logf("eigen value:\n %s", eigenval.String())
	t.Logf("eigenvector:\n %s", eigenvector.String())
}

var p = [][2]float64{
	{0, 0}, {0, 0.1}, {0.1, 0}, {-0.1, -0.1},
	{1, 0}, {0, 1}, {-1, 0}, {0, -1},
	{1.414, 1.414}, {-1.4, 1.4}, {1.4, -1.4}, {-1.4, -1.4},
	{2, 0}, {-2, 0}, {0, 2}, {0, -2},
}

func TestLpEigen(t *testing.T) {
	t.Log("test spectral clustering...")
	m := cmath.NewMatrix(16, 2)
	for i := 0; i < len(p); i++ {
		m.Set(i, 0, p[i][0])
		m.Set(i, 1, p[i][1])
	}
	t.Logf("Raw Matrix:\v%s", m.String())
	c := kmeans.KmeansCross(m, 3)
	t.Log("classify raw data into 3 clusters:")
	for i := 0; i < c.Len(); i++ {
		t.Logf("class : %d , center: %s", i, c.Center(i).String())
		for _, v := range c.Class(i) {
			t.Logf("%s", m.Vector(v).String())
		}
	}
	s := cmath.NewMatrix(16, 16)
	for i := 0; i < m.DimN(); i++ {
		for j := 0; j < m.DimN(); j++ {
			s.Set(i, j, cmath.EuclidDistance(m.Vector(i), m.Vector(j)))
		}
	}
	t.Logf("Similarity Matrix:\n%s", s.String())
	lap := spectral.Laplacian(s)
	t.Logf("Laplacian Matrix:\n %s", lap.String())
	eigenval, eigenvector := spectral.LpEigen(lap)
	t.Logf("eigen value:\n %s", eigenval.String())
	t.Logf("eigenvector:\n %s", eigenvector.String())
	sort := csort.MaxN(eigenval, 3)
	fs := cmath.NewMatrix(16, 3)
	for i, v := range sort {
		for j := 0; j < eigenvector.Vector(v).Dim(); j++ {
			fs.Set(j, i, eigenvector.Vector(v).At(j))
		}
	}
	spec := kmeans.KmeansCross(fs, 3)
	t.Log("spectral clustering(3 clusters):")
	for i := 0; i < spec.Len(); i++ {
		t.Logf("class : %d , center: %s", i, spec.Center(i).String())
		for _, v := range spec.Class(i) {
			t.Logf("%s", m.Vector(v).String())
		}
	}
}

func TestNASGaussianKernel(t *testing.T){
	t.Log("test NAS gaussian kernel")
	m := cmath.NewMatrix(16, 2)
	for i := 0; i < len(p); i++ {
		m.Set(i, 0, p[i][0])
		m.Set(i, 1, p[i][1])
	}
	t.Logf("matrix:\n %s", m.String())
	mm := cmath.NewMatrix(16,16)
	for i:=0;i<m.DimN();i++{
		for j:=0; j< m.DimN();j++{
			mm.Set(i,j, cmath.EuclidDistance(m.Vector(i),m.Vector(j)))
		}
	}
	t.Logf("similarity matrix:\n%s",mm.String())
	t.Logf("NAS:\n%s",spectral.NASGaussianKernel(mm).String())
	t.Logf("GaussianKernel:\n%s", spectral.GaussianKernel(mm).String())
	t.Logf("Poly:\n%s", spectral.PolynomialKernel(mm).String())
	t.Logf("linear:\n%s", spectral.LinearKernel(mm).String())
}
