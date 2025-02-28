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
	"time" 
	"math/rand"
	"github.com/AleksZieba/pokedexcli/internal/pokecache"
) 

var cmd *cliCommand 

var commands map[string]*cliCommand 

type cliCommand struct {
	name 		string 
	description	string   
	callback 	func() error 
	parameters 	string
} 

func initCommands() {
	commands = map[string]*cliCommand{
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
		"explore": {
			name:			"explore",
			description:	"Displays the Pokemon in a location", 
			callback:		commandExplore,
		}, 
		"catch": {
			name:			"catch",
			description:	"Attempts to catch a given Pokemon",
			callback:		commandCatch,
		}, 
		"inspect": {
			name: 			"inspect", 
			description:	"Gives details about a captured Pokemon", 
			callback:		commandInspect, 
		},
	}
}

func main() { 
	initCommands()
	
	repl()
}

func repl() { 
	scanner := bufio.NewScanner(os.Stdin) 
	for { 
		fmt.Print("Pokedex > ")
		if scanner.Scan() == true { 
			input := scanner.Text() 
			strslice := cleanInput(input) 
			if command, ok := commands[strslice[0]]; ok {
				if command.name == "explore" || command.name == "catch" || command.name == "inspect" {
					if len(strslice) < 2 {
						fmt.Println("Missing the location or Pokemon name, please try again") 
						repl()
					} else {
					command.parameters = strslice[1] 
					cmd = command
					}
				}
				command.callback() 
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
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, command := range commands {
		fmt.Println("\n" + command.name + ": " + command.description)
	} 
	return nil 
} 

func commandMap() error {
	finalIndex += 20 
	mapIndex := finalIndex - 20 
	for mapIndex <= finalIndex {
		reqURL := "https://pokeapi.co/api/v2/location-area/" + strconv.FormatUint(mapIndex, 10)
		if body, ok := cache.Entries[reqURL]; ok {
			
			location := location{} 
			err := json.Unmarshal(body.Val, &location)
			if err != nil {
				errors.New("json.Unmarshal() failed")
			}
			fmt.Println(location.Name) 
		} else {
			res, err := http.Get(reqURL)
			if err != nil {
				errors.New("Get request failed")
			}

			body, err := io.ReadAll(res.Body)
			if err != nil {
				errors.New("io.ReadAll() failed")
			} //body = []bytes
			cache.Add(reqURL, body)
			
			location := location{} 
			err = json.Unmarshal(body, &location)
			if err != nil {
				errors.New("json.Unmarshal() failed")
			}
			fmt.Println(location.Name) 
			defer res.Body.Close()
		}
	
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
		reqURL := "https://pokeapi.co/api/v2/location-area/" + strconv.FormatUint(mapIndex, 10)
		if body, ok := cache.Entries[reqURL]; ok {
			location := location{} 
			err := json.Unmarshal(body.Val, &location)
			if err != nil {
				errors.New("json.Unmarshal() failed")
		}
		fmt.Println(location.Name) 
		} else {
			res, err := http.Get("https://pokeapi.co/api/v2/location-area/" + strconv.FormatUint(mapIndex, 10)) 
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
		} 

		mapIndex++
	}
	return nil
} 

func commandExplore() error {
	reqURL := "https://pokeapi.co/api/v2/location-area/" + cmd.parameters
	if body, ok := cache.Entries[reqURL]; ok {
			
		location := location{} 
		err := json.Unmarshal(body.Val, &location)
		if err != nil {
			errors.New("json.Unmarshal() failed")
		}
		//fmt.Println(location.Name) 
		fmt.Printf("Exploring %s...\nFound Pokemon:\n", location.Name) 
		for _, pokemon := range(location.PokemonEncounters) {
			//pokemon = strings.ReplaceAll(pokemon, "}", "")
			fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
		}
	} else {
		res, err := http.Get(reqURL)
		if err != nil {
			errors.New("Get request failed")
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			errors.New("io.ReadAll() failed")
		} //body = []bytes
		cache.Add(reqURL, body)
		
		location := location{} 
		err = json.Unmarshal(body, &location)
		if err != nil {
			errors.New("json.Unmarshal() failed")
		}
		fmt.Printf("Exploring %s...\nFound Pokemon:\n", location.Name) 
		for _, pokemon := range(location.PokemonEncounters) {
			//pokemon = strings.ReplaceAll(pokemon, "}", "")
			fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
		}
		defer res.Body.Close()
	}

	return nil 
}

/* type cmdParameters struct {
	name		string 
	id			int 
} */

type location struct {
	ID      				int     `json:"id"`
	Name					string	`json:"name"` // should this be capitalized???
	PokemonEncounters    	[]PokemonEncounters    `json:"pokemon_encounters"`
} 

type Pokemon struct {
	Name string `json:"name"`
//	URL  string `json:"url"`
} 

type PokemonEncounters struct {
	Pokemon        Pokemon          `json:"pokemon"`
//	VersionDetails []VersionDetails `json:"version_details"`
}

var finalIndex uint64 = 1 

var cache *pokecache.Cache = pokecache.NewCache(15 * time.Second) 

var pokedex = make(map[string]PokemonDetails, 0) 

type PokemonDetails struct {
	Name                   string        `json:"name"`
	BaseExperience         int           `json:"base_experience"`
	Height                 int           `json:"height"`
	Weight                 int           `json:"weight"`
//	Stats                  []Stats       `json:"stats"`
	Types                  []Types       `json:"types"`
} 

type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
} 

type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}


//var cmd2 *cliCommand 

func commandCatch() error {
	fmt.Printf("Throwing a Pokeball at %s...", cmd.parameters)
	reqURL := "https://pokeapi.co/api/v2/pokemon/" + cmd.parameters //+"/"
	res, err := http.Get(reqURL) 
		if err != nil {
			errors.New("Get request failed")
		}
		defer res.Body.Close() //should work here, refactor others??

		body, err := io.ReadAll(res.Body)
		if err != nil {
			errors.New("io.ReadAll() failed")
		}

		pokemon := PokemonDetails{} 
		err = json.Unmarshal(body, &pokemon)
		if err != nil {
			errors.New("json.Unmarshal() failed")
		}
		chance := int(rand.Float32() * 100)
		if chance < pokemon.BaseExperience {
			fmt.Printf("\n%s escaped!\n", pokemon.Name)
		} else {
			if _, ok := pokedex[pokemon.Name]; !ok {
				pokedex[pokemon.Name] = pokemon
			}
			fmt.Printf("\n%s was caught!\n", pokemon.Name) 
		}
		
	return nil
} 

func commandInspect() error {
	if pokemon, ok := pokedex[cmd.parameters]; !ok {
		fmt.Println("This Pokemon is not in the Pokedex yet.")
	} else {
		fmt.Printf("\nName: %s", pokemon.Name) 
		fmt.Printf("\nHeight: %d", pokemon.Height) 
		fmt.Printf("\nWeight: %d\n", pokemon.Weight)
		fmt.Println("Types:")
		for _, poketype := range(pokemon.Types) {
			fmt.Printf("  - %s\n", poketype.Type.Name)
		}
	}
	return nil 
}