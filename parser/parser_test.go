package parser

import (
	"os"
	"testing"

	"github.com/adamvonbaron/gointerpreter/ast"
	"github.com/adamvonbaron/gointerpreter/lexer"
)

func testLetStatement(t *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "let" {
		t.Errorf("statement.TokenLiter() not 'let' got=%q", statement.TokenLiteral())
		return false
	}

	letStatement, ok := statement.(*ast.LetStatement)
	if !ok {
		t.Errorf("statement not *ast.Statment, got=%T", letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("letStatement.Name.TokenLiteral() got=%s want=%s", letStatement.Name.TokenLiteral(), name)
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser had %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf(msg)
	}
	t.FailNow()
}

func TestReturnStatements(t *testing.T) {
	fileInput, err := os.ReadFile("./parser_test_return_statements.monkey")
	if err != nil {
		panic("error reading parser_test_return_statements.monkey")
	}

	input := string(fileInput)

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements got=%d want=3", len(program.Statements))
	}

	for _, statement := range program.Statements {
		returnStatement, ok := statement.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("statement not *ast.ReturnStatement, got=%T", returnStatement)
		}

		if returnStatement.TokenLiteral() != "return" {
			t.Errorf("statement.TokenLiteral() got=%s want=return", returnStatement.TokenLiteral())
		}
	}
}

func TestLetStatements(t *testing.T) {
	fileInput, err := os.ReadFile("./parser_test_let_statements.monkey")
	if err != nil {
		panic("error reading parser_test_let_statements.monkey")
	}

	input := string(fileInput)

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	// before confirming that our
	// ast is properly created,
	// make sure that the lexer has
	// given us tokens in the
	// proper order
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements got=%d want=3", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, testCase := range tests {
		statement := program.Statements[i]

		if !testLetStatement(t, statement, testCase.expectedIdentifier) {
			return
		}
	}
}
