// Package runes implements utilities for manipulating slices of runes
package runes

// CloneSlice return a copy of the supplied rune slice
func CloneSlice(r []rune) []rune {
	return append([]rune(nil), r...)
}

// Concat joins a number of rune slices together
func Concat(r ...[]rune) []rune {
	if len(r) == 0 {
		return []rune(nil)
	} else if len(r) == 1 {
		return r[0]
	} else {
		res := CloneSlice(r[0])
		for _, v := range r[1:] {
			res = append(res, v...)
		}
		return res
	}
}

// InsertAt inserts the rune r into the slice s at index i.
// All runes after index i are moved one position right, growing the slice by one rune
// the modified slice is returned
func InsertAt(s []rune, r rune, i int) []rune {
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = r
	return s
}

// InsertSliceAt inserts the rune slice r into the slice s at index i.
// All runes after the insertion point are moved one position right, growing the slice by the length of r
// the modified slice is returned
func InsertSliceAt(s, r []rune, i int) []rune {
	return append(s[:i], append(r, s[i:]...)...)
}

// CutSliceAt cuts a slice from the rune slice s from index i, with length cnt.
// returns the origial slice with the cut removed, and a slice containing the cut runes.
func CutSliceAt(s []rune, i, cnt int) ([]rune, []rune) {
	cut := CloneSlice(s[i : i+cnt])
	return append(s[:i], s[i+cnt:]...), cut
}

// DeleteAt deletes the rune at index i from slice s
// All runes after index i are moved one position left, shrinking the slice by one rune
// the modified slice is returned
func DeleteAt(s []rune, i int) []rune {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

// TrimLeft returns a slice of r with all leading Unicode code points c satisfying f(c) removed
func TrimLeft(r []rune, f func(rune) bool) []rune {
	for i := 0; i < len(r); i++ {
		if !f(r[i]) {
			return r[i:]
		}
	}
	return []rune{}
}

// TrimRight returns a slice of the r with all trailing Unicode code points c satisfying f(c) removed.
func TrimRight(r []rune, f func(rune) bool) []rune {
	for i := len(r) - 1; i >= 0; i-- {
		if !f(r[i]) {
			return r[:i+1]
		}
	}
	return []rune{}
}

// Trim returns a slice of the r with all leading and trailing Unicode code points c satisfying f(c) removed.
func Trim(r []rune, f func(rune) bool) []rune {
	return TrimRight(TrimLeft(r, f), f)
}
