package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/juanjuanzero/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the monkey language! \n", user.Username)
	fmt.Printf("Type the commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
