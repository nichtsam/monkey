package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/nichtsam/monkey/repl"
)

func main() {
	r := repl.New(os.Stdin, os.Stdout)

	if user, err := user.Current(); err == nil {
		fmt.Printf("Hello %s! This is the Monkey programming language!\n",
			user.Username)
	}

	fmt.Printf("Feel free to type in commands\n")
	fmt.Printf("Type \".help\" for more information.\n")

	r.Start()
}
