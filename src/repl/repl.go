package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kriptonian1/BroLang/src/lexer"
	"github.com/kriptonian1/BroLang/src/token"
)

const PROMPT = "✨ >> "

/*
Start starts the REPL (Read-Eval-Print-Loop)

@param in io.Reader - The input to read from (usually os.Stdin) (type: io.Reader) (required)

@param out io.Writer - The output to write to (usually os.Stdout) (type: io.Writer) (required)
*/
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		if line[0] == '.' {
			cmdHelper(line)
			continue
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() { // Iterate through the tokens until EOF is reached
			fmt.Printf("%+v\n", tok)
		}
	}

}
