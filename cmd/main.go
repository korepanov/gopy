package main

import (
	"fmt"
	"os"

	"github.com/korepanov/gopy/internal/program"
)

func main() {
	var p program.Program
	err := p.ReadProgram()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	p.WriteProgram()
}
