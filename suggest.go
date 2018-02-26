package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type hyposol struct {
	guess string
}

func bestword(sols *Bitmap, words []string, index map[string]*Bitmap) {
	possibleSolWords := bitmapToSlice(words, sols)
	bestelim := float64(0)
	bestguess := ""
	middle := len(possibleSolWords) * len(words) / 2
	for _, guess := range words {
		eliminates := 0
		count := 0
		for _, possible := range possibleSolWords {
			copySolMap := sols.Copy()
			if guess == possible {
				// forget about correct guesses...
				continue
			}
			hypoCount := countCorrect(guess, possible)
			addResult(guess, hypoCount, copySolMap, index, words)
			hypospace := bitmapToSlice(words, copySolMap)
			if len(hypospace) > 0 {
				count = count + len(hypospace)
				eliminates = eliminates + (len(possibleSolWords) - len(hypospace))
			}
		}
		fmt.Println(guess, math.Abs(float64(eliminates-middle)), middle)
		if math.Abs(float64(count-middle)) < math.Abs(bestelim-float64(middle)) {
			bestelim = float64(count)
			bestguess = guess
		}
	}
	fmt.Println("Best guess is", bestguess, "elimates", bestelim, "on average")
	for i := 0; i <= len(bestguess); i++ {
		copySolMap := sols.Copy()
		addResult(bestguess, i, copySolMap, index, words)
		hypospace := bitmapToSlice(words, copySolMap)
		fmt.Println("Correct:", i, "left", hypospace, len(hypospace))
	}
}

func countCorrect(a, b string) int {
	as := strings.Split(a, "")
	bs := strings.Split(b, "")

	sort.Strings(as)
	sort.Strings(bs)

	for i := 0; i < len(as); i++ {
		if i > len(bs)-1 || as[i] != bs[i] {
			return i
		}
	}
	return 0
}
