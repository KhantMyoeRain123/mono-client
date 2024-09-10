package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/chzyer/readline"
)

type SimpleCompleter struct {
	commands []string
}

// Complete provides autocompletions based on the input
func (c *SimpleCompleter) Do(line []rune, pos int) ([][]rune, int) {
	var newLine [][]rune

	for _, cmd := range c.commands {
		sline := string(line)
		sline = strings.TrimPrefix(sline, "$")
		sline = strings.TrimSpace(sline)
		trimmed_cmd := strings.TrimPrefix(cmd, sline)
		newLine = append(newLine, []rune(trimmed_cmd))
	}
	return newLine, pos
}

func main() {
	completer := &SimpleCompleter{
		commands: []string{
			"exit",
			"cd",
			"ls",
			"echo",
		},
	}

	reader, err := readline.NewEx(&readline.Config{
		Prompt:          ">> ",
		HistoryFile:     "history.txt",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
	})

	if err != nil {
		fmt.Println("Error creating readline instance")
		return
	}

	fmt.Println("Welcome to Mono")
	fmt.Println("---------------------")

	currentDir, err := os.Getwd()

	if err != nil {
		fmt.Println("Cannot get current working directory")
	}

	for {
		fmt.Println(currentDir)
		input, _ := reader.Readline()
		input = strings.TrimSpace(input)
		if strings.HasPrefix(input, "$") {
			input = strings.TrimPrefix(input, "$")
			//trim again
			input = strings.TrimSpace(input)

			//check for exit
			if input == "exit" {
				fmt.Println("Exiting...")
				break
			}
			//check for cd
			if strings.HasPrefix(input, "cd ") {
				dir := strings.TrimPrefix(input, "cd ")
				dir = strings.TrimSpace(dir)
				err := os.Chdir(dir)
				if err != nil {
					fmt.Println(err)
				}
				currentDir, err = os.Getwd()
				if err != nil {
					fmt.Println("Cannot get current working directory")
				}
				continue
			}

			// Execute the input as a shell command
			err := executeShellCommand(input)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}

	}

}

func executeShellCommand(cmd string) error {
	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		return nil
	}
	command := exec.Command(parts[0], parts[1:]...)

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	return command.Run()
}
