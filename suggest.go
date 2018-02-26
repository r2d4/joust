package main

import (
	"fmt"
	"math"
	"strings"
)

type hyposol struct {
	guess string
}

func bestword(sols *Bitmap, words []string, index map[string]*Bitmap) {
	middle := float64(len(bitmapToSlice(words, sols))) / 2
	bestguess := ""
	bestd := middle * 3
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
		fmt.Println("correct", i, n)
	}
}

func countCorrect(a, b string) int {
	a = strings.ToUpper(a)
	b = strings.ToUpper(b)
	as := strings.Split(a, "")
	bs := strings.Split(b, "")

	amap := map[string]int{}
	bmap := map[string]int{}

	match := 0
	for i := 0; i < len(as); i++ {
		amap[as[i]] = amap[as[i]] + 1
	}

	for i := 0; i < len(bs); i++ {
		bmap[bs[i]] = bmap[bs[i]] + 1
	}

	for k, v := range amap {
		min := math.Min(float64(v), float64(bmap[k]))
		match = match + int(min)
	}

	return match
}
