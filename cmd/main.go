package main

import (
	"gopy.com/internal"
)

func main() {
	var p internal.Program
	err := p.ReadProgram()

	if err != nil {

	}

	p.WriteProgram()
}
