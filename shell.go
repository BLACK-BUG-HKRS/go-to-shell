package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		//read keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		// handle execution of the input
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	// remove the new line character
	input = strings.TrimSuffix(input, "\n")
	// Prepare the command to execute
	cmd := exec.Command(input)
	// Set the correct output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	// execute the command and return the error
	return cmd.Run()
}
