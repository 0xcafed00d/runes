package runes

import (
//"unicode"
)

func CloneRuneSlice(r []rune) []rune {
	var res []rune = nil
	return append(res, r...)
}

func InsertRuneAt(s []rune, r rune, i int) []rune {
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = r
	return s
}

func DeleteRuneAt(s []rune, i int) []rune {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

func TrimRunesLeft(r []rune, f func(rune) bool) []rune {
	for i := 0; i < len(r); i++ {
		if !f(r[i]) {
			return r[i:]
		}
	}
	return []rune{}
}

func TrimRunesRight(r []rune, f func(rune) bool) []rune {
	for i := len(r) - 1; i >= 0; i-- {
		if !f(r[i]) {
			return r[:i+1]
		}
	}
	return []rune{}
}

func TrimRunes(r []rune, f func(rune) bool) []rune {
	return TrimRunesRight(TrimRunesLeft(r, f), f)
}
