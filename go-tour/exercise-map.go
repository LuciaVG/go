package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	mp := make(map[string]int)
	words := strings.Fields(s)
	for _, v := range words {
		mp[v]++
	}
	return mp
}

func main() {
	wc.Test(WordCount)
}
