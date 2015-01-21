package cstrings_test

import (
	"clustering/strings"
	"testing"
)

func TestNewFastSentence(t *testing.T) {
	t.Log("test FastSentence...")
	t.Logf("TestData: %v", testdata)
	s := cstrings.NewFastSentence()
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
