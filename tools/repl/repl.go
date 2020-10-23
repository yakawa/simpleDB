package repl

import (
	"bufio"
	"bytes"
	"fmt"
	"io"

	"github.com/yakawa/simpleDB/common/result"
	"github.com/yakawa/simpleDB/compiler/lexer"
	"github.com/yakawa/simpleDB/compiler/parser"
	"github.com/yakawa/simpleDB/compiler/translator"
	"github.com/yakawa/simpleDB/vm"
)

const PROMPT = ">>"

func Start(in io.ReadCloser, out io.Writer) {
	r := bufio.NewReader(in)
	for i := 0; ; i++ {
		fmt.Fprintf(out, PROMPT)

		var bb *bytes.Buffer
		for {
			buf, isPrefix, err := r.ReadLine()
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}

			if i == 0 {
				bb = bytes.NewBuffer(buf)
			} else {
				if _, err = bb.Write(buf); err != nil {
					panic(err)
				}
			}
			if !isPrefix {
				break
			}
		}

		tokens := lexer.Lex(bb.String())
		a := parser.Parse(tokens)
		vc := translator.Translate(a)
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
		fmt.Fprintf(out, "\n")
	}
}
