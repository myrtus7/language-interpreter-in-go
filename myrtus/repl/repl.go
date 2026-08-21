package repl

import (
	"bufio"
	"fmt"
	"myrtus/lexer"
	"myrtus/token"
	"os"
)

const PROMPT = ">> "

func Start() {

	for {
		fmt.Print(PROMPT)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		input := scanner.Text()
		l := lexer.New(input)

		for tok := l.GenerateToken(); tok.Type != token.EOF; tok = l.GenerateToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
