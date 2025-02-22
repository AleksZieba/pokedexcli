package main 

import(
	"strings"
	"bufio" 
	"fmt"
	"os" 
	//"errors"
) 

func main() { 
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ") 
	commands = map[string]cliCommand{
		"help": {
			name:			"help",
			description:	"Displays a help message",
			callback:		commandHelp,
		},
		"exit": {
			name:			"exit", 
			description:	"Exit the Pokedex",
			callback:		commandExit, 
		},  
	}
	for {
		if scanner.Scan() == true { 
			input := scanner.Text() 
			strslice := cleanInput(input) 
			if cmd, ok := commands[strslice[0]]; ok {
				cmd.callback() 
			} else {
				fmt.Println("Unknown command")
			}
		} 
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}  

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0) 
	return nil
} 

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, command := range commands {
		fmt.Println(command.name + ": " + command.description)
	} 
	return nil 
}

type cliCommand struct {
	name 		string 
	description	string 
	callback 	func() error 
} 

var commands map[string]cliCommand
