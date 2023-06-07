package parser

import (
	"fmt"

	"github.com/adamvonbaron/gointerpreter/ast"
	"github.com/adamvonbaron/gointerpreter/lexer"
	"github.com/adamvonbaron/gointerpreter/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// read two tokens
	// so curToken and
	// peekToken are set
	p.nextToken()
	p.nextToken()

	return p

}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("next token got=%s want=%s", p.peekToken.Type, t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek will assert that the parser's peekToken
// TokenType is equal to t, and then advance to the next
// token and return true otherwise, do not advance to
// the next token and return false
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	// log an error noting that our
	// parser expected peekToken
	// to be t, but was sadly
	// p.peekToken.Type
	p.peekError(t)
	return false
}

// parse a let statement in monkey code
func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: p.curToken}

	// since we know the first token was a token.LET
	// our next token must be a token.IDENT
	// according to our loose grammar

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	// next token must be token.ASSIGN ("=")
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// skip over expression for now

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

// parse a return statement in monkey code
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{
		Token: p.curToken,
	}

	// next token should be expression
	p.nextToken()

	// but we are skipping for now
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for !p.curTokenIs(token.EOF) {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}

	return program
}
