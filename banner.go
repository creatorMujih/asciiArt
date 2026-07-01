package main

import (
	"fmt"
	"os"
	"strings"
)

func loadBanner(filename string) (map[rune][]string, error) {
	font := make(map[rune][]string)
	data, err := os.ReadFile(filename)

	if len(data) == 0 {
		return nil, fmt.Errorf("Empty")
	}

	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) < 856 {
		return nil, fmt.Errorf("Incomplete")
	}

	for c := ' '; c <= '~'; c++ {
		start := (int(c) - 32) * 9
		font[c] = lines[start+1 : start+9]
	}
	return font, err
}
