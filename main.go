package main 

import(
	"strings"
	"bufio" 
	"fmt"
	"os"
) 

func main() { 
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for {
		if scanner.Scan() == true { 
			input := scanner.Text() 
			strslice := cleanInput(input) 
			fmt.Print("Your command was: " + strslice[0])
		} 
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}  