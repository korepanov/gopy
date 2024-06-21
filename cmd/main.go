package main

import (
	"fmt"

	"gopy.com/internal"
)

func main() {
	var c internal.Command

	for {
		err := c.NextCommand()
		if err != nil {
			panic(err)
		}
		fmt.Println(c.Input)
	}

}
