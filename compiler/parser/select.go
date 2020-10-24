package parser

import (
	"github.com/yakawa/simpleDB/common/ast"
	"github.com/yakawa/simpleDB/common/value"
)

func (p *parser) parseSELECTStatement() (*ast.SELECTStatement, error) {
	statement := &ast.SELECTStatement{}
	loop := true

	for {
		switch p.currentToken.Value.Type {
		case value.K_SELECT:
			selectClause, err := p.parseSELECTClause()
			if err != nil {
				return statement, err
			}
			statement.Select = selectClause
		default:
			loop = false
		}
		if !loop {
			break
		}
	}
	return statement, nil
}

func (p *parser) parseSELECTClause() (*ast.SELECTClause, error) {
	clause := &ast.SELECTClause{}
	loop := true
	p.readToken()

	for {
		switch p.currentToken.Value.Type {
		case value.INTEGER:
			cols, err := p.parseResultColumns()
			if err != nil {
				return clause, err
			}
			clause.ResultColumns = cols
		default:
			loop = false
		}
		if !loop {
			break
		}
	}
	return clause, nil
}

func (p *parser) parseResultColumns() ([]ast.ResultColumn, error) {
	cols := []ast.ResultColumn{}
	loop := true
	for {
		switch p.currentToken.Value.Type {
		case value.INTEGER:
			expr, err := p.parseExpression(LOWEST)
			if err != nil {
				return cols, err
			}
			cols = append(cols, ast.ResultColumn{Expression: expr})
			p.readToken()
		default:
			loop = false
			break
		}
		if !loop {
			break
		}
	}
	return cols, nil
}
