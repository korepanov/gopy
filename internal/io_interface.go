package internal

import (
	"bufio"
	"os"
)

const chunkSize = 1024

type Command struct {
	Input  string
	Tokens []string
}

func (c *Command) NextCommand() {

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	c.Input = s.Text()
}
