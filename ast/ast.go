package ast

import "github.com/adamvonbaron/gointerpreter/token"

type Node interface {
	// return literal value of token that
	// node is associated with
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// examples of let statements in monkey
// let x = 5;
// let y = 8383838;
// let hello = "hello world";
// let something = myFunc(x, y, hello);
// let <identifier> <assigns> <expression>
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier // the name of the variable
	Value Expression  // the expression to bind the name to
}

func (l *LetStatement) statementNode() {
}

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

// examples of return statements in monkey
// return 5;
// return false;
// return add(x, y);
// return <expression>
type ReturnStatement struct {
	Token       token.Token // the token.RETURN token
	ReturnValue Expression
}

func (r *ReturnStatement) statementNode() {}
func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}
