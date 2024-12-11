package parser

import (
	"fmt"
	"java/ast"
	"java/lexer"
	"java/tokens"
	"testing"
)

func TestIntegerAssignmentStatement(t *testing.T) {
	input := `
		int x = 5;
		int y = 2;
		int num = 20;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statement. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"num"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testIntegerAssignmentStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestStringAssignmentStatement(t *testing.T) {
	input := `
		String x = "Hello world";
		String h = "hello";
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	fmt.Print(program)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 2 statement. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"h"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testStringAssignmentStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
		return 5;
		return 10;
		return 993322;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
				returnStmt.TokenLiteral())
		}
	}
}

func testStringAssignmentStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "String" {
		t.Errorf("s.TokenLiteral not 'String'. got=%q", s.TokenLiteral())
		return false
	}

	stmt, ok := s.(*ast.StringAssignmentStatement)
	if !ok {
		t.Errorf("s not *ast.StringAssignmentStatement. got=%T", s)
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

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.IntegerAssignmentStatement{
				Token: tokens.Token{Type: tokens.INTEGER_DT, Literal: "int"},
				Name: &ast.Identifier{
					Token: tokens.Token{Type: tokens.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &ast.Identifier{
					Token: tokens.Token{Type: tokens.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	fmt.Println(program.String())

	if program.String() != "int myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := `int x = y;`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.IntegerAssignmentStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.IntegerAssignmentStatement. got=%T",
			program.Statements[0])
	}
	ident, ok := stmt.Value.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Value)
	}
	if ident.Value != "y" {
		t.Errorf("ident.Value not %s. got=%s", "y", ident.Value)
	}
	if ident.TokenLiteral() != "y" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "y",
			ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "int x = 5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.IntegerAssignmentStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.IntegerAssignmentStatement. got=%T",
			program.Statements[0])
	}
	literal, ok := stmt.Value.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Value)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5",
			literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	input := "int x = -5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.IntegerAssignmentStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.IntegerAssignmentStatement. got=%T",
			program.Statements[0])
	}
	exp, ok := stmt.Value.(*ast.PrefixExpression)
	if !ok {
		t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Value)
	}
	if exp.Operator != "-" {
		t.Fatalf("exp.Operator is not '%s'. got=%s",
			"-", exp.Operator)
	}
	if !testIntegerLiteral(t, exp.Right, 5) {
		return
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value,
			integ.TokenLiteral())
		return false
	}

	return true
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"int x = 5 + 5;", 5, "+", 5},
		{"int x = 5 - 5;", 5, "-", 5},
		{"int x = 5 * 5;", 5, "*", 5},
		{"int x = 5 / 5;", 5, "/", 5},
		{"int x = 5 > 5;", 5, ">", 5},
		{"int x = 5 < 5;", 5, "<", 5},
		// TODO
		//{"boolean x = 5 == 5;", 5, "==", 5},
		//{"boolean x = 5 != 5;", 5, "!=", 5},}
	}
	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.IntegerAssignmentStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.IntegerAssignmentStatement. got=%T",
				program.Statements[0])
		}

		exp, ok := stmt.Value.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("exp is not ast.InfixExpression. got=%T", stmt.Value)
		}
		if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
			return
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
			return
		}
	}
}
