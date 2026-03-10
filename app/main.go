package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// TODO: Uncomment the code below to pass the first stage
	for {
		fmt.Print("$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		command = strings.TrimSpace(command)

		if command == "exit" {
			break
		} else if strings.HasPrefix(command, "echo ") {
			command = strings.TrimPrefix(command, "echo ")
			fmt.Println(command)
		} else {
			fmt.Printf("%s: command not found\n", command)
		}
	}

}
