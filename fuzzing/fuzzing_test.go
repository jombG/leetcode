package fuzzing

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode/utf8"
)

func TestReverse1(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			in:  "Hello, world",
			out: "dlrow ,olleH",
		},
		{
			in:  " ",
			out: " ",
		},
		{
			in:  "1234",
			out: "4321",
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tt.out, Reverse1(tt.in))
		})
	}
}

//func FuzzReverse1(f *testing.F) {
//	cases := []string{"Hello, world", " ", "1234"}
//
//	for _, c := range cases {
//		f.Add(c)
//	}
//
//	f.Fuzz(func(t *testing.T, orig string) {
//		rev := Reverse1(orig)
//		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
//			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
//		}
//	})
//}

func FuzzReverse2(f *testing.F) {
	cases := []string{"Hello, world", " ", "1234"}

	for _, c := range cases {
		f.Add(c)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, err := Reverse2(orig)
		if err != nil {
			fmt.Println(err)
		}
		doubleRev, err := Reverse2(rev)
		if err != nil {
			fmt.Println(err)
		}
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
