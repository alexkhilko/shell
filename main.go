package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"io"
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
		commands := strings.Split(input, "|")
		// Initial input is from the standard input
		var inputPipe io.Reader = os.Stdin
		var cmds []*exec.Cmd
		for i, command := range commands {
			command = strings.TrimSpace(command)
			parts := strings.Fields(command)

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
			cmd.Stdin = inputPipe

			if i < len(commands) - 1 {
				outputPipe, err := cmd.StdoutPipe()
				if err != nil {
					fmt.Printf("Error creating stdout pipe: %s\n", err)
					break
				}
				inputPipe = outputPipe
			} else {
				cmd.Stdout = os.Stdout
			}

			if err := cmd.Start(); err != nil {
				fmt.Printf("Error starting command: %s\n", err)
				break
			}
			cmds = append(cmds, cmd)
		}
		for _, cmd := range cmds {
			if err := cmd.Wait(); err != nil {
				fmt.Printf("Error waiting for command: %s, %v\n", strings.Join(cmd.Args, " "), err)
				if exitErr, ok := err.(*exec.ExitError); ok {
					fmt.Println("Exit status:", exitErr.ExitCode())
				}
			}
		}
	}
}