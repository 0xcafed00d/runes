package runes

import (
	"testing"
	"unicode"
)

type testtrim struct {
	data     string
	expected string
	f        func([]rune, func(rune) bool) []rune
}

var testdata_trim = []testtrim{
	{"123", "123", TrimLeft},
	{" 123", "123", TrimLeft},
	{"  123", "123", TrimLeft},
	{"   ", "", TrimLeft},
	{"", "", TrimLeft},

	{"456", "456", TrimRight},
	{"456 ", "456", TrimRight},
	{"456  ", "456", TrimRight},
	{"   ", "", TrimRight},
	{"", "", TrimRight},

	{"789", "789", Trim},
	{"789 ", "789", Trim},
	{"789  ", "789", Trim},
	{" 789", "789", Trim},
	{"  789", "789", Trim},
	{" 789 ", "789", Trim},
	{"  789  ", "789", Trim},
	{"   ", "", Trim},
	{"", "", Trim},
}

func TestTrim(t *testing.T) {
	for _, tst := range testdata_trim {
		got := tst.f([]rune(tst.data), unicode.IsSpace)
		if string(got) != tst.expected {
			t.Fatalf("Trim Test: [%s] Expected: [%s] Got: [%s]",
				tst.data, tst.expected, string(got))
		}
	}
}

type testinsertdel struct {
	data            string
	insert_expected string
	delete_expected string
	index           int
}

var testdata_id = []testinsertdel{
	{"123", "x123", "23", 0},
	{"123", "1x23", "13", 1},
	{"123", "12x3", "12", 2},
	{"123", "123x", "ignore", 3},
}

func TestInsertDelete(t *testing.T) {
	for i, tst := range testdata_id {
		data := CloneSlice([]rune(tst.data))

		if tst.insert_expected != "ignore" {
			res := string(InsertAt(data, 'x', tst.index))
			if res != tst.insert_expected {
				t.Fatalf("Insert Rune Test: [%d] Expected: [%s] Got: [%s]",
					i, tst.insert_expected, res)
			}
		}

		data = CloneSlice([]rune(tst.data))
		if tst.delete_expected != "ignore" {
			res := string(DeleteAt(data, tst.index))
			if res != tst.delete_expected {
				t.Fatalf("Delete Rune Test: [%d] Expected: [%s] Got: [%s]",
					i, tst.delete_expected, res)
			}
		}
	}
}

type testinsertslice struct {
	data     string
	expected string
	index    int
}

var testdata_is = []testinsertslice{
	{"123", "xxx123", 0},
	{"123", "1xxx23", 1},
	{"123", "12xxx3", 2},
	{"123", "123xxx", 3},
}

func TestInsertSlice(t *testing.T) {
	for i, tst := range testdata_is {
		data := CloneSlice([]rune(tst.data))
		res := string(InsertSliceAt(data, []rune("xxx"), tst.index))
		if res != tst.expected {
			t.Fatalf("Insert Rune Slice Test: [%d] Expected: [%s] Got: [%s]",
				i, tst.expected, res)
		}
	}
}

type testcut struct {
	data         string
	expected     string
	expected_cut string
	index        int
	count        int
}

var testdata_cut = []testcut{
	{"123456", "123456", "", 0, 0},
	{"123456", "23456", "1", 0, 1},
	{"123456", "3456", "12", 0, 2},
	{"123456", "456", "123", 0, 3},
	{"123456", "56", "1234", 0, 4},
	{"123456", "6", "12345", 0, 5},
	{"123456", "", "123456", 0, 6},
}

func TestCutSlice(t *testing.T) {
	for i, tst := range testdata_cut {
		data := CloneSlice([]rune(tst.data))
		res, cut := CutSliceAt(data, tst.index, tst.count)

		if string(res) != tst.expected || string(cut) != tst.expected_cut {
			t.Fatalf("Insert Rune Slice Test: [%d] Expected: [%s,%s] Got: [%s,%s]",
				i, tst.expected, tst.expected_cut, string(res), string(cut))
		}
	}
}
