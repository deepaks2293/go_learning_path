package main

type cliCommand struct {
	name        string
	description string
	callback    string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    "ok",
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    "yes",
		},
	}
}

// func main() {

// 	// Function with clicommand (commandName) passed here

// 	i := map[string]cliCommand{
// 		"help": {
// 			name:        "help",
// 			description: "Displays a help message",
// 			callback:    "ok",
// 		},
// 		"exit": {
// 			name:        "exit",
// 			description: "Exit the Pokedex",
// 			callback:    "yes",
// 		},
// 	}
// 	x, y := i["exit"]

// 	// two variable to check command exist and get value
// 	exists, command := getCommands()["help"]
// 	fmt.Println(command)
// 	fmt.Println(exists)
// 	fmt.Println(x)
// 	fmt.Println(y)
// }
