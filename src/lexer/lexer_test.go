package lexer

import (
	"java/tokens"
	"testing"
)

func TestLexerOneLine(t *testing.T) {
	const smallJavaCode = `int x = 5;`

	lexer := New(smallJavaCode)
	expectedResult := []tokens.Token{
		tokens.Token{Type: tokens.INTEGER_DT, Literal: "int"},
		tokens.Token{Type: tokens.IDENT, Literal: "x"},
		tokens.Token{Type: tokens.ASSIGN, Literal: "="},
		tokens.Token{Type: tokens.INT, Literal: "5"},
		tokens.Token{Type: tokens.SEMICOLON, Literal: ";"},
	}

	for _, tok := range expectedResult {
		result := lexer.NextToken()
		if result.Type != tok.Type {
			t.Errorf("For token: %v. Invalid token type! Token type should've been %v, but was %v\n", tok.Literal, tok.Type, result.Type)
		}
		if tok.Literal != result.Literal {
			t.Errorf("Invalid token literal! Token literal should've been %v, but was %v\n", tok.Literal, result.Literal)
		}
	}
}

func TestLexerString(t *testing.T) {
	const javaStringCode = `String str = "Hello world";`
	lexer := New(javaStringCode)
	expectedResult := []tokens.Token{
		tokens.Token{Type: tokens.STRING_DT, Literal: "String"},
		tokens.Token{Type: tokens.IDENT, Literal: "str"},
		tokens.Token{Type: tokens.ASSIGN, Literal: "="},
		tokens.Token{Type: tokens.STRING, Literal: "Hello world"},
		tokens.Token{Type: tokens.SEMICOLON, Literal: ";"},
	}

	for _, tok := range expectedResult {
		result := lexer.NextToken()
		if result.Type != tok.Type {
			t.Errorf("For token: %v. Invalid token type! Token type should've been %v, but was %v\n", tok.Literal, tok.Type, result.Type)
		}
		if tok.Literal != result.Literal {
			t.Errorf("Invalid token literal! Token literal should've been %v, but was %v\n", tok.Literal, result.Literal)
		}
	}

}

