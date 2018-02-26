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
	middle := float64(len(bitmapToSlice(words, sols))) / 2
	bestguess := ""
	bestd := middle
	for _, guess := range words {
		d := float64(0)
		for i := 0; i <= len(guess); i++ {
			copySolMap := sols.Copy()
			addResult(guess, i, copySolMap, index, words)
			hypospace := bitmapToSlice(words, copySolMap)
			n := len(hypospace)
			if n == 0 {
				continue
			}
			d = d + math.Abs(middle-float64(n))
		}
		if d < bestd {
			bestguess = guess
			bestd = d
		}
	}
	fmt.Println("best guess", bestguess, "d", bestd)
	for i := 0; i <= len(bestguess); i++ {
		copySolMap := sols.Copy()
		addResult(bestguess, i, copySolMap, index, words)
		hypospace := bitmapToSlice(words, copySolMap)
		n := len(hypospace)
		fmt.Println("correct", i, hypospace, n)
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
