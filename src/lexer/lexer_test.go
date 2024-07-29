package lexer

import (
	"java/interpreter/tokens"
	"testing"
)

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

func TestLexer(t *testing.T) {
	lexer := New(javaCode)
	expectedResult := []tokens.Token{
		tokens.Token{Type: tokens.CLASS, Literal: "class"},
		tokens.Token{Type: tokens.IDENT, Literal: "Main"},
		tokens.Token{Type: tokens.LBRACE, Literal: "{"},
		tokens.Token{Type: tokens.PUBLIC, Literal: "public"},
		tokens.Token{Type: tokens.STATIC, Literal: "static"},
		tokens.Token{Type: tokens.VOID, Literal: "void"},
		tokens.Token{Type: tokens.IDENT, Literal: "main"},
		tokens.Token{Type: tokens.STRING, Literal: "String"},
		tokens.Token{Type: tokens.LSPAREN, Literal: "["},
		tokens.Token{Type: tokens.RSPAREN, Literal: "]"},
		tokens.Token{Type: tokens.RBRACE, Literal: "}"},
		tokens.Token{Type: tokens.INT, Literal: "int"},
		tokens.Token{Type: tokens.IDENT, Literal: "x"},
		tokens.Token{Type: tokens.ASSIGN, Literal: "="},
		tokens.Token{Type: tokens.INT, Literal: "5"},
		tokens.Token{Type: tokens.SEMICOLON, Literal: ";"},
		tokens.Token{Type: tokens.INT, Literal: "int"},
		tokens.Token{Type: tokens.IDENT, Literal: "y"},
		tokens.Token{Type: tokens.ASSIGN, Literal: "="},
		tokens.Token{Type: tokens.INT, Literal: "10"},
		tokens.Token{Type: tokens.SEMICOLON, Literal: ";"},
		tokens.Token{Type: tokens.STRING, Literal: "String"},
		tokens.Token{Type: tokens.IDENT, Literal: "test"},
		tokens.Token{Type: tokens.ASSIGN, Literal: "="},
		tokens.Token{Type: tokens.STRING, Literal: "Hello world"},
		tokens.Token{Type: tokens.SEMICOLON, Literal: ";"},
		tokens.Token{Type: tokens.PRINT, Literal: "System.out.println"},
		tokens.Token{Type: tokens.LPAREN, Literal: "("},
		tokens.Token{Type: tokens.IDENT, Literal: "test"},
		tokens.Token{Type: tokens.RPAREN, Literal: ")"},
		tokens.Token{Type: tokens.SEMICOLON, Literal: ";"},
		tokens.Token{Type: tokens.PRINT, Literal: "System.out.println"},
		tokens.Token{Type: tokens.LPAREN, Literal: "("},
		tokens.Token{Type: tokens.IDENT, Literal: "x"},
		tokens.Token{Type: tokens.PLUS, Literal: "+"},
		tokens.Token{Type: tokens.IDENT, Literal: "y"},
		tokens.Token{Type: tokens.RPAREN, Literal: ")"},
		tokens.Token{Type: tokens.SEMICOLON, Literal: ";"},
		tokens.Token{Type: tokens.RBRACE, Literal: "}"},
		tokens.Token{Type: tokens.PUBLIC, Literal: "public"},
		tokens.Token{Type: tokens.STATIC, Literal: "static"},
		tokens.Token{Type: tokens.INT, Literal: "int"},
		tokens.Token{Type: tokens.IDENT, Literal: "sum"},
		tokens.Token{Type: tokens.LPAREN, Literal: "("},
		tokens.Token{Type: tokens.INT, Literal: "int"},
		tokens.Token{Type: tokens.IDENT, Literal: "a"},
		tokens.Token{Type: tokens.COMMA, Literal: ","},
		tokens.Token{Type: tokens.INT, Literal: "int"},
		tokens.Token{Type: tokens.IDENT, Literal: "b"},
		tokens.Token{Type: tokens.RPAREN, Literal: ")"},
		tokens.Token{Type: tokens.LBRACE, Literal: "{"},
		tokens.Token{Type: tokens.RETURN, Literal: "return"},
		tokens.Token{Type: tokens.IDENT, Literal: "a"},
		tokens.Token{Type: tokens.PLUS, Literal: "+"},
		tokens.Token{Type: tokens.IDENT, Literal: "b"},
		tokens.Token{Type: tokens.SEMICOLON, Literal: ";"},
		tokens.Token{Type: tokens.LBRACE, Literal: "}"},
		tokens.Token{Type: tokens.LBRACE, Literal: "}"},
	}
}
