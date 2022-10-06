package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func shell() {
	reader := bufio.NewReader(os.Stdin)
	for {
		//read keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
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
}
