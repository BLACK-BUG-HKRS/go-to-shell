package shell

import (
	"bufio"
	"os"
)

func shell() {
	reader := bufio.NewReader(os.Stdin)
	for {
		//read keyboard input
		input, err := reader.ReadString('\n')
	}
}
