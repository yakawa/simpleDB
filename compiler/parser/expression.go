package parser

import (
	"errors"

	"github.com/yakawa/simpleDB/common/ast"
	"github.com/yakawa/simpleDB/common/value"
)

func (p *parser) parseExpression(precedence int) (*ast.Expression, error) {
	expr := &ast.Expression{}
	unary, exists := p.unaryParseFunc[p.currentToken.Value.Type]
	if !exists {
		return expr, errors.New("Unknown Unary Operator")
	}

	left, err := unary()
	if err != nil {
		return expr, err
	}

	for p.getNextToken().Value.Type != value.EOS && precedence < p.getNextTokenPrecedence() {
		binary, exists := p.binaryParseFunc[p.getNextToken().Value.Type]
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

func (p *parser) parseInteger() (*ast.Expression, error) {
	expr := &ast.Expression{
		Literal: &ast.Literal{
			Numeric: &ast.Numeric{
				Integral: p.currentToken.Value.Integer,
			},
		},
	}
	return expr, nil
}

func (p *parser) parseBinaryExpr(left *ast.Expression) (*ast.Expression, error) {
	expr := &ast.Expression{
		BinaryOperation: &ast.BinaryOpe{
			Left: left,
		},
	}

	switch p.currentToken.Value.Type {
	case value.S_PLUS:
		expr.BinaryOperation.Operator = ast.B_PLUS
	case value.S_MINUS:
		expr.BinaryOperation.Operator = ast.B_MINUS
	case value.S_ASTERISK:
		expr.BinaryOperation.Operator = ast.B_ASTERISK
	case value.S_SOLIDAS:
		expr.BinaryOperation.Operator = ast.B_SOLIDAS
	case value.S_PERCENT:
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
