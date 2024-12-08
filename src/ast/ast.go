package ast

import (
	"java/tokens"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type IntegerAssignmentStatement struct {
	Token tokens.Token // the token.INT token
	Name  *Identifier
	Value Expression
}

type StringAssignmentStatement struct {
	Token tokens.Token // the token.STRING token
	Name  *Identifier
	Value Expression
}

func (ls *IntegerAssignmentStatement) statementNode()       {}
func (ls *IntegerAssignmentStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *StringAssignmentStatement) statementNode()       {}
func (ls *StringAssignmentStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token tokens.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
