/*

  Clustering Analysis Package
  Strings:StringMath Unit
  Copyright (c) Zhai HongJie 2012

*/

package cstrings

import (
	"strings"
)

func CompareString(a, b string) int {
	c, d := []byte(a), []byte(b)

	for i := 0; i < len(c) && i < len(d); i++ {
		if m := int(c[i]) - int(d[i]); m != 0 {
			return m
		}
	}
	if l := int(len(c)) - int(len(d)); l > 0 {
		return int(c[len(d)])
	} else {
		if l < 0 {
			return -int(d[len(c)])
		}
	}
	return 0
}

func CompareStringIngoreCase(a, b string) int {
	return CompareString(strings.ToUpper(a), strings.ToUpper(b))
}
