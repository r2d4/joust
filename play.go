package main

import (
	"fmt"
	"sort"
	"strings"
)

type round struct {
	word    string
	correct int
}

func cliPlay(words []string) {
	m := reverseIndex(words)
	sols := NewBitmapAllOn(len(words))

	fmt.Println("pick a word")
	var actual string
	fmt.Scanln(&actual)
	fmt.Println("playing with word", actual)
	rounds := 0
	for {
		fmt.Println("finding best word...")
		bestword(sols, words, m, []string{})
		fmt.Println("pick a word")
		var pick string
		fmt.Scanln(&pick)
		rounds = rounds + 1
		if pick == actual {
			break
		}
		c := countCorrect(actual, pick)
		fmt.Println("CORRECT", c)
		playround(words, round{pick, c}, m, sols)
		fmt.Println("sols", bitmapToSlice(words, sols))
	}

	fmt.Println("won, rounds", rounds)
}

func playRounds(words []string, rs ...round) *Bitmap {
	m := reverseIndex(words)
	sols := NewBitmapAllOn(len(words))
	for _, r := range rs {
		r.word = strings.ToUpper(r.word)
		addResult(r.word, r.correct, sols, m, words)
	}
	return sols
}

func playround(words []string, r round, idx map[string]*Bitmap, sol *Bitmap) {
	r.word = strings.ToUpper(r.word)
	addResult(r.word, r.correct, sol, idx, words)
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
