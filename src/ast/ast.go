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

type ReturnStatement struct {
	Token       tokens.Token // the tokens.RETURN
	ReturnValue Expression
}

type IntegerLiteral struct {
	Token tokens.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

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
