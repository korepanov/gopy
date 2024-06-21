package internal

import (
	"bufio"
	"os"
)

func (p *Program) nextCommand() {

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	var c Command
	c.Input = s.Text()
	c.
		p.InputProgram = p.InputProgram.append(p.InputProgram)
}
