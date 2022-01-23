package main

import (
	"fmt"
	"testing"
)

func TestSmbGo(t *testing.T) {
	var tests = []struct {
		Namefs string
		Datafs string
	}{
		{"", "\n"},
		{"", ""},
		{"\t", "one\ttwo\tthree\n"},
		{"Data for test", "Number 99999 to data test"},
		{"Yes, no", "No, or, yes"},
	}

	var prevNamefs string
	for _, test := range tests {
		if test.Namefs != prevNamefs {
			fmt.Printf("\n%s\n", test.Namefs)
			prevNamefs = test.Namefs
		}
	}

	var prevDatafs string
	for _, test := range tests {
		if test.Datafs != prevDatafs {
			fmt.Printf("\n%s\n", test.Datafs)
			prevDatafs = test.Datafs
		}
	}
}
