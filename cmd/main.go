package main

import (
	"fmt"
	"os"

	"github.com/korepanov/gopy/internal/cerrors"
	"github.com/korepanov/gopy/internal/dispatcher"
)

func main() {

	err := dispatcher.Compile()

	if err == cerrors.ErrHelp {
		os.Exit(0)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
