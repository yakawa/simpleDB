package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/yakawa/simpleDB/common/ast"
	"github.com/yakawa/simpleDB/common/token"
)

func (p *parser) parseExpression(precedence int) (*ast.Expression, error) {
	expr := &ast.Expression{}
	unary, exists := p.unaryParseFunc[p.currentToken.Type]
	if !exists {
		return expr, errors.New(fmt.Sprintf("Unknown Unary Operator: %s", p.currentToken.Type))
	}

	left, err := unary()
	if err != nil {
		return expr, err
	}

	for p.getNextToken().Type != token.EOS && precedence < p.getNextTokenPrecedence() {
		binary, exists := p.binaryParseFunc[p.getNextToken().Type]
		if !exists {
			return left, nil
		}
		p.readToken()
		left, err = binary(left)
		if err != nil {
			return left, err
		}
	}

	return left, nil
}

func (p *parser) parseNumber() (*ast.Expression, error) {
	expr := &ast.Expression{
		Literal: &ast.Literal{
			Numeric: &ast.Numeric{
				Integral: p.currentToken.Value.Integer,
			},
		},
	}
	return expr, nil
}

func (p *parser) parseIdent() (*ast.Expression, error) {
	if p.getNextToken().Type == token.S_LPAREN {
		expr, err := p.parseFunctionCallExpr()
		if err != nil {
			return &ast.Expression{}, err
		}
		return expr, nil
	} else {
		expr, err := p.parseColumnExpr()
		if err != nil {
			return &ast.Expression{}, err
		}
		return expr, nil
	}
}

func (p *parser) parseBinaryExpr(left *ast.Expression) (*ast.Expression, error) {
	expr := &ast.Expression{
		BinaryOperation: &ast.BinaryOpe{
			Left: left,
		},
	}

	switch p.currentToken.Type {
	case token.S_PLUS:
		expr.BinaryOperation.Operator = ast.B_PLUS
	case token.S_MINUS:
		expr.BinaryOperation.Operator = ast.B_MINUS
	case token.S_ASTERISK:
		expr.BinaryOperation.Operator = ast.B_ASTERISK
	case token.S_SOLIDAS:
		expr.BinaryOperation.Operator = ast.B_SOLIDAS
	case token.S_PERCENT:
		expr.BinaryOperation.Operator = ast.B_PERCENT
	}
	precedence := p.getCurrentTokenPrecedence()

	p.readToken()

	ex, err := p.parseExpression(precedence)
	if err != nil {
		return expr, err
	}
	expr.BinaryOperation.Right = ex
	return expr, nil
}

func (p *parser) parseGroupedExpr() (*ast.Expression, error) {
	expr := &ast.Expression{}
	p.readToken()

	ex, err := p.parseExpression(LOWEST)
	if err != nil {
		return expr, err
	}
	if p.getNextToken().Type != token.S_RPAREN {
		return expr, errors.New("Unknown Format")
	}
	p.readToken()
	return ex, nil
}

func (p *parser) parsePrefixExpr() (*ast.Expression, error) {
	expr := &ast.Expression{
		UnaryOperation: &ast.UnaryOpe{},
	}
	switch p.currentToken.Type {
	case token.S_PLUS:
		expr.UnaryOperation.Operator = ast.U_PLUS
	case token.S_MINUS:
		expr.UnaryOperation.Operator = ast.U_MINUS
	default:
		return expr, errors.New("Unknown Prefix Operator")
	}

	p.readToken()
	ex, err := p.parseExpression(LOWEST)
	if err != nil {
		return expr, err
	}

	expr.UnaryOperation.Expr = ex

	return expr, nil
}

func (p *parser) parseFunctionCallExpr() (*ast.Expression, error) {
	expr := &ast.Expression{
		FunctionCall: &ast.FunctionCall{},
	}

	expr.FunctionCall.Name = strings.ToUpper(p.currentToken.Literal)

	for {
		p.readToken()
		ex, err := p.parseExpression(LOWEST)
		if err != nil {
			return expr, err
		}
		expr.FunctionCall.Args = append(expr.FunctionCall.Args, *ex)
		if p.currentToken.Type == token.S_RPAREN {
			break
		}
	}

	return expr, nil
}

func (p *parser) parseColumnExpr() (*ast.Expression, error) {
	expr := &ast.Expression{
		Column: &ast.Column{},
	}
	expr.Column.Column = p.currentToken.Literal

	return expr, nil
}
