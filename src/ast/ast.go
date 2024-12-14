package ast

import (
	"bytes"
	"java/tokens"
	"strings"
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

type IfExpression struct {
	Token       tokens.Token // The 'if' token
	Condition   Expression
	Branches    []ElseIfExpression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

type ElseIfExpression struct {
	Token       tokens.Token // The else if token
	Condition   Expression
	Consequence *BlockStatement
}

func (eif *ElseIfExpression) expressionNode()      {}
func (eif *ElseIfExpression) TokenLiteral() string { return eif.Token.Literal }
func (eif *ElseIfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("else if")
	out.WriteString(eif.Condition.String())
	out.WriteString(" ")
	out.WriteString(eif.Consequence.String())
	return out.String()
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}

type ExpressionStatement struct {
	Token      tokens.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type BlockStatement struct {
	Token      tokens.Token // the { token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type Boolean struct {
	Token tokens.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

type CallExpression struct {
	Token     tokens.Token // The '(' token
	Function  Expression   // Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

type Parameter struct {
	DataType      tokens.Token
	ParameterName *Identifier
}

func (p *Parameter) expressionNode()      {}
func (p *Parameter) TokenLiteral() string { return p.DataType.Literal }
func (p *Parameter) String() string {
	var out bytes.Buffer
	out.WriteString(p.DataType.Literal + " ")
	out.WriteString(p.ParameterName.Value)
	return out.String()
}

type FunctionLiteral struct {
	Name       *Identifier
	Accessor   tokens.Token // e.g PUBLIC/PRIVATE
	ReturnType tokens.Token // e.g String, int,...
	Token      tokens.Token // The Accessor token
	Parameters []*Parameter
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}

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
	out.WriteString(bs.TokenLiteral() + " ")
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
