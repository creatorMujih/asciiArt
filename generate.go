package main

import (
	"strings"
)

func generateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}

	input = strings.ReplaceAll(input, "\n", "\\n")

	word := strings.Split(input, "\\n")
	var result strings.Builder
	for index, value := range word {
		if value == "" {
			if index == len(word)-1 {
				continue
			}
			result.WriteByte('\n')
			continue
		}
		v := renderline(value, banner)
		for _, value := range v {
			result.WriteString(value)
			result.WriteString("\n")
		}
	}
	return result.String()
}
