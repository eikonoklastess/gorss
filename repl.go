package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func startRepl(cfg *apiConfig) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcom to gorss a beautiful rss client/viewer for the cli\nEnter help for some help ;)")
	for {
		fmt.Print("> ")
		if reader.Scan() {
			input := cleanInput(reader.Text())
			if len(input) == 0 {
				continue
			}
			command, ok := getCommands()[input[0]]
			if ok {
				var cliArg *string
				if len(input) > 1 {
					arg := input[1]
					cliArg = &arg
				}
				err := command.callback(cfg, cliArg)
				if err != nil {
					fmt.Println(err.Error())
				}
			} else {
				fmt.Println("Unknown Command")
			}
		} else {
			break
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*apiConfig, *string) error
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    commandExit,
		},
		"clear": {
			name:        "clear",
			description: "Clear the screen",
			callback:    commandClear,
		},
		"follow": {
			name:        "follow {feed url}",
			description: "follow let you add a rss feed to your database for which gorss will aggregate data every few minutes when its open if it just opened gorss will only aggregate today's posts",
			callback:    commandFollow,
		},
	}
	return commands
}

func cleanInput(text string) []string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	outputs := strings.Split(output, " ")
	return outputs
}

func commandClear(*apiConfig, *string) error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}

func commandHelp(*apiConfig, *string) error {
	println()
	println("here are the available commands")
	println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	println()
	return nil
}

func commandExit(*apiConfig, *string) error {
	os.Exit(0)
	return nil
}
