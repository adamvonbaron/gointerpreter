package lexer

import "github.com/adamvonbaron/gointerpreter/token"

type Lexer struct {
	input        string
	position     int  // current position in input
	readPosition int  // current reading position in input (after position)
	ch           byte // current character being checked
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	// initialize position, readPosition,
	// and ch before returning
	// to caller
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekNextChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// skip over all whitespace between
	// l.position and the next character
	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekNextChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{
				Type:    token.EQ,
				Literal: literal,
			}
		} else {
			tok = token.Token{
				Type:    token.ASSIGN,
				Literal: string(l.ch),
			}
		}
	case ';':
		tok = token.Token{
			Type:    token.SEMICOLON,
			Literal: string(l.ch),
		}
	case '(':
		tok = token.Token{
			Type:    token.LPAREN,
			Literal: string(l.ch),
		}
	case ')':
		tok = token.Token{
			Type:    token.RPAREN,
			Literal: string(l.ch),
		}
	case ',':
		tok = token.Token{
			Type:    token.COMMA,
			Literal: string(l.ch),
		}
	case '+':
		tok = token.Token{
			Type:    token.PLUS,
			Literal: string(l.ch),
		}
	case '-':
		tok = token.Token{
			Type:    token.MINUS,
			Literal: string(l.ch),
		}
	case '!':
		if l.peekNextChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{
				Type:    token.NOT_EQ,
				Literal: literal,
			}
		} else {
			tok = token.Token{
				Type:    token.BANG,
				Literal: string(l.ch),
			}
		}
	case '*':
		tok = token.Token{
			Type:    token.ASTERISK,
			Literal: string(l.ch),
		}
	case '/':
		tok = token.Token{
			Type:    token.SLASH,
			Literal: string(l.ch),
		}
	case '<':
		tok = token.Token{
			Type:    token.LT,
			Literal: string(l.ch),
		}
	case '>':
		tok = token.Token{
			Type:    token.GT,
			Literal: string(l.ch),
		}
	case '{':
		tok = token.Token{
			Type:    token.LBRACE,
			Literal: string(l.ch),
		}
	case '}':
		tok = token.Token{
			Type:    token.RBRACE,
			Literal: string(l.ch),
		}
	case 0:
		tok = token.Token{
			Type:    token.EOF,
			Literal: "",
		}
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = token.Token{
				Type:    token.ILLEGAL,
				Literal: string(l.ch),
			}
		}
	}
	l.readChar()
	return tok
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'

}
