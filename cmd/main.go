package main

import (
	"fmt"

	"gopy.com/internal"
)

func main() {
	var c internal.Command

	for {
		c.NextCommand()
		fmt.Println(c.Input)
	}

}
