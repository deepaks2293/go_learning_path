package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl() {
	// To take argument
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		//Passing string to clean input function
		words := cleanInput(reader.Text())
		// checking len of words is not 0 if 0 than got to ---> fmt.Print("Pokedex > ")
		if len(words) == 0 {
			continue
		}

		// Remove string from words 0 element
		commandName := words[0]

		//Calling getCommands function input as command name
		command, exists := getCommands()[commandName]
		//command variable hold struct vaulue and exists hold bool value (true or false)
		if exists {
			// Check err and call callback() function in struct
			err := command.callback()
			// if err print err and back to ---> fmt.Print("Pokedex > ")
			if err != nil {
				fmt.Println(err)
			}
			continue
			//if  command not found in struct print else block
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

// take input as string and return output as list of string
func cleanInput(text string) []string {
	//Change text to lower case
	output := strings.ToLower(text)
	// Fields used to split and return as list
	words := strings.Fields(output)
	return words
}

// Creating struct for command
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Function with clicommand (commandName) passed here and return the key if exit in struct getCommands()["help"] confused refer main.go
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    CommandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    CommandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations areas in the Pokemon world",
			callback:    CommandExit,
		},
	}
}
