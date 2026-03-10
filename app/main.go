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
	for {
		fmt.Print("$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		command = strings.TrimSpace(command)

		builtinTypes := []string{"type", "echo", "exit"}

		if command == "exit" {
			break
		} else if strings.HasPrefix(command, "echo ") {
			command = strings.TrimPrefix(command, "echo ")
			fmt.Println(command)
		} else if strings.HasPrefix(command, "type ") {
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
			fmt.Printf("%s: command not found\n", command)

		}
	}
}
