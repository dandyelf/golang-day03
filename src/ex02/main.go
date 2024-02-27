package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	args := os.Args[1:]
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			cmd := exec.Command(args[0], append(args[1:], line)...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}
