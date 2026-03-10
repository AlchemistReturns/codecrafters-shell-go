package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func main() {
	//REPL
	for {
		fmt.Print("$ ")

		//Read the command and trim
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		command = strings.TrimSpace(command)

		//array to maintain builtin types
		builtinTypes := []string{"type", "echo", "exit"}

		if command == "exit" {
			break

		} else if strings.HasPrefix(command, "echo ") {
			//echoes a command
			command = strings.TrimPrefix(command, "echo ")
			fmt.Println(command)

		} else if strings.HasPrefix(command, "type ") {
			//Checks the command type
			//If is a builtin, displays that it is a builtin command
			//Else shows the absolute file path
			command = strings.TrimPrefix(command, "type ")

			if slices.Contains(builtinTypes, command) {
				fmt.Printf("%s is a shell builtin\n", command)
			} else {
				path, err := exec.LookPath(command)
				if err != nil {
					fmt.Printf("%s: not found\n", command)
				} else {
					fmt.Printf("%s is %s\n", command, path)
				}
			}

		} else {
			//splits the command line arguments into an argv slice
			//checks the first argument to see if it is an executable
			//If it is an executable, calls the execute command with the provided arguments
			//Displays the output of executing the executable
			argv := strings.Split(command, " ")
			_, err := exec.LookPath(argv[0])

			if err != nil {
				fmt.Printf("%s: command not found\n", command)
			} else {
				cmd := exec.Command(argv[0], argv[1:]...)
				output, err := cmd.CombinedOutput()
				if err != nil {
					fmt.Printf("%s: %s\n", command, err)
				} else {
					fmt.Printf("%s", output)
				}
			}

		}
	}
}
