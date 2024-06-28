package internal

import (
	"bufio"
	"fmt"
	"os"
)

func (p *Program) ReadProgram() error {
	s := bufio.NewScanner(os.Stdin)

	for p.nextCommand(s) {
	}

	return nil
}

func (p *Program) LexicalAnalyze() error {
	for idx, c := range p.input {
		err := c.lexicalAnalyze()
		return fmt.Errorf("%q : %q : %d\n%q", ErrCompile, err, idx+1)
	}
	return fmt.Errorf("%q : %q", ErrReadProgram, err)
}

func (p *Program) WriteProgram() {
	for _, command := range p.input {
		for _, token := range command.tokens {
			fmt.Print(token.lex + " ; ")
		}
		fmt.Println()
	}
}

// reads next command in the command input
func (p *Program) nextCommand(s *bufio.Scanner) bool {
	if !s.Scan() {
		return false
	}
	var c command
	c.input = s.Text()
	p.input = append(p.input, c)

	return true
}
