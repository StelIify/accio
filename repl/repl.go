package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/StelIify/accio/lexer"
	"github.com/StelIify/accio/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.NewLexer(line)
		for tok := l.NextToken(); tok.Type != token.EOF; {
			fmt.Fprintf(out, "%+v\n", tok)
			tok = l.NextToken()
		}
	}
}
