package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/schnell3526/writing_an_interpreter_in_go/src/lexer"
	"github.com/schnell3526/writing_an_interpreter_in_go/src/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" {
			fmt.Println("See you later!")
			return
		}
		l := lexer.NewLexer(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
