package parser

import "github.com/yakawa/simpleDB/common/ast"

func (p *parser) parseFROMClause() (*ast.FROMClause, error) {
	from := &ast.FROMClause{}

	tbl, err := p.parseTable()
	if err != nil {
		return from, err
	}

	from.Table = tbl

	p.readToken()

	return from, nil
}

func (p *parser) parseTable() (*ast.Table, error) {
	tbl := &ast.Table{}

	tbl.Table = p.currentToken.Literal

	return tbl, nil
}
