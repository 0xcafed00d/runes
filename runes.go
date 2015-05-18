package runes

func CloneSlice(r []rune) []rune {
	return append([]rune(nil), r...)
}

func InsertAt(s []rune, r rune, i int) []rune {
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = r
	return s
}

func InsertSliceAt(s, r []rune, i int) []rune {
	return append(s[:i], append(r, s[i:]...)...)
}

func CutSliceAt(s []rune, i, cnt int) ([]rune, []rune) {
	cut := CloneSlice(s[i : i+cnt])
	return append(s[:i], s[i+cnt:]...), cut
}

func DeleteAt(s []rune, i int) []rune {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

func TrimLeft(r []rune, f func(rune) bool) []rune {
	for i := 0; i < len(r); i++ {
		if !f(r[i]) {
			return r[i:]
		}
	}
	return []rune{}
}

func TrimRight(r []rune, f func(rune) bool) []rune {
	for i := len(r) - 1; i >= 0; i-- {
		if !f(r[i]) {
			return r[:i+1]
		}
	}
	return []rune{}
}

func Trim(r []rune, f func(rune) bool) []rune {
	return TrimRight(TrimLeft(r, f), f)
}
