package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/nichtsam/monkey/lexer"
	"github.com/nichtsam/monkey/token"
)

type Repl struct {
	reader io.Reader
	writer io.Writer
}

const PROMPT = ">> "

func New(reader io.Reader, writer io.Writer) *Repl {
	return &Repl{reader, writer}
}

func (r *Repl) Start() {
	scanner := bufio.NewScanner(r.reader)

	for {
		fmt.Fprint(r.writer, PROMPT)
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		line = strings.Trim(line, " ")

		if line == ".help" {
			fmt.Fprintln(r.writer, ".exit     Exit the REPL")
			fmt.Fprintln(r.writer, ".help     Print this help message")
			continue
		}

		if line == ".exit" {
			break
		}

		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(r.writer, "%+v\n", tok)
		}
	}
}
