package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}

type LocationAreaResponse struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(strings.ToLower(text))
	words := strings.Fields(text)
	return words
}

func commandExit(*Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range commandRegistry {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func FetchLocationAreas(url string) (*LocationAreaResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var locationAreaResponse LocationAreaResponse
	if err := json.NewDecoder(resp.Body).Decode(&locationAreaResponse); err != nil {
		return nil, err
	}

	return &locationAreaResponse, nil
}

func commandMap(config *Config) error {
	url := ""
	if config.Next != nil {
		url = *config.Next
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	locationAreas, err := FetchLocationAreas(url)
	if err != nil {
		return err
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	config.Next = locationAreas.Next
	config.Previous = locationAreas.Previous
	return nil
}

func commandMapBack(config *Config) error {
	if config.Previous == nil {
		fmt.Println("You're on the first page.")
		return nil
	}

	locationAreas, err := FetchLocationAreas(*config.Previous)
	if err != nil {
		return err
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	config.Next = locationAreas.Next
	config.Previous = locationAreas.Previous
	return nil
}

var commandRegistry map[string]cliCommand

func init() {
	commandRegistry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world",
			callback:    commandMapBack,
		},
	}
}

func main() {
	config := &Config{
		Next: getDefaultLocationAreaURL(),
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		words := cleanInput(input)

		if len(words) > 0 {
			command, exists := commandRegistry[words[0]]
			if exists {
				err := command.callback(config)
				if err != nil {
					fmt.Println("Error:", err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func getDefaultLocationAreaURL() *string {
	url := "https://pokeapi.co/api/v2/location-area/"
	return &url
}
