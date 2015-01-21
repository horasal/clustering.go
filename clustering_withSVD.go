/*

  Clustering Analysis Package
  Clustering:F Test Unit
  Copyright (c) Zhai HongJie 2012

*/

package main

import (
	"bufio"
	"clustering/gsl"
	"clustering/kmeans"
	"clustering/math"
	"clustering/sort"
	"clustering/spectral"
	"clustering/statistics"
	"clustering/strings"
	"clustering/tfidf"
	"flag"
	"fmt"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"
)

var (
	input     = flag.String("input", "", "input directory")
	output    = flag.String("output", "./", "output directory")
	userbase  = flag.Bool("user", true, "base on user space")
	threshold = flag.Float64("threshold", 0.5, "max tfidf value for selecting feature terms")
	k         = flag.Int("k", 20, "class to classify in")
	sig       = flag.String("sig", "classical", `set sig for clustering classical - normal kmeans
 cross - extended kmeans,also need to set variance parameter
 KMedoids - KMEDOIDS 
 *fuzzy - fuzzy kmeans, this only return fuzzy matrix
 `)
	variance   = flag.Float64("variance", 0.6, "variance threshold for extended kmeans")
	fuzzym     = flag.Float64("m", 1.5, "fuzzy set parameter M")
	dimreduce  = flag.Bool("dimreduce", false, "reduce dim when laplacian eigenmap")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file *enable this option will slow down the program")
	SVD        = flag.Bool("svd", false, "use svd when get user - topic relation")
	kmeansPP   = flag.Bool("p", false, "use kmeans++ to select the initial center *enable this option will slow down the program")
	sigma      = flag.Float64("sigma", 0, "sigma value in spectral clustering, exp(-|x_i-x_j|^2/\\sigma^2).")
	nas        = flag.Int("nas", 3, "nearest point to select in NAS")
	kernel     = flag.String("kernel", "gaussian", `kernel to use in spectral clustering
	gaussian - gaussian kernel, a sigma parameter is needed *default
	linear - linear kernel
	polynomial - polynomial kernel
	nas - nas gaussian kernel, a nas parameter is needed`)
	boolean = flag.Bool("bool", false, "user boolean matrix for user")
)

const dirperm = 0777

type token struct {
	word string
}

func (a token) Word() string { return strings.ToUpper(a.word) }
func (a token) Equal(t cstrings.IToken) int {
	return cstrings.CompareStringIngoreCase(a.Word(), t.Word())
}
func (a *token) set(s string) { a.word = strings.TrimSpace(s) }

func newtoken(s string) cstrings.IToken {
	t := new(token)
	t.word = s
	return t
}

func getsent(f *bufio.Reader) cstrings.ISentence {
	s := cstrings.NewFastSentence()
	a, err := f.ReadString('\n')
	if len(a) == 0 && err != nil {
		return nil
	}
	arr := strings.Split(a, "\t")
	for i := 1; i < len(arr); i++ {
		la := strings.Split(strings.TrimSpace(arr[i]), " ")
		for j := 0; j < len(la); j++ {
			s.Add(newtoken(la[j]))
		}
	}
	return s
}

func GetArticle(fi os.FileInfo) cstrings.IArticle {
	a := cstrings.NewArticle()
	f, err := os.Open(*input + "/" + fi.Name())
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer f.Close()
	bf := bufio.NewReader(f)
	for s := getsent(bf); s != nil; s = getsent(bf) {
		a.Add(s)
	}
	return a
}

func getuserNameList(fi os.FileInfo) string {
	n := ""
	arr := strings.Split(fi.Name(), "_")
	for i := 1; i < len(arr)-1; i++ {
		n += arr[i] + "_"
	}
	return n[:len(n)-1] //strings.TrimRight(n, "_")
}

func buildSimilarityMatrix(matrix cmath.IMatrix) cmath.IMatrix {
	m := cmath.NewMatrix(matrix.DimN(), matrix.DimN())
	for i := 0; i < matrix.DimN(); i++ {
		for j := 0; j < matrix.DimN(); j++ {
			m.Set(i, j, cmath.EuclidDistance(matrix.Vector(i), matrix.Vector(j)))
		}
	}
	return m
}

