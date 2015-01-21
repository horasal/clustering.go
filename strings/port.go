/*

  Clustering Analysis Package
  Strings:Interface Unit
  Copyright (c) Zhai HongJie 2012

*/

package cstrings

type IToken interface {
	Word() string
	Equal(t IToken) int
}

type ISentence interface {
	Token(i int) IToken
	TokenUnique(i int) IToken
	TokenCount(i int) int
	Len() int
	LenUnique() int

	Add(t IToken) bool
	Contain(t IToken) bool
	Search(t IToken) int
}

type IArticle interface {
	Sentence(i int) ISentence
	Len() int
	Add(s ISentence)
	Contain(t IToken) bool
}
