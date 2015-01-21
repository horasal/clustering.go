/*

  Clustering Analysis Package
  TF-IDF:Interface Unit
  Copyright (c) Zhai HongJie 2012

*/
package tfidf

import (
	"clustering/math"
	"clustering/strings"
)

type ITokenValue interface {
	Token(i int) cstrings.IToken
	Value(i int) float64
	Len() int
	Vector() cmath.IVector
	Search(t cstrings.IToken) int
}

type tokenValue struct {
	token cstrings.ISentence
	value []float64
}

func (v tokenValue) Token(i int) cstrings.IToken {
	if i < 0 || i >= v.Len() {
		return nil
	}
	return v.token.TokenUnique(i)
}

func (v tokenValue) Value(i int) float64 {
	if i < 0 || i >= v.Len() {
		return 0
	}
	return v.value[i]
}

func (v tokenValue) Len() int {
	if v.token.LenUnique() != len(v.value) {
		return 0
	}
	return v.token.LenUnique()
}

func (v tokenValue) Vector() cmath.IVector {
	c := cmath.NewVector(v.Len())
	for i := 0; i < c.Dim(); i++ {
		c.Set(i, v.value[i])
	}
	return c
}

func (v tokenValue) Search(t cstrings.IToken) int {
	return v.token.Search(t)
}

func (v *tokenValue) add(t cstrings.IToken, f float64) {
	if v.token.Add(t) {
		v.value = append(v.value, f)
	} else {
		if i := v.token.Search(t); i >= 0 {
			v.value[i] = f
		}
	}
}

func newTokenValue() *tokenValue {
	t := new(tokenValue)
	t.token = cstrings.NewFastSentence()
	t.value = make([]float64, 0)
	return t
}
