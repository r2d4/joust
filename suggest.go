package main

import (
	"fmt"
	"math"
	"strings"
)

func bestword(sols *Bitmap, words []string, index map[string]*Bitmap, blacklist []string) string {
	bestguess := ""
	minBucket := float64(len(words))
	var skip bool
	for _, guess := range words {
		for _, b := range blacklist {
			if b == guess {
				skip = true
				break
			}
		}
		if skip {
			skip = false
			continue
		}
		maxBucket := float64(0)
		for i := 0; i <= len(guess); i++ {
			copySolMap := sols.Copy()
			addResult(guess, i, copySolMap, index, words)
			hypospace := bitmapToSlice(words, copySolMap)
			n := len(hypospace)
			if n == 0 {
				continue
			}
			maxBucket = math.Max(float64(n), maxBucket)
		}
		if maxBucket < minBucket {
			minBucket = maxBucket
			bestguess = guess
		}
	}
	fmt.Printf("The best guess is %s (d=%f)\n", bestguess, minBucket)
	for i := 0; i <= len(bestguess); i++ {
		copySolMap := sols.Copy()
		addResult(bestguess, i, copySolMap, index, words)
		hypospace := bitmapToSlice(words, copySolMap)
		n := len(hypospace)
		if n > 0 {
			fmt.Printf("(correct=%d) %d solutions remain\n", i, n)
			if n < 30 && n > 0 {
				fmt.Println("solutions:", hypospace)
			}
		}

	}
	return bestguess
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
