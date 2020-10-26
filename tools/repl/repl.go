package repl

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/yakawa/simpleDB/common/result"
	"github.com/yakawa/simpleDB/compiler/lexer"
	"github.com/yakawa/simpleDB/compiler/parser"
	"github.com/yakawa/simpleDB/compiler/planner"
	"github.com/yakawa/simpleDB/vm"
)

const PROMPT = ">>"

func Start(in io.ReadCloser, out io.Writer) {
	r := bufio.NewReader(in)
	for {
		fmt.Fprintf(out, PROMPT)

		var bb *bytes.Buffer
		bb = bytes.NewBuffer([]byte(""))
		for i := 0; ; i++ {
			buf, cont, err := r.ReadLine()
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}

			if i == 0 {
				bb = bytes.NewBuffer([]byte(""))
			}
			_, err = bb.Write(buf)
			if err != nil {
				panic(err)
			}
			// bb.Write(buf)
			if cont {
				continue
			}
			if len(buf) == 0 {
				break
			} else {
				bb.Write([]byte("\n"))
			}
		}

		line := bb.String()

		if strings.HasPrefix(line, ".") {
			if parseCommand(line, out) {
				return
			}
		} else {
			tokens := lexer.Lex(line)
			a, _ := parser.Parse(tokens)
			vc := planner.Translate(a)
			for _, c := range vc {
				fmt.Fprintf(out, "%#+v\n", c)
			}
			rs := vm.Run(vc)
			for i, col := range rs {
				switch col.Type {
				case result.Integral:
					fmt.Fprintf(out, "%d", col.Integral)
				}
				if i != (len(rs) - 1) {
					fmt.Fprintf(out, ",")
				}
			}
		}
		fmt.Fprintf(out, "\n")
	}
}

func parseCommand(line string, out io.Writer) bool {
	switch line {
	case ".exit":
		return true
	default:
		fmt.Fprintf(out, "Unknown Command")
		return false
	}
}
