package main

import (
	"fmt"
	"os"

	"github.com/StelIify/accio/repl"
)

func main() {
	fmt.Println("Hello, this is Accio programming language!")
	fmt.Println("Feel free to type in commands: ")
	repl.Start(os.Stdin, os.Stdout)
}
