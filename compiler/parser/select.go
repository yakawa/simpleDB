package parser

import (
	"errors"

	"github.com/yakawa/simpleDB/common/ast"
	"github.com/yakawa/simpleDB/common/token"
)

func (p *parser) parseSELECTStatement() (*ast.SELECTStatement, error) {
	statement := &ast.SELECTStatement{}

	if p.currentToken.Type == token.K_SELECT {
		selectClause, err := p.parseSELECTClause()
		if err != nil {
			return statement, err
		}
		statement.Select = selectClause
		if p.currentToken.Type == token.K_FROM {
			p.readToken()
			fromClause, err := p.parseFROMClause()
			if err != nil {
				return statement, err
			}
			statement.From = fromClause
		}
	} else {
		return statement, errors.New("SELECT missing")
	}
	return statement, nil
}

func (p *parser) parseSELECTClause() (*ast.SELECTClause, error) {
	clause := &ast.SELECTClause{}
	p.readToken()

	cols, err := p.parseResultColumns()
	if err != nil {
		return clause, err
	}
	clause.ResultColumns = cols
	return clause, nil
}

func (p *parser) parseResultColumns() ([]ast.ResultColumn, error) {
	cols := []ast.ResultColumn{}
	loop := true
	for {
		switch p.currentToken.Type {
		case token.EOS, token.S_SEMICOLON, token.K_FROM:
			loop = false
		case token.S_COMMA:
			p.readToken()
		default:
			expr, err := p.parseExpression(LOWEST)
			if err != nil {
				return cols, err
			}
			cols = append(cols, ast.ResultColumn{Expression: expr})
			p.readToken()
		}
		if !loop {
			break
		}
	}
	return cols, nil
}
