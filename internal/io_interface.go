package internal

import (
	"bufio"
	"fmt"
	"os"
)

func (p *Program) ReadProgram() {
	s := bufio.NewScanner(os.Stdin)

	for p.nextCommand(s) {
		p.input[len(p.input)-1].lexicalAnalyze()
	}
}

func (p *Program) WriteProgram() {
	for _, command := range p.input {
		for _, token := range command.tokens {
			fmt.Print(token.lexeme + " ; ")
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
