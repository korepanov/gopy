package internal

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

const chunkSize = 1024

type Command struct {
	Input  string
	Tokens []string
}

func (c *Command) NextCommand() error {

	c.Input = ""
	c.Tokens = nil

	r := bufio.NewReader(os.Stdin)

	b := make([]byte, chunkSize)
	for {
		_, err := r.Read(b)

		if err != nil {
			return err
		}

		c.Input += string(b)

		endIndex := bytes.Index(b, []byte("\n"))

		if endIndex != -1 {
			c.Input = c.Input[:strings.Index(c.Input, "\n")]

			if len(b) > endIndex+1 {
				w := bufio.NewWriter(os.Stdin)
				w.Write(b[endIndex+1:])
			}

			break
		}
	}

	return nil
}
