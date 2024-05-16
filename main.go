package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	
	for {
		fmt.Print("sh> ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		input = strings.TrimSpace(input)
		parts := strings.Fields(input)
		if len(parts) == 0 {
			return
		}
		if parts[0] == "exit" {
			return
		}
		if parts[0] == "cd" {
			if len(parts) < 2 {
				fmt.Println("cd: missing argument")
				continue
			}
			if len(parts) > 2 {
				fmt.Println("cd: too many arguments")
				continue
			}
			path := parts[1]
			err := os.Chdir(path)
			if err != nil {	
				fmt.Printf("cd: %s: %s\n", path, err)
				continue
			}
		}
		cmd := exec.Command(parts[0], parts[1:]...)

		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error executing command:", err)
			continue
		}

		fmt.Print(string(output))
	}
}