package main

func renderline(text string, banner map[rune][]string) []string {
	result := []string{}
	new := ""

	for i := range 8 {
		for _, c := range text {
			new += banner[c][i]
		}
		result = append(result, new)
		new = ""
	}
	return result
}
