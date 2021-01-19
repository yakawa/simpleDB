package main

import (
	"os"

	"github.com/yakawa/simpleDB/frontend/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