func buildDocumentVector(as []cstrings.IArticle, sc cstrings.ISentence, ti []tfidf.ITokenValue) cmath.IMatrix {
	count := 0
	for _, v := range as {
		count += v.Len()
	}
	matrix := cmath.NewMatrix(sc.LenUnique(), count)
	for j := 0; j < sc.LenUnique(); j++ {
		i := 0
		for k := 0; k < len(as); k++ {
			for u := 0; u < as[k].Len(); u++ {
				if as[k].Sentence(u).Contain(sc.Token(j)) {
					matrix.Set(j, i, 1)
				} else {
					matrix.Set(j, i, 0)
				}
				i++
			}
		}
	}
	return matrix
}

func buildUserVector(as []cstrings.IArticle, sc cstrings.ISentence, ti []tfidf.ITokenValue) cmath.IMatrix {
	matrix := cmath.NewMatrix(sc.LenUnique(), len(as))
	for j := 0; j < sc.LenUnique(); j++ {
		for i, v := range ti {
			if v.Search(sc.TokenUnique(j)) > 0 {
				matrix.Set(j, i, 1)
			} else {
				matrix.Set(j, i, 0)
			}
		}
	}
	return matrix
}

func buildUserVectorN(as []cstrings.IArticle, sc cstrings.ISentence, ti []tfidf.ITokenValue) cmath.IMatrix {
	matrix := cmath.NewMatrix(sc.LenUnique(), len(as))
	for j := 0; j < sc.LenUnique(); j++ {
		for i, v := range ti {
			if v.Search(sc.TokenUnique(j)) > 0 {
				matrix.Set(j, i, v.Value(v.Search(sc.TokenUnique(j))))
			} else {
				matrix.Set(j, i, 0)
			}
		}
	}
	return matrix.NormalizeN()
}

func Cosine(space cmath.IMatrix, class kmeans.IClasses) cmath.IVector {
	c := cmath.NewVector(class.Len())
	for i := 0; i < class.Len(); i++ {
		count := 0.0
		for _, v := range class.Class(i) {
			for _, n := range class.Class(i) {
				if v == n {
					continue
				}
				c.Set(i, c.At(i)+statistics.Cosine(space.Vector(v), space.Vector(n)))
				count++
			}
		}
		c.Set(i, c.At(i)/count)
	}
	return c
}

func cluster(m cmath.IMatrix, sig string) kmeans.IClasses {
	fmt.Print("clustering...")
	var c kmeans.IClasses
	switch strings.ToUpper(sig) {
	case "CLASSICAL":
		c = kmeans.Kmeans(m, *k)
	case "CROSS":
		c = kmeans.KmeansCross(m, *k)
	case "KMEDOIDS":
		c = kmeans.KMedoids(m, *k)
	/*case "FUZZY":
	c = kmeans.KmeansFuzzy(m, *k)*/
	default:
		c = kmeans.Kmeans(m, *k)
	}
	fmt.Println("ok")
	return c
}

func print(m cmath.IMatrix, c kmeans.IClasses, sc cstrings.ISentence) {
	cos := Cosine(m, c)
	for i := 0; i < c.Len(); i++ {
		fmt.Printf("class: %d cosine: %f\n", i, cos.At(i))
		for _, j := range c.Class(i) {
			fmt.Printf("%s, ", sc.TokenUnique(j).Word())
		}
		fmt.Println()
	}
}

//user struct
//name : string
//center for document in words space
//Imatrix - for document
/*
		w1	w2	w3
u1	d1
	d2	x1	x2	x3
	d3
u2	d1
	d2	x4	x5	x6
	d3
*/

//topic struct
//Imatrix for words
//center for words in document space
/*
		d1	d2	d3
t1	w1
	w2	x1	x2	x3
	w3
t2	w1
	w2	x4	x5	x6
	w3
*/

type user struct {
	username string
	document []cstrings.IArticle
	matrix   cmath.IMatrix
}

func newUser() *user {
	u := new(user)
	u.document = make([]cstrings.IArticle, 0)
	u.matrix = cmath.NewMatrix(0, 0)
	return u
}

type topic struct {
	center cmath.IVector
	words  cstrings.ISentence
	matrix cmath.IMatrix
}

func newTopic() *topic {
	t := new(topic)
	t.words = cstrings.NewSentence()
	t.matrix = cmath.NewMatrix(0, 0)
	t.center = cmath.NewVector(0)
	return t
}

const FuzzyM = 1.1

