package shell

import (
	"bufio"
	"os"
)

func shell() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
}
