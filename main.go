package main 

import(
	"strings"
) 

func main() {
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}