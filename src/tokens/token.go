package tokens

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT" // add, foobar, x, y, ...
	INT    = "INT"
	STRING = "STRING"

	// Operators
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	INCREMENT = "++"
	DECREMENT = "--"
	SLASH     = "/"
	PERIOD    = "."

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN    = "("
	RPAREN    = ")"
	LSPAREN   = "["
	RSPAREN   = "]"
	LBRACE    = "{"
	RBRACE    = "}"
	QUOTATION = "\""

	// Keywords
	TRUE    = "TRUE"
	FALSE   = "FALSE"
	IF      = "IF"
	ELSE    = "ELSE"
	ELSE_IF = "ELSE IF"
	RETURN  = "RETURN"
	CLASS   = "CLASS"
	STATIC  = "STATIC"

	// Access modifiers
	PUBLIC  = "PUBLIC"
	PRIVATE = "PRIVATE"

	// return type
	VOID = "VOID"

	// Data types
	INTEGER_DT   = "int"
	STRING_DT    = "String"
	CHARACTER_DT = "char"
	BOOLEAN_DT   = "boolean"

	// External System library
	SYSTEM  = "System"
	OUT     = "OUT"
	PRINTLN = "PRINTLN"
)

var keywords = map[string]TokenType{
	"class":   CLASS,
	"true":    TRUE,
	"false":   FALSE,
	"else":    ELSE,
	"public":  PUBLIC,
	"private": PRIVATE,
	"static":  STATIC,
	"void":    VOID,
	"System":  SYSTEM,
	"out":     OUT,
	"println": PRINTLN,
	"int":     INTEGER_DT,
	"String":  STRING_DT,
	"return":  RETURN,
	"boolean": BOOLEAN_DT,
	"if":      IF,
	"else if": ELSE_IF,
}

func LookupIdentifier(s string) TokenType {
	if tok, ok := keywords[s]; ok {
		return tok
	}
	return IDENT
}

type Token struct {
	Type    TokenType
	Literal string
}
