package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	Shell       = []string{"/bin/sh", "-c"}
	Panic       = true
	Trace       = false
	TracePrefix = "+"

	exit = os.Exit
)
var Tee io.Writer

func assert(err error) {
	if err != nil {
		panic(err)
	}
}

func Path(parts ...string) string {
	return filepath.Join(parts...)
}

func PathTemplate(parts ...string) func(...interface{}) string {
	return func(values ...interface{}) string {
		return fmt.Sprintf(Path(parts...), values...)
	}
}

func Quote(arg string) string {
	return fmt.Sprintf("'%s'", strings.Replace(arg, "'", "'\\''", -1))
}

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
	arrInputs := strings.Fields(input)

	switch arrInputs[0] {
	case "exit":
		//adding custom commands
		os.Exit(0)
	}

	// Prepare the command to execute
	cmd := exec.Command(arrInputs[0], arrInputs[1:]...)

	// Set the correct output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// execute the command and return the error
	return cmd.Run()
}