func mycenter(data cmath.IMatrix, class []int) cmath.IVector {
	v := cmath.NewVector(data.DimM())
	for i := 0; i < len(class); i++ {
		for j := 0; j < data.DimM(); j++ {
			v.Set(j, v.At(j)+data.At(class[i], j))
		}
	}
	for j := 0; j < v.Dim(); j++ {
		v.Set(j, v.At(j)/float64(len(class)))
	}
	return v
}

func dis(i, j int, destdata cmath.IMatrix, center cmath.IMatrix) float64 {
	return cmath.EuclidDistance(destdata.Vector(i), center.Vector(j))
}

func getFuzzyset(destdata cmath.IMatrix, Centers cmath.IMatrix) cmath.IMatrix {
	m := cmath.NewMatrix(destdata.DimN(), Centers.DimN())
	for i := 0; i < destdata.DimN(); i++ {
		for j := 0; j < Centers.DimN(); j++ {
			c := 0.0
			v := dis(i, j, destdata, Centers)
			for k := 0; k < Centers.DimN(); k++ {
				vk := dis(i, k, destdata, Centers)
				if v == vk {
					c += 1
				} else {
					c += math.Pow(v/vk, 2/(FuzzyM-1))
				}
			}
			m.Set(i, j, 1/c)
		}
	}
	return m
}

func getFuzzysetU(destdata cmath.IMatrix, Centers cmath.IMatrix) cmath.IMatrix {
	m := cmath.NewMatrix(destdata.DimN(), destdata.DimM())
	for i := 0; i < destdata.DimN(); i++ {
		c := 0.0
		for j := 0; j < destdata.DimM(); j++ {
			c += destdata.At(i, j)
		}
		for j := 0; j < destdata.DimM(); j++ {
			m.Set(i, j, destdata.At(i, j)/c)
		}
	}
	return m
}

func getFuzzysetZ(destdata cmath.IMatrix, Centers cmath.IMatrix) cmath.IMatrix {
	m := cmath.NewMatrix(destdata.DimN(), Centers.DimN())
	for i := 0; i < destdata.DimN(); i++ {
		for j := 0; j < Centers.DimN(); j++ {
			m.Set(i, j, dis(i, j, destdata, Centers))
		}
	}
	return m
}

func linearNormalize(m cmath.IMatrix) cmath.IMatrix {
	matrix := cmath.NewMatrix(m.DimN(), m.DimM())
	for i := 0; i < m.DimN(); i++ {
		t := 0.0
		for j := 0; j < m.DimM(); j++ {
			t += m.At(i, j)
		}
		for j := 0; j < m.DimM(); j++ {
			matrix.Set(i, j, m.At(i, j)/t)
		}
	}
	return matrix
}

func entropy(v cmath.IVector) float64 {
	r := 0.0
	for i := 0; i < v.Dim(); i++ {
		if v.At(i) == 0 {
			continue
		}
		r += -v.At(i) * math.Log(v.At(i))
	}
	return r
}

func sent(s cstrings.ISentence) string {
	r := ""
	for i := 0; i < s.Len(); i++ {
		r += s.Token(i).Word()
	}
	return r
}

var elaspedTime = time.Now().Unix()

