package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/junkboy0315/go-interpreter/lexer"
	"github.com/junkboy0315/go-interpreter/token"
)

// PROMPT はプロンプトの左端に表示する記号
const PROMPT = ">> "

// Start はプロンプトを開始する
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.CreateNewLexer(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
