package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/adamvonbaron/gointerpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("hello %s! this is the monkey programming langauge!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