func TestLexerLarge(t *testing.T) {
	const javaCode = `
	class Main {
		public static void main(String[] args) {
			int x = 5;
			int y = 10;
			String test = "Hello world";
			System.out.println(test);
			System.out.println(x + y);
		}

		public static int sum(int a, int b) {
			return a + b;
		}
	}
	`
	lexer := New(javaCode)
	expectedResult := []tokens.Token{
		tokens.Token{Type: tokens.CLASS, Literal: "class"},
		tokens.Token{Type: tokens.IDENT, Literal: "Main"},
		tokens.Token{Type: tokens.LBRACE, Literal: "{"},
		tokens.Token{Type: tokens.PUBLIC, Literal: "public"},
		tokens.Token{Type: tokens.STATIC, Literal: "static"},
		tokens.Token{Type: tokens.VOID, Literal: "void"},
		tokens.Token{Type: tokens.IDENT, Literal: "main"},
		tokens.Token{Type: tokens.LPAREN, Literal: "("},
		tokens.Token{Type: tokens.STRING_DT, Literal: "String"},
		tokens.Token{Type: tokens.LSPAREN, Literal: "["},
		tokens.Token{Type: tokens.RSPAREN, Literal: "]"},
		tokens.Token{Type: tokens.IDENT, Literal: "args"},
		tokens.Token{Type: tokens.RPAREN, Literal: ")"},
		tokens.Token{Type: tokens.LBRACE, Literal: "{"},
		tokens.Token{Type: tokens.INTEGER_DT, Literal: "int"},
		tokens.Token{Type: tokens.IDENT, Literal: "x"},
		tokens.Token{Type: tokens.ASSIGN, Literal: "="},
		tokens.Token{Type: tokens.INT, Literal: "5"},
		tokens.Token{Type: tokens.SEMICOLON, Literal: ";"},
		tokens.Token{Type: tokens.INTEGER_DT, Literal: "int"},
		tokens.Token{Type: tokens.IDENT, Literal: "y"},
		tokens.Token{Type: tokens.ASSIGN, Literal: "="},
		tokens.Token{Type: tokens.INT, Literal: "10"},
		tokens.Token{Type: tokens.SEMICOLON, Literal: ";"},
		tokens.Token{Type: tokens.STRING_DT, Literal: "String"},
		tokens.Token{Type: tokens.IDENT, Literal: "test"},
		tokens.Token{Type: tokens.ASSIGN, Literal: "="},
		tokens.Token{Type: tokens.STRING, Literal: "Hello world"},
		tokens.Token{Type: tokens.SEMICOLON, Literal: ";"},
		tokens.Token{Type: tokens.SYSTEM, Literal: "System"},
		tokens.Token{Type: tokens.PERIOD, Literal: "."},
		tokens.Token{Type: tokens.OUT, Literal: "out"},
		tokens.Token{Type: tokens.PERIOD, Literal: "."},
		tokens.Token{Type: tokens.PRINTLN, Literal: "println"},
		tokens.Token{Type: tokens.LPAREN, Literal: "("},
		tokens.Token{Type: tokens.IDENT, Literal: "test"},
		tokens.Token{Type: tokens.RPAREN, Literal: ")"},
		tokens.Token{Type: tokens.SEMICOLON, Literal: ";"},
		tokens.Token{Type: tokens.SYSTEM, Literal: "System"},
		tokens.Token{Type: tokens.PERIOD, Literal: "."},
		tokens.Token{Type: tokens.OUT, Literal: "out"},
		tokens.Token{Type: tokens.PERIOD, Literal: "."},
		tokens.Token{Type: tokens.PRINTLN, Literal: "println"},
		tokens.Token{Type: tokens.LPAREN, Literal: "("},
		tokens.Token{Type: tokens.IDENT, Literal: "x"},
		tokens.Token{Type: tokens.PLUS, Literal: "+"},
		tokens.Token{Type: tokens.IDENT, Literal: "y"},
		tokens.Token{Type: tokens.RPAREN, Literal: ")"},
		tokens.Token{Type: tokens.SEMICOLON, Literal: ";"},
		tokens.Token{Type: tokens.RBRACE, Literal: "}"},
		tokens.Token{Type: tokens.PUBLIC, Literal: "public"},
		tokens.Token{Type: tokens.STATIC, Literal: "static"},
		tokens.Token{Type: tokens.INTEGER_DT, Literal: "int"},
		tokens.Token{Type: tokens.IDENT, Literal: "sum"},
		tokens.Token{Type: tokens.LPAREN, Literal: "("},
		tokens.Token{Type: tokens.INTEGER_DT, Literal: "int"},
		tokens.Token{Type: tokens.IDENT, Literal: "a"},
		tokens.Token{Type: tokens.COMMA, Literal: ","},
		tokens.Token{Type: tokens.INTEGER_DT, Literal: "int"},
		tokens.Token{Type: tokens.IDENT, Literal: "b"},
		tokens.Token{Type: tokens.RPAREN, Literal: ")"},
		tokens.Token{Type: tokens.LBRACE, Literal: "{"},
		tokens.Token{Type: tokens.RETURN, Literal: "return"},
		tokens.Token{Type: tokens.IDENT, Literal: "a"},
		tokens.Token{Type: tokens.PLUS, Literal: "+"},
		tokens.Token{Type: tokens.IDENT, Literal: "b"},
		tokens.Token{Type: tokens.SEMICOLON, Literal: ";"},
		tokens.Token{Type: tokens.RBRACE, Literal: "}"},
		tokens.Token{Type: tokens.RBRACE, Literal: "}"},
	}
	for i, token := range expectedResult {
		nextToken := lexer.NextToken()
		if nextToken.Type != token.Type {
			t.Errorf("For token: %v. Invalid token type! Token type should've been %v, but was %v. for index: %v\n", token.Literal, token.Type, nextToken.Type, i)
		}
		if nextToken.Literal != nextToken.Literal {
			t.Errorf("Invalid token literal! Token literal should've been %v, but was %v. for index: %v\n", token.Literal, nextToken.Literal, i)
		}
	}
}
