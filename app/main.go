package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"slices"
	"strings"
)

func handleInput(input string) []string {
	// This regex matches an "argument" which is a sequence of:
	// Either a quoted block '...' OR one or more non-space characters [^\s]
	// The (?:...) logic ensures they are treated as one continuous token.
	re := regexp.MustCompile(`(?:'((?:''|[^'])*)'|[^\s'])+`)

	matches := re.FindAllString(input, -1)
	var result []string

	for _, match := range matches {
		// For each full match (e.g., script''hello), we manually
		// strip the outer quotes of any quoted sections and dissolve ''

		// 1. Handle inner doubled quotes first: '' -> (empty)
		temp := strings.ReplaceAll(match, "''", "")

		// 2. Strip any remaining single quotes
		temp = strings.ReplaceAll(temp, "'", "")

		result = append(result, temp)
	}

	return result
}

func main() {
	//array to maintain builtin types
	builtinTypes := []string{"type", "echo", "exit", "pwd"}

	//REPL
	for {
		fmt.Print("$ ")

		//Read the command and trim
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)

		//Split the command
		command := strings.Split(input, " ")[0]
		args := strings.Join(strings.Split(input, " ")[1:], " ")
		argv := handleInput(args)

		//Perform according to the input command
		switch command {
		case "exit":
			os.Exit(0)

		case "echo":
			output := strings.Join(argv, " ")
			fmt.Println(output)

		case "type":
			//Checks the command type
			//If is a builtin, displays that it is a builtin command
			//Else shows the absolute file path

			if slices.Contains(builtinTypes, argv[0]) {
				fmt.Printf("%s is a shell builtin\n", argv[0])
			} else {
				path, err := exec.LookPath(argv[0])
				if err != nil {
					fmt.Printf("%s: not found\n", argv[0])
				} else {
					fmt.Printf("%s is %s\n", argv[0], path)
				}
			}

		case "pwd":
			//prints the current working directory

			pwd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(pwd)

		case "cd":
			//Changes the current working directory

			if argv[0] == "~" {
				argv[0] = os.Getenv("HOME")
			}

			err := os.Chdir(argv[0])
			if err != nil {
				fmt.Printf("cd: %s: No such file or directory\n", argv[0])
			}

		default:
			//Checks the first argument to see if it is an executable
			//If it is an executable, calls the execute command with the provided arguments
			//Displays the output of executing the executable

			_, err := exec.LookPath(command)

			if err != nil {
				fmt.Printf("%s: command not found\n", command)
			} else {
				cmd := exec.Command(command, argv[0:]...)
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
