package cstrings_test

import (
	"clustering/strings"
	"testing"
)

func TestCompareString(t *testing.T) {
	t.Log("test CompareString...")
	t.Logf("abcdEG VS abcdEG: %d", cstrings.CompareString("abcdEG", "abcdEG"))
	t.Logf("abcdEG VS abcdEGF: %d", cstrings.CompareString("abcdEG", "abcdEGF"))
	t.Logf("abcdEGF VS abcdEG: %d", cstrings.CompareString("abcdEGF", "abcdEG"))
	t.Logf("abddEG VS abcdEG: %d", cstrings.CompareString("abddEG", "abcdEG"))
	t.Logf("abcdEG VS abddEG: %d", cstrings.CompareString("abcdEG", "abddEG"))
	t.Logf("abddEG VS abcdEG: %d", cstrings.CompareString("abddEGF", "abcdEG"))
	t.Logf("abcdEG VS abddEG: %d", cstrings.CompareString("abcdEGF", "abddEG"))
	t.Logf("abcdEG VS \"\": %d", cstrings.CompareString("abcdEGF", ""))
	t.Logf("\"\" VS abcdEG: %d", cstrings.CompareString("", "abcdEGF"))
	t.Logf("ABCDEG VS abcdEG: %d", cstrings.CompareString("ABCDEG", "abcdEG"))
	t.Logf("abcdEG VS abcdeg: %d", cstrings.CompareString("abcdEG", "abcdeg"))
}

func TestCompareStringIngoreCase(t *testing.T) {
	t.Log("test CompareString...")
	t.Logf("abcdEG VS abcdEG: %d", cstrings.CompareStringIngoreCase("abcdEG", "abcdEG"))
	t.Logf("abcdEG VS abcdEGF: %d", cstrings.CompareStringIngoreCase("abcdEG", "abcdEGF"))
	t.Logf("abcdEGF VS abcdEG: %d", cstrings.CompareStringIngoreCase("abcdEGF", "abcdEG"))
	t.Logf("abddEG VS abcdEG: %d", cstrings.CompareStringIngoreCase("abddEG", "abcdEG"))
	t.Logf("abcdEG VS abddEG: %d", cstrings.CompareStringIngoreCase("abcdEG", "abddEG"))
	t.Logf("abddEG VS abcdEG: %d", cstrings.CompareStringIngoreCase("abddEGF", "abcdEG"))
	t.Logf("abcdEG VS abddEG: %d", cstrings.CompareStringIngoreCase("abcdEGF", "abddEG"))
	t.Logf("abcdEG VS \"\": %d", cstrings.CompareStringIngoreCase("abcdEGF", ""))
	t.Logf("\"\" VS abcdEG: %d", cstrings.CompareStringIngoreCase("", "abcdEGF"))
	t.Logf("ABCDEG VS abcdEG: %d", cstrings.CompareStringIngoreCase("ABCDEG", "abcdEG"))
	t.Logf("abcdEG VS abcdeg: %d", cstrings.CompareStringIngoreCase("abcdEG", "abcdeg"))
}
