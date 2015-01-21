/*

  Clustering Analysis Package
  TF-IDF:TFIDF Unit
  Copyright (c) Zhai HongJie 2012

*/
package tfidf

import (
	"clustering/strings"
	"math"
)

func tf(v cstrings.IArticle) ITokenValue {
	sc := cstrings.NewSentence()
	for i := 0; i < v.Len(); i++ {
		for j := 0; j < v.Sentence(i).Len(); j++ {
			sc.Add(v.Sentence(i).Token(j))
		}
	}
	c := newTokenValue()
	for i := 0; i < sc.LenUnique(); i++ {
		c.add(sc.TokenUnique(i), float64(sc.TokenCount(i))/float64(sc.Len()))
	}
	return c
}

func TF(A []cstrings.IArticle) []ITokenValue {
	c := make([]ITokenValue, len(A))
	for i, v := range A {
		c[i] = tf(v)
	}
	return c
}

func IDF(A []ITokenValue) ITokenValue {
	sc := cstrings.NewSentence()
	for _, v := range A {
		for i := 0; i < v.Len(); i++ {
				sc.Add(v.Token(i))
		}
	}
	c := newTokenValue()
	for i := 0; i < sc.LenUnique(); i++ {
		count := 0
		for _, v := range A {
			if v.Search(sc.TokenUnique(i))>0 {
				count++
			}
		}
		c.add(sc.TokenUnique(i), math.Log(float64(len(A))/float64(count)))
	}
	return c
}
/*
func IDF(A []cstrings.IArticle, tf []ITokenValue) ITokenValue {
	sc := cstrings.NewSentence()
	for _, v := range A {
		for i := 0; i < v.Len(); i++ {
			for j := 0; j < v.Sentence(i).Len(); j++ {
				sc.Add(v.Sentence(i).Token(j))
			}
		}
	}
	c := newTokenValue()
	for i := 0; i < sc.LenUnique(); i++ {
		count := 0
		if i%100==0{
			fmt.Println(i,sc.LenUnique())
		}
		for _, v := range A {
			if v.Contain(sc.TokenUnique(i)) {
				count++
			}
		}
		c.add(sc.TokenUnique(i), math.Log(float64(len(A))/float64(count)))
	}
	return c
}
*/
func TFIDF(tf []ITokenValue, idf ITokenValue) []ITokenValue {
	c := make([]ITokenValue, len(tf))
	for i, v := range tf {
		sc := newTokenValue()
		for j := 0; j < v.Len(); j++ {
			sc.add(v.Token(j), v.Value(j)*idf.Value(idf.Search(v.Token(j))))
		}
		c[i] = sc
	}
	return c
}
