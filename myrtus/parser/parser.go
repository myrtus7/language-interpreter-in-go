package parser

import (
	"myrtus/ast"
	"myrtus/lexer"
	"myrtus/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// at first nextToken() curToken nil and peekToken at first generated token
	p.nextToken()
	p.nextToken()
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.GenerateToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
