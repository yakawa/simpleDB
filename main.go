package main

import (
	"os"

	"github.com/yakawa/simpleDB/tools/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
