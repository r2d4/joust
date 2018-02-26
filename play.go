package main

import (
	"sort"
	"strings"
)

type round struct {
	word    string
	correct int
}

func play(words []string, rs ...round) *Bitmap {
	m := reverseIndex(words)
	sols := NewBitmapAllOn(len(words))
	for _, r := range rs {
		r.word = strings.ToUpper(r.word)
		addResult(r.word, r.correct, sols, m, words)
	}
	return sols
}

func addResult(word string, correct int, sols *Bitmap, index map[string]*Bitmap, words []string) {
	s := strings.Split(word, "")
	sort.Strings(s)
	roundSols := NewBitmap(len(words))
	subsets := combinations(s, correct)
	if correct == 0 {
		subsets = combinations(s, -1)
	}

	for _, set := range subsets {
		roundSols.Union(index[set])
	}

	if correct == 0 {
		roundSols.Complement()
	}

	sols.Intersect(roundSols)
}
