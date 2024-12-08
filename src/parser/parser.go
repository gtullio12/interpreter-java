package parser

import (
	"fmt"
	"java/ast"
	"java/lexer"
	"java/tokens"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  tokens.Token
	peekToken tokens.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t tokens.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != tokens.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case tokens.INTEGER_DT:
		return p.parseIntStatement()
	case tokens.STRING_DT:
		return p.parseStringStatement()
	default:
		return nil
	}
}

func (p *Parser) parseStringStatement() *ast.StringAssignmentStatement {
	stmt := &ast.StringAssignmentStatement{Token: p.curToken}

	if !p.expectPeek(tokens.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(tokens.ASSIGN) {
		return nil
	}

	// Skipping parsing expression
	for !p.curTokenIs(tokens.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseIntStatement() *ast.IntegerAssignmentStatement {
	stmt := &ast.IntegerAssignmentStatement{Token: p.curToken}

	if !p.expectPeek(tokens.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(tokens.ASSIGN) {
		return nil
	}

	// Skipping parsing expression
	for !p.curTokenIs(tokens.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t tokens.TokenType) bool {
	return p.curToken.Type == t
}
func (p *Parser) peekTokenIs(t tokens.TokenType) bool {
	return p.peekToken.Type == t
}
func (p *Parser) expectPeek(t tokens.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
