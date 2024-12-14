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

func TestBooleanAssignmentStatement(t *testing.T) {
	input := "boolean a = false;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statement. got=%d", len(program.Statements))
	}

	testBooleanStatement(t, program.Statements[0], "a")
}

func TestReturnStatementExpression(t *testing.T) {
	input := `return x + y;`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d",
			len(program.Statements))
	}

	returnStmt, _ := program.Statements[0].(*ast.ReturnStatement)

	infixExpression, _ := returnStmt.ReturnValue.(*ast.InfixExpression)

	if !testInfixExpression(t, infixExpression, "x", "+", "y") {
		return
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

func testBooleanStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "boolean" {
		t.Errorf("s.TokenLiteral not 'boolean'. got=%q", s.TokenLiteral())
		return false
	}

	stmt, ok := s.(*ast.BooleanAssignmentStatement)
	if !ok {
		t.Errorf("s not *ast.BooleanAssignmentStatement. got=%T", s)
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
	testIdentifier(t, stmt.Value, "y")
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

		exp, _ := stmt.Value.(*ast.InfixExpression)
		testInfixExpression(t, exp, tt.leftValue, tt.operator, tt.rightValue)
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"add(b * c)",
			"add((b * c))",
		},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a + b) + ((c * d) / f)) + g))",
		},
		{
			"boolean a = b > c;",
			"boolean a = (b > c);",
		},
		{
			"int a = 1 + (2 + 3) + 4;",
			"int a = ((1 + (2 + 3)) + 4);",
		},
		{
			"int a = (5 + 5) * 2;",
			"int a = ((5 + 5) * 2);",
		},
		{
			"int a = 2 / (5 + 5);",
			"int a = (2 / (5 + 5));",
		},
		{
			"int a = -(5 + 5);",
			"int a = (-(5 + 5));",
		},
		{
			"int c = -a * b;",
			"int c = ((-a) * b);",
		},
		{
			"int d = a + b + c;",
			"int d = ((a + b) + c);",
		},
		{
			"int d = a + b - c;",
			"int d = ((a + b) - c);",
		},
		{
			"int d = a * b * c;",
			"int d = ((a * b) * c);",
		},
		{
			"int d = a * b / c;",
			"int d = ((a * b) / c);",
		},
		{
			"int d = a + b / c;",
			"int d = (a + (b / c));",
		},
		{
			"int g = a + b * c + d / e - f;",
			"int g = (((a + (b * c)) + (d / e)) - f);",
		},
		{
			"boolean a = 5 > 4 == 3 < 4;",
			"boolean a = ((5 > 4) == (3 < 4));",
		},
		{
			"boolean a = 5 < 4 != 3 > 4;",
			"boolean a = ((5 < 4) != (3 > 4));",
		},
		{
			"boolean a = 3 + 4 * 5 == 3 * 1 + 4 * 5;",
			"boolean a = ((3 + (4 * 5)) == ((3 * 1) + (4 * 5)));",
		},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)

	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}

	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}

	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value,
			ident.TokenLiteral())
		return false
	}
	return true
}

func testLiteralExpression(
	t *testing.T,
	exp ast.Expression,
	expected interface{},
) bool {
	switch v := expected.(type) {
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testIdentifier(t, exp, v)
	}

	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{},
	operator string, right interface{}) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.InfixExpression. got=%T(%s)", exp, exp)
		return false
	}
	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}
	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s'. got=%q", operator, opExp.Operator)
		return false
	}
	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}
	return true
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	bo, ok := exp.(*ast.Boolean)
	if !ok {
		t.Errorf("exp not *ast.Boolean. got=%T", exp)
		return false
	}

	if bo.Value != value {
		t.Errorf("bo.Value not %t. got=%t", value, bo.Value)
		return false
	}

	if bo.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("bo.TokenLiteral not %t. got=%s",
			value, bo.TokenLiteral())
		return false
	}

	return true
}

func TestIfExpression(t *testing.T) {
	input := `if (x > y){ return x + y; }`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 1, len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.IfExpression. got=%T", stmt.Expression)
	}

	if !testInfixExpression(t, exp.Condition, "x", ">", "y") {
		return
	}

	if len(exp.Consequence.Statements) != 1 {
		t.Errorf("consequence is not 1 statements. got=%d\n", len(exp.Consequence.Statements))
	}

	consequence, ok := exp.Consequence.Statements[0].(*ast.ReturnStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ReturnStatement. got=%T", exp.Consequence.Statements[0])
	}

	infixExpression, _ := consequence.ReturnValue.(*ast.InfixExpression)

	if !testInfixExpression(t, infixExpression, "x", "+", "y") {
		return
	}

	if exp.Alternative != nil {
		t.Errorf("exp.Alternative.Statements was not nil. got=%+v", exp.Alternative)
	}
}

func TestIfElseExpression(t *testing.T) {
	input := `if (x > y) {
	return x + y;
	} else {
	 return x - y;
	}`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 1, len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.IfExpression. got=%T", stmt.Expression)
	}

	if !testInfixExpression(t, exp.Condition, "x", ">", "y") {
		return
	}

	if len(exp.Consequence.Statements) != 1 {
		t.Errorf("consequence is not 1 statements. got=%d\n", len(exp.Consequence.Statements))
	}

	consequence, ok := exp.Consequence.Statements[0].(*ast.ReturnStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ReturnStatement. got=%T", exp.Consequence.Statements[0])
	}

	infixExpression, _ := consequence.ReturnValue.(*ast.InfixExpression)

	if !testInfixExpression(t, infixExpression, "x", "+", "y") {
		return
	}

	alt := exp.Alternative

	if len(alt.Statements) != 1 {
		t.Fatalf("Else statements should only have 1 statement, got=%d\n", len(alt.Statements))
	}
	altInfixExpression, _ := alt.Statements[0].(*ast.ReturnStatement)

	if !testInfixExpression(t, altInfixExpression.ReturnValue, "x", "-", "y") {
		return
	}
}

