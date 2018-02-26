package main

import (
	"fmt"
	"sort"
	"strings"
)

func stats(sols *Bitmap, words []string) {
	solWords := bitmapToSlice(words, sols)
	m := map[string]int{}
	for _, sol := range solWords {
		combos := combinations(strings.Split(sol, ""), -1)
		for _, c := range combos {
			m[c]++
		}
	}

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for _, kv := range ss {
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)
	}
}
