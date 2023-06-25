package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execCommand(input string) error {
	// Remove the end line
	input = strings.TrimSuffix(input, "\n")

	// split the commands
	args := strings.Split(input, " ")
	cmd := exec.Command(args[0])

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	fmt.Print(cmd.Stderr)
	return cmd.Run() // Returns the output.
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error given from reading file : %v", err)
		}

		execCommand(input)
	}
}
