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
	//array to maintain builtin types
	builtinTypes := []string{"type", "echo", "exit", "pwd"}

	//REPL
	for {
		fmt.Print("$ ")

		//Read the command and trim
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		command = strings.TrimSpace(command)

		//Split the command
		argv := strings.Split(command, " ")

		//Perform according to the input command
		switch argv[0] {
		case "exit":
			os.Exit(0)

		case "echo":
			fmt.Println(argv[1])

		case "type":
			//Checks the command type
			//If is a builtin, displays that it is a builtin command
			//Else shows the absolute file path

			if slices.Contains(builtinTypes, argv[1]) {
				fmt.Printf("%s is a shell builtin\n", argv[1])
			} else {
				path, err := exec.LookPath(argv[1])
				if err != nil {
					fmt.Printf("%s: not found\n", argv[1])
				} else {
					fmt.Printf("%s is %s\n", argv[1], path)
				}
			}

		case "pwd":
			//prints the current working directory
			fmt.Println(os.Getwd())

		default:
			//Checks the first argument to see if it is an executable
			//If it is an executable, calls the execute command with the provided arguments
			//Displays the output of executing the executable

			_, err := exec.LookPath(argv[0])

			if err != nil {
				fmt.Printf("%s: command not found\n", argv[0])
			} else {
				cmd := exec.Command(argv[0], argv[1:]...)
				output, err := cmd.CombinedOutput()
				if err != nil {
					fmt.Printf("%s: %s\n", argv[0], err)
				} else {
					fmt.Printf("%s", output)
				}
			}
		}
	}
}
