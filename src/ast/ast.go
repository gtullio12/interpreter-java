package ast

import (
	"bytes"
	"java/tokens"
)

type Node interface {
	TokenLiteral() string
	String() string
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
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

type Boolean struct {
	Token tokens.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

type IntegerAssignmentStatement struct {
	Token tokens.Token // the token.INT token
	Name  *Identifier
	Value Expression
}

type BooleanAssignmentStatement struct {
	Token tokens.Token // token.TRUE
	Name  *Identifier
	Value Expression
}

type StringAssignmentStatement struct {
	Token tokens.Token // the token.STRING token
	Name  *Identifier
	Value Expression
}

type ReturnStatement struct {
	Token       tokens.Token // the tokens.RETURN
	ReturnValue Expression
}

type PrefixExpression struct {
	Token    tokens.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

type InfixExpression struct {
	Token    tokens.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}

type IntegerLiteral struct {
	Token tokens.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

func (bs *BooleanAssignmentStatement) String() string {
	var out bytes.Buffer
	out.WriteString(bs.Name.TokenLiteral())
	out.WriteString(bs.Name.String())
	out.WriteString(" = ")

	if bs.Value != nil {
		out.WriteString(bs.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

func (is *IntegerAssignmentStatement) String() string {
	var out bytes.Buffer
	out.WriteString(is.TokenLiteral() + " ")
	out.WriteString(is.Name.String())
	out.WriteString(" = ")
	if is.Value != nil {
		out.WriteString(is.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

func (ss *StringAssignmentStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ss.TokenLiteral() + " ")
	out.WriteString(ss.Name.String())
	out.WriteString(" = ")
	if ss.Value != nil {
		out.WriteString(ss.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

func (i *Identifier) String() string { return i.Value }

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (bs *BooleanAssignmentStatement) statementNode()       {}
func (bs *BooleanAssignmentStatement) TokenLiteral() string { return bs.Token.Literal }

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