func main() {
	flag.Parse()
	if *input == "" {
		flag.PrintDefaults()
		return
	}
	if *output == "" {
		*output = "./"
	}
	if !strings.HasSuffix(*output, "/") {
		*output += "/"
	}
	err := os.MkdirAll(*output, dirperm)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			fmt.Println(err)
			return
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	kmeans.Variance = *variance
	kmeans.KMeansPP = *kmeansPP
	spectral.Sigma = *sigma
	spectral.NAS = *nas
	dir, err := os.Open(*input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fi, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print("collecting information...")
	as := make([]cstrings.IArticle, len(fi))
	am := make([]string, len(fi))
	count := 0
	for i, v := range fi {
		as[i] = GetArticle(v)
		am[i] = getuserNameList(v)
		count += as[i].Len()
	}
	fmt.Println("ok")
	fmt.Printf("total tweets: %d\n", count)
	fmt.Print("tf...")
	tf := tfidf.TF(as)
	fmt.Println("ok")
	fmt.Print("idf...")
	idf := tfidf.IDF(tf)
	fmt.Println("ok")
	fmt.Print("tfidf...")
	ti := tfidf.TFIDF(tf, idf)
	fmt.Println("ok")
	fmt.Print("collect words...")
	sc := cstrings.NewSentence()
	for _, v := range ti {
		s := csort.Sort(v.Vector())
		for j := 0; j < len(s) && v.Value(s[j]) > (v.Value(s[0])+v.Value(s[len(s)-1]))*(*threshold); j++ {
			sc.Add(v.Token(s[j]))
		}
	}
	fmt.Println("ok")
	fmt.Printf("total %d words choosed\n", sc.LenUnique())

	fmt.Printf("write relation file...")
	ff, err := os.Create(*output + "user-term.txt")
	if err != nil {
		panic(err.Error())
	}
	bfuser := bufio.NewWriter(ff)
	for i := 0; i < len(tf); i++ {
		bfuser.WriteString(fmt.Sprintf("\"%s\",", am[i]))
		for j := 0; j < sc.LenUnique(); j++ {
			if tf[i].Search(sc.TokenUnique(j)) > 0 {
				bfuser.WriteString(fmt.Sprintf("\"%s\",", sc.TokenUnique(j).Word()))
			}
		}
		bfuser.WriteString("\n")
	}
	bfuser.Flush()
	ff.Close()
	ff, err = os.Create(*output + "document-term.txt")
	if err != nil {
		panic(err.Error())
	}
	bfdocument := bufio.NewWriter(ff)
	for i := 0; i < len(as); i++ {
		for j := 0; j < as[i].Len(); j++ {
			bfdocument.WriteString(fmt.Sprintf("\"%s\",", am[i]))
			bfdocument.WriteString(fmt.Sprintf("\"%s\",", sent(as[i].Sentence(j))))
			for k := 0; k < as[i].Sentence(j).LenUnique(); k++ {
				if sc.Contain(as[i].Sentence(j).TokenUnique(k)) {
					bfdocument.WriteString(fmt.Sprintf("\"%s\",", as[i].Sentence(j).TokenUnique(k).Word()))
				}
			}
			/*
				for k := 0; k < sc.LenUnique(); k++ {
					if as[i].Sentence(j).Contain(sc.TokenUnique(k)) {
						bfdocument.WriteString(fmt.Sprintf("\"%s\",", sc.TokenUnique(k).Word()))
					}
				}
			*/
			bfdocument.WriteString("\n")
		}
	}
	bfdocument.Flush()
	ff.Close()
	fmt.Println("ok")

	fmt.Print("build matrix...")
	var matrix cmath.IMatrix
	if *userbase {
		if !*boolean {
			matrix = buildUserVectorN(as, sc, tf)
		} else {
			matrix = buildUserVector(as, sc, ti)
		}
	} else {
		matrix = buildDocumentVector(as, sc, ti)
	}
	fmt.Println("ok")
	matrix = matrix.NormalizeN()
	fmt.Print("clear memory...")
	runtime.GC()
	fmt.Println("ok")
	fmt.Printf("time spend: %ds\n", time.Now().Unix()-elaspedTime)
	fmt.Print("kernel...")
	sMatrix := buildSimilarityMatrix(matrix)
	switch strings.ToUpper(*kernel) {
	case "GAUSSIAN":
		fmt.Print("gaussian...")
		sMatrix = spectral.GaussianKernel(sMatrix)
	case "NAS":
		fmt.Print("nas...")
		sMatrix = spectral.NASGaussianKernel(sMatrix)
	case "POLYNOMIAL":
		fmt.Print("poly...")
		sMatrix = spectral.PolynomialKernel(sMatrix)
	default:
		fmt.Print("default(linear)...")
	}
	fmt.Println("ok")
	fmt.Print("laplacian...")
	eigenvalue, eigenvector := spectral.LpEigen(spectral.Laplacian(sMatrix))
	fmt.Println("ok")
	fmt.Print("build topic...")
	var m cmath.IMatrix
	if *dimreduce {
		s := csort.MaxN(eigenvalue, *k)
		m = cmath.NewMatrix(eigenvector.DimM(), *k)
		for j, v := range s {
			for i := 0; i < eigenvector.Vector(v).Dim(); i++ {
				m.Set(i, j, eigenvector.Vector(v).At(i))
			}
		}
	} else {
		m = cmath.NewMatrix(eigenvector.DimM(), eigenvalue.Dim())
		for j := 0; j < eigenvalue.Dim(); j++ {
			for i := 0; i < eigenvector.Vector(j).Dim(); i++ {
				m.Set(i, j, eigenvector.Vector(j).At(i))
			}
		}
	}
	m = m.NormalizeN()
	fmt.Println("ok")
	c := cluster(m, *sig)
	fmt.Printf("preset cluster: %d, cluster get: %d\n", *k, c.Len())
	fmt.Print("save topic information...")
	topicinfo, err := os.Create(*output + "topic.txt")
	if err != nil {
		panic(err.Error())
	}
	bftopic := bufio.NewWriter(topicinfo)
	for i := 0; i < c.Len(); i++ {
		bftopic.WriteString(fmt.Sprintf("Class%d:\n", i))
		for _, v := range c.Class(i) {
			bftopic.WriteString(fmt.Sprintf("%s, ", sc.TokenUnique(v).Word()))
		}
		bftopic.WriteString("\n")
	}
	bftopic.Flush()
	topicinfo.Close()
	fmt.Println("ok")
	fmt.Println("cosine:", Cosine(matrix, c).String())
	fmt.Print("build user term relation...")
	userTermMatrix := cmath.NewMatrix(len(fi), sc.LenUnique())
	for i := 0; i < len(tf); i++ {
		for j := 0; j < sc.LenUnique(); j++ {
			if a := tf[i].Search(sc.TokenUnique(j)); a > 0 && a < tf[i].Len() {
				userTermMatrix.Set(i, j, 1)
			} else {
				userTermMatrix.Set(i, j, 0)
			}
		}
	}
	fmt.Println("ok")

	fmt.Print("get topic center...")
	cen := cmath.NewMatrix(c.Len(), matrix.DimM())
	for i := 0; i < c.Len(); i++ {
		c := mycenter(matrix, c.Class(i))
		for j := 0; j < c.Dim(); j++ {
			cen.Set(i, j, c.At(j))
		}
	}
	cen = cen.NormalizeN()
	fmt.Println("ok")

	fmt.Print("get user document relation...")
	var userDocumentMatrix cmath.IMatrix
	if *userbase {
		userDocumentMatrix = cen.Transposition()
		cen = cmath.NewMatrix(c.Len(), c.Len())
		for i := 0; i < c.Len(); i++ {
			cen.Set(i, i, 1)
		}
	} else {
		if *SVD {
			cc, _, _ := gslgo.SVDDecompostion(matrix)
			userDocumentMatrix = userTermMatrix.Mul(cc).NormalizeN()
		} else {
			userDocumentMatrix = userTermMatrix.Mul(matrix).NormalizeN()
		}
	}
	fmt.Println("ok")

	fmt.Print("get user topic relation ...")
	var userTopic cmath.IMatrix
	if *userbase {
		userTopic = getFuzzysetU(userDocumentMatrix, cen)
	} else {
		userTopic = getFuzzyset(userDocumentMatrix, cen)
	}
	fmt.Println("ok")

	fmt.Print("save user ranking...")
	evector := cmath.NewVector(userTopic.DimN())
	for i := 0; i < userTopic.DimN(); i++ {
		evector.Set(i, entropy(userTopic.Vector(i)))
	}
	cs := csort.Sort(evector)
	utfile, err := os.Create(*output + "user_ranking.txt")
	if err != nil {
		panic(err.Error())
	}
	utbuffer := bufio.NewWriter(utfile)
	for i := 0; i < len(cs); i++ {
		utbuffer.WriteString(fmt.Sprintf("\"%s\",", am[cs[i]]))
		utbuffer.WriteString(fmt.Sprintf("%f\n", evector.At(cs[i])))
	}
	utbuffer.Flush()
	utfile.Close()
	fmt.Println("ok")

	fmt.Print("save user - topic matrix...")
	umfile, err := os.Create(*output + "user_topic_matrix.txt")
	if err != nil {
		panic(err.Error())
	}
	umbuffer := bufio.NewWriter(umfile)
	for i := 0; i < len(cs); i++ {
		umbuffer.WriteString(fmt.Sprintf("\"%s\",", am[cs[i]]))
		umbuffer.WriteString(fmt.Sprintf("%s\n", userTopic.Vector(cs[i])))
	}
	umbuffer.Flush()
	umfile.Close()
	fmt.Println("ok")
	tms := time.Now().Unix() - elaspedTime
	tmm := tms / 60
	tmh := tmm / 60
	fmt.Printf("total time: %dh %dm %ds\n", tmh%60, tmm%60, tms%60)
}
