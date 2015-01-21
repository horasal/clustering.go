/*

  Clustering Analysis Package
  Strings:String Unit
  Copyright (c) Zhai HongJie 2012

*/

package cstrings

type uniqueList struct {
	content []IToken
	count   []int
}

func (u uniqueList) Len() int           { return len(u.content) }
func (u uniqueList) Less(a, b int) bool { return u.content[a].Equal(u.content[b]) > 0 }
func (u *uniqueList) Swap(a, b int) {
	u.content[a], u.content[b] = u.content[b], u.content[a]
	u.count[a], u.count[b] = u.count[b], u.count[a]
}

func (u *uniqueList) insert(t IToken, n int) {
	if n < 0 {
		u.content = append([]IToken{t}, u.content...)
		u.count = append([]int{1}, u.count...)
	} else if n > len(u.content) {
		u.content = append(u.content, t)
		u.count = append(u.count, 1)
	} else {
		u.content = append(u.content[:n], append([]IToken{t}, u.content[n:]...)...)
		u.count = append(u.count[:n], append([]int{1}, u.count[n:]...)...)
	}
}

func (u *uniqueList) Add(t IToken) bool {
	left, right := 0, len(u.content)-1
	for i := (left + right) / 2; left <= right; i = (left + right) / 2 {
		if a := u.content[i].Equal(t); a == 0 {
			u.count[i]++
			return false
		} else if a < 0 {
			right = i - 1
		} else {
			left = i + 1
		}
	}
	u.insert(t, left)
	return true
}

func (u uniqueList) Search(t IToken) int {
	left, right := 0, len(u.content)-1
	for i := (left + right) / 2; left <= right; i = (left + right) / 2 {
		if a := u.content[i].Equal(t); a == 0 {
			return i
		} else if a < 0 {
			right = i - 1
		} else {
			left = i + 1
		}
	}
	return -1
}

func newUniqueList() *uniqueList {
	u := new(uniqueList)
	u.content = make([]IToken, 0)
	u.count = make([]int, 0)
	return u
}

type sentence struct {
	tokenNormal []IToken
	tokenUnique *uniqueList
}

func (s sentence) Contain(t IToken) bool {
	if i := s.Search(t); i >= 0 && i < s.LenUnique() {
		return true
	}
	return false
}

func (s sentence) Search(t IToken) int {
	return s.tokenUnique.Search(t)
}

func (s sentence) Token(i int) IToken {
	if i >= 0 && i < len(s.tokenNormal) {
		return s.tokenNormal[i]
	}
	return nil
}

func (s sentence) TokenUnique(i int) IToken {
	if i >= 0 && i < s.tokenUnique.Len() {
		return s.tokenUnique.content[i]
	}
	return nil
}

func (s sentence) TokenCount(i int) int {
	if i >= 0 && i < s.tokenUnique.Len() {
		return s.tokenUnique.count[i]
	}
	return 0
}

func (s sentence) Len() int { return len(s.tokenNormal) }

func (s sentence) LenUnique() int { return s.tokenUnique.Len() }

func (s *sentence) Add(t IToken) bool {
	s.tokenNormal = append(s.tokenNormal, t)
	return s.tokenUnique.Add(t)
}

type article struct {
	sent []ISentence
}

func (a article) Sentence(i int) ISentence {
	if i < 0 || i >= len(a.sent) {
		return nil
	}
	return a.sent[i]
}

func (a article) Contain(t IToken) bool {
	for _, v := range a.sent {
		if v.Contain(t) {
			return true
		}
	}
	return false
}

func (a article) Len() int { return len(a.sent) }

func (s *article) Add(t ISentence) {
	s.sent = append(s.sent, t)
}
func NewSentence() ISentence {
	s := new(fastsentence)
	s.tokenCount = make([]int, 0)
	s.tokenNormal = make([]IToken, 0)
	s.tokenUnique = make(map[string]int)
	s.tokenUniqueList = make([]IToken, 0)
	return s
}
/*
func NewSentence() ISentence {
	s := new(sentence)
	s.tokenNormal = make([]IToken, 0)
	s.tokenUnique = newUniqueList()
	return s
}*/

func NewArticle() IArticle {
	a := new(article)
	a.sent = make([]ISentence, 0)
	return a
}
