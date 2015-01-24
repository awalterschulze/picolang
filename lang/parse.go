package lang

import (
	"github.com/awalterschulze/picolang/lang/ast"
	"github.com/awalterschulze/picolang/lang/lexer"
	"github.com/awalterschulze/picolang/lang/parser"
)

func Parse(s string) ([]ast.Statement, error) {
	p := parser.NewParser()
	ss, err := p.Parse(lexer.NewLexer([]byte(s)))
	if err != nil {
		return nil, err
	}
	return ss.([]ast.Statement), nil
}
