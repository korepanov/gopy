package main

import (
	"fmt"
	"os"

	"github.com/korepanov/gopy/internal/cerrors"
	"github.com/korepanov/gopy/internal/flagparse"
	"github.com/korepanov/gopy/internal/program"
)

func main() {
	err := flagparse.ProcessFlags()

	if err == cerrors.ErrHelp {
		os.Exit(0)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s : %s\n", cerrors.ErrCompile, err)
		os.Exit(1)
	}

	var p program.Program
	err = p.ReadProgram()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s : %s\n", cerrors.ErrCompile, err)
		os.Exit(1)
	}

	p.WriteProgram()
}
