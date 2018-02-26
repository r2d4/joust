package main

import (
	"io/ioutil"
	"sort"
	"strings"
)

func load(file string, size int) ([]string, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	out := []string{}
	for _, line := range strings.Split(string(b), "\n") {
		for _, word := range strings.Split(line, " ") {
			if len(word) != size {
				continue
			}
			out = append(out, word)
		}
	}

	sort.Strings(out)
	return out, nil
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func reverseIndex(words []string) map[string]*Bitmap {
	out := map[string]*Bitmap{}
	for i, w := range words {
		s := strings.Split(w, "")
		sort.Strings(s)
		cs := combinations(s, -1)
		for _, c := range cs {
			if out[c] == nil {
				out[c] = NewBitmap(len(words))
			}
			out[c].Set(i)
		}
	}
	return out
}

func bitmapToSlice(words []string, b *Bitmap) []string {
	out := []string{}
	if b == nil {
		return out
	}
	for i, w := range words {
		if b.IsSet(i) {
			out = append(out, w)
		}
	}
	return out
}

func combinations(elems []string, size int) []string {
	out := []string{}
	n := len(elems)
	for num := 0; num < (1 << uint(n)); num++ {
		combination := []string{}
		for ndx := 0; ndx < n; ndx++ {
			if num&(1<<uint(ndx)) != 0 {
				combination = append(combination, elems[ndx])
			}
		}
		if len(combination) == 0 {
			continue
		}
		if len(combination) == size || size == -1 {
			out = append(out, strings.Join(combination, ""))
		}
	}
	return out
}