func TestIfElseIfStatement(t *testing.T) {
	input := `if (x > y) {
				return x + y;
			  } else if (x < y) {
				return x - y;
			  } else {
	 			return x / y;
			  }`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 1, len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.IfExpression. got=%T", stmt.Expression)
	}

	if !testInfixExpression(t, exp.Condition, "x", ">", "y") {
		return
	}

	if len(exp.Consequence.Statements) != 1 {
		t.Errorf("consequence is not 1 statements. got=%d\n", len(exp.Consequence.Statements))
	}

	consequence, ok := exp.Consequence.Statements[0].(*ast.ReturnStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ReturnStatement. got=%T", exp.Consequence.Statements[0])
	}

	infixExpression, _ := consequence.ReturnValue.(*ast.InfixExpression)

	if !testInfixExpression(t, infixExpression, "x", "+", "y") {
		return
	}

	alt := exp.Alternative

	if len(alt.Statements) != 1 {
		t.Fatalf("Else statements should only have 1 statement, got=%d\n", len(alt.Statements))
	}
	altInfixExpression, _ := alt.Statements[0].(*ast.ReturnStatement)

	if !testInfixExpression(t, altInfixExpression.ReturnValue, "x", "/", "y") {
		return
	}
}

func TestIfElseStatement2(t *testing.T) {
	input := `if (x > y) {
				return x + y;
			  } else if (x < y) {
				return x - y;
			  } else if (false) {
			   return x * y;
			  } else {
	 			return x / y;
			  }`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 1, len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.IfExpression)

	elseIfs := exp.Branches

	if len(elseIfs) != 2 {
		t.Fatalf("Should have 2 Else If Expressions but got=%d\n", len(elseIfs))
	}

	if !ok {
		t.Fatalf("stmt.Expression is not ast.IfExpression. got=%T", stmt.Expression)
	}

	if !testInfixExpression(t, exp.Condition, "x", ">", "y") {
		return
	}

	if len(exp.Consequence.Statements) != 1 {
		t.Errorf("consequence is not 1 statements. got=%d\n", len(exp.Consequence.Statements))
	}

	consequence, ok := exp.Consequence.Statements[0].(*ast.ReturnStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ReturnStatement. got=%T", exp.Consequence.Statements[0])
	}

	infixExpression, _ := consequence.ReturnValue.(*ast.InfixExpression)

	if !testInfixExpression(t, infixExpression, "x", "+", "y") {
		return
	}

	alt := exp.Alternative

	if len(alt.Statements) != 1 {
		t.Fatalf("Else statements should only have 1 statement, got=%d\n", len(alt.Statements))
	}
	altInfixExpression, _ := alt.Statements[0].(*ast.ReturnStatement)

	if !testInfixExpression(t, altInfixExpression.ReturnValue, "x", "/", "y") {
		return
	}

}

func TestFunctionParsing(t *testing.T) {
	input := `public String append(String x, String y) { return x + y; }`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	function, ok := stmt.Expression.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.FunctionLiteral. got=%T",
			stmt.Expression)
	}

	if len(function.Parameters) != 2 {
		t.Fatalf("function literal parameters wrong. want 2, got=%d\n", len(function.Parameters))
	}

	testParameterExpression(t, *function.Parameters[0], "String", "x")
	testParameterExpression(t, *function.Parameters[1], "String", "y")

	if len(function.Body.Statements) != 1 {
		t.Fatalf("function.Body.Statements has not 1 statements. got=%d\n", len(function.Body.Statements))
	}

	bodyStmt, ok := function.Body.Statements[0].(*ast.ReturnStatement)
	if !ok {
		t.Fatalf("function body stmt is not ast.ReturnStatement. got=%T",
			function.Body.Statements[0])
	}

	testInfixExpression(t, bodyStmt.ReturnValue, "x", "+", "y")
}

func testParameterExpression(t *testing.T, p ast.Parameter, datatype interface{}, name interface{}) bool {
	if p.ParameterName.Value != name {
		t.Fatalf("Parameter name should be %s. but was %s\n", p.ParameterName.Value, name)
		return false
	}

	if p.DataType.Literal != datatype {
		t.Fatalf("Data type for parameter should be %s. but was %s\n", p.DataType.Literal, datatype)
		return false
	}
	return true
}

func TestCallExpressionParsing(t *testing.T) {
	input := "add(1, 2 * 3, 4 + 5);"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stmt is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.CallExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.CallExpression. got=%T", stmt.Expression)
	}

	if !testIdentifier(t, exp.Function, "add") {
		return
	}

	if len(exp.Arguments) != 3 {
		t.Fatalf("wrong length of arguments. got=%d", len(exp.Arguments))
	}

	testLiteralExpression(t, exp.Arguments[0], int64(1))
	testInfixExpression(t, exp.Arguments[1], int64(2), "*", int64(3))
	testInfixExpression(t, exp.Arguments[2], int64(4), "+", int64(5))
}
