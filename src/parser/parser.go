package parser

import (
	"fmt"
	"java/ast"
	"java/lexer"
	"java/tokens"
	"strconv"
)

const (
	int = iota
	_
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

type Parser struct {
	l         *lexer.Lexer
	curToken  tokens.Token
	peekToken tokens.Token
	errors    []string

	prefixParseFns map[tokens.TokenType]prefixParseFn
	infixParseFns  map[tokens.TokenType]infixParseFn
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

var precedences = map[tokens.TokenType]int64{
	tokens.EQ:       EQUALS,
	tokens.NOT_EQ:   EQUALS,
	tokens.LT:       LESSGREATER,
	tokens.GT:       LESSGREATER,
	tokens.PLUS:     SUM,
	tokens.MINUS:    SUM,
	tokens.SLASH:    PRODUCT,
	tokens.ASTERISK: PRODUCT,
	tokens.LPAREN:   CALL,
}

// [...]
func (p *Parser) peekPrecedence() int64 {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}
func (p *Parser) curPrecedence() int64 {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) parseExpression(precedence int64) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]

	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	leftExp := prefix()

	for !p.peekTokenIs(tokens.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}
	return leftExp
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

func (p *Parser) registerPrefix(tokenType tokens.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
func (p *Parser) registerInfix(tokenType tokens.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()

	p.prefixParseFns = make(map[tokens.TokenType]prefixParseFn)
	p.registerPrefix(tokens.IDENT, p.parseIdentifier)
	p.registerPrefix(tokens.STRING, p.parseStringLiteral)
	p.registerPrefix(tokens.INT, p.parseIntegerLiteral)
	p.registerPrefix(tokens.MINUS, p.parsePrefixExpression)
	p.registerPrefix(tokens.TRUE, p.parseBoolean)
	p.registerPrefix(tokens.FALSE, p.parseBoolean)
	p.registerPrefix(tokens.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(tokens.IF, p.parseIfExpression)
	p.registerPrefix(tokens.PUBLIC, p.parseFunctionLiteral)
	p.registerPrefix(tokens.PRIVATE, p.parseFunctionLiteral)
	p.registerPrefix(tokens.VOID, p.parseFunctionLiteral)
	p.registerPrefix(tokens.BANG, p.parsePrefixExpression)
	p.infixParseFns = make(map[tokens.TokenType]infixParseFn)

	p.registerInfix(tokens.LPAREN, p.parseCallExpression)
	p.registerInfix(tokens.PLUS, p.parseInfixExpression)
	p.registerInfix(tokens.MINUS, p.parseInfixExpression)
	p.registerInfix(tokens.SLASH, p.parseInfixExpression)
	p.registerInfix(tokens.ASTERISK, p.parseInfixExpression)
	p.registerInfix(tokens.EQ, p.parseInfixExpression)
	p.registerInfix(tokens.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(tokens.LT, p.parseInfixExpression)
	p.registerInfix(tokens.GT, p.parseInfixExpression)

	return p
}

func (p *Parser) parseStringLiteral() ast.Expression {
	str := &ast.StringLiteral{Token: p.curToken}

	str.Value = p.curToken.Literal

	return str
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseCallArguments()
	return exp
}

func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if p.peekTokenIs(tokens.RPAREN) {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpression(LOWEST))

	for p.peekTokenIs(tokens.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(tokens.RPAREN) {
		return nil
	}
	return args
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	// public void getString()
	// private void getString()
	// void getString()
	// public String getString()
	// public int getString()

	lit := &ast.FunctionLiteral{Token: p.curToken}

	if p.curToken.Type == tokens.PUBLIC || p.curToken.Type == tokens.PRIVATE {
		lit.Accessor = p.curToken
		p.nextToken()
	}

	if p.curToken.Type == tokens.VOID || p.curToken.Type == tokens.STRING_DT || p.curToken.Type == tokens.INTEGER_DT || p.curToken.Type == tokens.CHARACTER_DT {
		lit.ReturnType = p.curToken
		p.nextToken()
	}

	lit.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(tokens.LPAREN) {
		return nil
	}

	p.nextToken()

	parameters := []*ast.Parameter{}

	for !p.curTokenIs(tokens.RPAREN) {
		if p.curTokenIs(tokens.COMMA) {
			p.nextToken()
		}
		param := &ast.Parameter{}
		param.DataType = p.curToken
		p.nextToken()
		param.ParameterName = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		parameters = append(parameters, param)
		p.nextToken()
	}

	if !p.expectPeek(tokens.LBRACE) {
		return nil
	}

	lit.Parameters = parameters

	lit.Body = p.parseBlockStatement()
	return lit
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()
	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(tokens.RPAREN) {
		return nil
	}

	return exp
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(tokens.TRUE)}
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

func (p *Parser) noPrefixParseFnError(t tokens.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
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

func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.curToken}
	if !p.expectPeek(tokens.LPAREN) {
		return nil
	}

	p.nextToken()

	expression.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(tokens.RPAREN) {
		return nil
	}

	if !p.expectPeek(tokens.LBRACE) {
		return nil
	}

	expression.Consequence = p.parseBlockStatement()

	elseIfExpressions := []ast.ElseIfExpression{}

	for p.peekTokenIs(tokens.ELSE_IF) {
		elseIfExpression := &ast.ElseIfExpression{Token: p.curToken}

		p.nextToken()

		if !p.expectPeek(tokens.LPAREN) {
			return nil
		}

		elseIfExpression.Condition = p.parseExpression(LOWEST)

		if !p.expectPeek(tokens.LBRACE) {
			return nil
		}

		elseIfExpression.Consequence = p.parseBlockStatement()

		elseIfExpressions = append(elseIfExpressions, *elseIfExpression)

	}

	expression.Branches = elseIfExpressions

	if p.peekTokenIs(tokens.ELSE) {
		p.nextToken()
		if !p.expectPeek(tokens.LBRACE) {
			return nil
		}
		p.nextToken()
		expression.Alternative = p.parseBlockStatement()
	}

	return expression
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}

	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(tokens.RBRACE) && !p.curTokenIs(tokens.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(tokens.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case tokens.IDENT:
		return p.parseIdentifierStatement()
	case tokens.INCREMENT:
		return p.parseIncrementStatement()
	case tokens.DECREMENT:
		return p.parseDecrementStatement()
	case tokens.BOOLEAN_DT:
		return p.parseBooleanStatement()
	case tokens.INTEGER_DT:
		return p.parseIntStatement()
	case tokens.STRING_DT:
		return p.parseStringStatement()
	case tokens.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseIdentifierStatement() ast.Statement {
	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if p.peekTokenIs(tokens.INCREMENT) {
		incrementStmt := &ast.IncrementStatement{Token: p.peekToken, Operand: ident, Side: "POSTFIX"}
		return incrementStmt
	} else if p.peekTokenIs(tokens.DECREMENT) {
		decrementStmt := &ast.DecrementStatement{Token: p.peekToken, Operand: ident, Side: "POSTFIX"}
		return decrementStmt
	}

	p.nextToken()

	return nil
}

func (p *Parser) parseIncrementStatement() *ast.IncrementStatement {
	if !p.peekTokenIs(tokens.IDENT) {
		return nil
	}
	p.nextToken()

	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	incrementStmt := &ast.IncrementStatement{Token: p.curToken, Operand: ident, Side: "PREFIX"}

	return incrementStmt
}

func (p *Parser) parseDecrementStatement() *ast.DecrementStatement {
	if !p.peekTokenIs(tokens.IDENT) {
		return nil
	}
	p.nextToken()

	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	decrementStmt := &ast.DecrementStatement{Token: p.curToken, Operand: ident, Side: "PREFIX"}

	return decrementStmt
}

func (p *Parser) parseBooleanStatement() *ast.BooleanAssignmentStatement {
	stmt := &ast.BooleanAssignmentStatement{Token: p.curToken}

	if !p.expectPeek(tokens.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(tokens.ASSIGN) {
		return nil
	}

	p.nextToken()

	for !p.curTokenIs(tokens.SEMICOLON) {
		exp := p.parseExpression(LOWEST)
		stmt.Value = exp
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	for !p.curTokenIs(tokens.SEMICOLON) {
		exp := p.parseExpression(LOWEST)
		stmt.ReturnValue = exp
		p.nextToken()
	}

	return stmt
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

	p.nextToken()

	for !p.curTokenIs(tokens.SEMICOLON) {
		exp := p.parseExpression(LOWEST)
		stmt.Value = exp
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

	p.nextToken()

	for !p.curTokenIs(tokens.SEMICOLON) {
		exp := p.parseExpression(LOWEST)
		stmt.Value = exp
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
