package parser

import (
	"github.com/yakawa/simpleDB/common/ast"
	"github.com/yakawa/simpleDB/common/token"
	"github.com/yakawa/simpleDB/common/value"
)

type parser struct {
	tokens       token.Tokens
	currentToken token.Token
	pos          int

	unaryParseFunc  map[value.Type]unaryOpeFunction
	binaryParseFunc map[value.Type]binaryOpeFunction
}

func Parse(tokens token.Tokens) (*ast.AST, error) {
	a := &ast.AST{}
	p := new(tokens)
	sql, err := p.parse()
	if err != nil {
		return a, err
	}
	a.SQL = sql
	return a, nil
}

type (
	unaryOpeFunction  func() (*ast.Expression, error)
	binaryOpeFunction func(*ast.Expression) (*ast.Expression, error)
)

const (
	_ int = iota
	LOWEST
	SUM     // + -
	PRODUCT // * /
	HIGHEST
)

var precedences = map[value.Type]int{
	value.S_PLUS:     SUM,
	value.S_MINUS:    SUM,
	value.S_ASTERISK: PRODUCT,
	value.S_SOLIDAS:  PRODUCT,
	value.S_PERCENT:  PRODUCT,
}

func new(tokens token.Tokens) *parser {
	p := &parser{
		tokens: tokens,
	}
	p.readToken()

	p.unaryParseFunc = make(map[value.Type]unaryOpeFunction)
	p.binaryParseFunc = make(map[value.Type]binaryOpeFunction)

	p.unaryParseFunc[value.INTEGER] = p.parseInteger

	p.binaryParseFunc[value.S_PLUS] = p.parseBinaryExpr
	p.binaryParseFunc[value.S_MINUS] = p.parseBinaryExpr
	p.binaryParseFunc[value.S_ASTERISK] = p.parseBinaryExpr
	p.binaryParseFunc[value.S_SOLIDAS] = p.parseBinaryExpr
	p.binaryParseFunc[value.S_PERCENT] = p.parseBinaryExpr

	return p
}

func (p *parser) readToken() {
	if p.pos >= len(p.tokens) {
		p.currentToken = token.Token{
			Type: token.EOS,
		}
		return
	}
	p.currentToken = p.tokens[p.pos]
	p.pos++
	return
}

func (p *parser) getNextToken() token.Token {
	if p.pos >= len(p.tokens) {
		return token.Token{
			Type: token.EOS,
			Value: value.Value{
				Type: value.EOS,
			},
		}
	}
	return p.tokens[p.pos]
}

func (p *parser) getCurrentTokenPrecedence() int {
	if p.pos > len(p.tokens) {
		return LOWEST
	}
	if p, exists := precedences[p.currentToken.Value.Type]; exists {
		return p
	}
	return LOWEST
}

func (p *parser) getNextTokenPrecedence() int {
	if p.pos+1 > len(p.tokens) {
		return LOWEST
	}
	if p, exists := precedences[p.getNextToken().Value.Type]; exists {
		return p
	}
	return LOWEST
}

func (p *parser) parse() ([]ast.SQL, error) {
	SQLs := []ast.SQL{}
	loop := true
	for {
		switch p.currentToken.Value.Type {
		case value.K_SELECT:
			ss, err := p.parseSELECTStatement()
			if err != nil {
				return SQLs, err
			}
			sql := ast.SQL{
				SELECTStatement: ss,
			}
			SQLs = append(SQLs, sql)
		default:
			loop = false
		}
		if !loop {
			break
		}
	}
	return SQLs, nil
}
