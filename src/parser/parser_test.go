package parser

import (
	"fmt"
	"java/ast"
	"java/lexer"
	"testing"
)

func TestIntegerAssignmentStatement(t *testing.T) {
	input := `
		int x = 5;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	fmt.Print(program)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statement. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testIntegerAssignmentStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testIntegerAssignmentStatement(t *testing.T, s ast.Statement, name string) bool {

	if s.TokenLiteral() != "int" {
		t.Errorf("s.TokenLiteral not 'int'. got=%q", s.TokenLiteral())
		return false
	}

	stmt, ok := s.(*ast.IntegerAssignmentStatement)
	if !ok {
		t.Errorf("s not *ast.IntegerAssignmentStatement. got=%T", s)
		return false
	}
	if stmt.Name.Value != name {
		t.Errorf("stmt.Name.Value not '%s'. got=%s", name, stmt.Name.Value)
		return false
	}
	if stmt.Name.TokenLiteral() != name {
		t.Errorf("stmt.Name.TokenLiteral() not '%s'. got=%s",
			name, stmt.Name.TokenLiteral())
		return false
	}
	return true

}
