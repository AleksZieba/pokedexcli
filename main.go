package main 

import(
	"strings"
	"bufio" 
	"fmt"
	"os" 
	"net/http" 
	"errors"
	"strconv"
	"io"
	"encoding/json" 
	"pokedexcli/internal/pokecache"
) 

func main() { 
	cache := NewCache(5 * time.Second)
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
		"map": {
			name:			"map", 
			description:	"Displays next 20 locations", 
			callback: 		commandMap,
		}, 
		"mapb": {
			name: 			"mapb",
			description:	"Displays previous 20 locations",
			callback:		commandMapB,
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

func commandMap() error {
	finalIndex += 20 
	mapIndex := finalIndex - 20
	for mapIndex <= finalIndex {
		res, err := http.Get("https://pokeapi.co/api/v2/location-area/" + strconv.FormatUint(uint64(mapIndex), 10)) 
		if err != nil {
			errors.New("Get request failed")
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			errors.New("io.ReadAll() failed")
		}

		location := location{} 
		err = json.Unmarshal(body, &location)
		if err != nil {
			errors.New("json.Unmarshal() failed")
		}
		fmt.Println(location.Name) 

		defer res.Body.Close()
		mapIndex++
	}
	return nil
} 

func commandMapB() error { 
	if finalIndex < 22 {
		fmt.Println("you're on the first page")
	}
	finalIndex -= 20 
	mapIndex := finalIndex - 20
	for mapIndex <= finalIndex {
		res, err := http.Get("https://pokeapi.co/api/v2/location-area/" + strconv.FormatUint(uint64(mapIndex), 10)) 
		if err != nil {
			errors.New("Get request failed")
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			errors.New("io.ReadAll() failed")
		}

		location := location{} 
		err = json.Unmarshal(body, &location)
		if err != nil {
			errors.New("json.Unmarshal() failed")
		}
		fmt.Println(location.Name) 

		defer res.Body.Close()
		mapIndex++
	}
	return nil
}

type cliCommand struct {
	name 		string 
	description	string 
	callback 	func() error 
} 

type location struct {
	Name	string	`json:"name"`
}

var commands map[string]cliCommand

var finalIndex uint16 = 1