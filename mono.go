package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Mono")
	fmt.Println("---------------------")

	currentDir, err := os.Getwd()

	if err != nil {
		fmt.Println("Cannot get current working directory")
	}

	for {
		fmt.Println(currentDir)
		fmt.Print("mono>>")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if strings.HasPrefix(input, "$") {
			input = strings.TrimPrefix(input, "$")
			//trim again
			input = strings.TrimSpace(input)

			//check for exit
			if input == "exit" {
				fmt.Println("Exiting CLI...")
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
