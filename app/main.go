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

// The following function was generated with the help of AI (Gemini)
func handleInput(input string) []string {
	// This regex matches an "argument" (one or more parts glued together).
	// Part 1: '((?:''|[^'])*)'  (Single quoted)
	// Part 2: "((?:""|[^"])*)"  (Double quoted)
	// Part 3: [^\s'"]+          (Plain text)
	reArg := regexp.MustCompile(`(?:'((?:''|[^'])*)'|"((?:""|[^"])*)"|[^\s'"]+)+`)

	// This sub-regex is used to find the quoted parts WITHIN a single argument
	reParts := regexp.MustCompile(`'((?:''|[^'])*)'|"((?:""|[^"])*)"`)

	matches := reArg.FindAllString(input, -1)
	var result []string

	for _, match := range matches {
		// We process the match by looking ONLY for the quoted segments
		// and replacing them with their stripped/dissolved content.
		processed := reParts.ReplaceAllStringFunc(match, func(m string) string {
			if strings.HasPrefix(m, "'") {
				// It's a single-quoted block: strip outer ' and dissolve ''
				content := m[1 : len(m)-1]
				return strings.ReplaceAll(content, "''", "")
			} else {
				// It's a double-quoted block: strip outer " and dissolve ""
				content := m[1 : len(m)-1]
				return strings.ReplaceAll(content, `""`, "")
			}
		})

		// Note: Plain text like script's is left untouched by ReplaceAllStringFunc
		// because it doesn't match the reParts pattern.
		result = append(result, processed)
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
