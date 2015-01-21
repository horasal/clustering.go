package cstrings_test

import (
	"clustering/strings"
	"strings"
	"testing"
)

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

var testdata = [...]string{"aaa", "a", "bbb", "bbb", "aaa",
	"b", "c", "d", "e", "ff", "g", "hijk", "lmn", "aaa", "bbb"}

func TestNewSentence(t *testing.T) {
	t.Log("test NewSentence...")
	t.Logf("TestData: %v", testdata)
	s := cstrings.NewSentence()
	for _, v := range testdata {
		s.Add(newtoken(v))
	}
	ss := ""
	for i := 0; i < s.Len(); i++ {
		ss += s.Token(i).Word() + "\t"
	}
	t.Logf("Sentence Content:\n %s\n", ss)
	ss = ""
	for i := 0; i < s.LenUnique(); i++ {
		ss += s.TokenUnique(i).Word() + "\t"
	}
	t.Logf("Sentence Unique Content:\n %s\n", ss)
	t.Logf("Sentence len:%d", s.Len())
	t.Logf("Unique len:%d", s.LenUnique())
	t.Logf("Contain \"aaa\": %v", s.Contain(newtoken("aaa")))
	t.Logf("Contain \"ccc\": %v", s.Contain(newtoken("ccc")))
	t.Logf("Search \"ff\":%d", s.Search(newtoken("ff")))
	t.Logf("Search \"f\":%d", s.Search(newtoken("f")))
	t.Logf("Unique token: %d is %s", s.Search(newtoken("ff")), s.TokenUnique(s.Search(newtoken("ff"))).Word())
	t.Logf("Token count:%s is %d", "e", s.TokenCount(s.Search(newtoken("e"))))
	t.Logf("Token count:%s is %d", "bbb", s.TokenCount(s.Search(newtoken("bbb"))))
}

func TestNewArticle(t *testing.T) {
	t.Log("test NewArticle...")
	t.Logf("TestData: %v", testdata)
	s := cstrings.NewSentence()
	a := cstrings.NewArticle()
	for _, v := range testdata {
		s.Add(newtoken(v))
		j := cstrings.NewSentence()
		for i := 0; i < s.Len(); i++ {
			j.Add(s.Token(i))
		}
		a.Add(j)
	}
	t.Logf("Article len: %d", a.Len())
	ss := ""
	for i := 0; i < a.Len(); i++ {
		for j := 0; j < a.Sentence(i).LenUnique(); j++ {
			ss += a.Sentence(i).TokenUnique(j).Word() + "\t"
		}
		ss += "\n"
	}
	t.Logf("Article content:\n %s", ss)
	t.Logf("contain \"a\":%v", a.Contain(newtoken("a")))
	t.Logf("contain \"lk\":%v", a.Contain(newtoken("lk")))
	t.Logf("contain \"lmn\":%v", a.Contain(newtoken("lmn")))
}
