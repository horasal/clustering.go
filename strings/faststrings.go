package cstrings

type fastsentence struct {
	tokenNormal     []IToken
	tokenUniqueList []IToken
	tokenCount      []int
	tokenUnique     map[string]int
}

func (s fastsentence) Contain(t IToken) bool {
	_, ok := s.tokenUnique[t.Word()]
	return ok
}

func (s fastsentence) Search(t IToken) int {
	i, ok := s.tokenUnique[t.Word()]
	if ok {
		return i
	}
	return -1
}

func (s fastsentence) Token(i int) IToken {
	if i >= 0 && i < len(s.tokenNormal) {
		return s.tokenNormal[i]
	}
	return nil
}

func (s fastsentence) TokenUnique(i int) IToken {
	if i >= 0 && i < s.LenUnique() {
		return s.tokenUniqueList[i]
	}
	return nil
}

func (s fastsentence) TokenCount(i int) int {
	if i >= 0 && i < s.LenUnique() {
		return s.tokenCount[i]
	}
	return 0
}

func (s fastsentence) Len() int { return len(s.tokenNormal) }

func (s fastsentence) LenUnique() int { return len(s.tokenUniqueList) }

func (s *fastsentence) Add(t IToken) bool {
	s.tokenNormal = append(s.tokenNormal, t)
	i, ok := s.tokenUnique[t.Word()]
	if !ok {
		s.tokenUniqueList = append(s.tokenUniqueList, t)
		s.tokenCount = append(s.tokenCount, 1)
		s.tokenUnique[t.Word()] = len(s.tokenUniqueList) - 1
		return true
	}
	s.tokenCount[i]++
	return false
}

func NewFastSentence() ISentence {
	s := new(fastsentence)
	s.tokenCount = make([]int, 0)
	s.tokenNormal = make([]IToken, 0)
	s.tokenUnique = make(map[string]int)
	s.tokenUniqueList = make([]IToken, 0)
	return s
}
