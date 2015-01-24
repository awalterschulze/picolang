
package parser

import (
	"github.com/awalterschulze/picolang/lang/ast"
)

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab {
	ProdTabEntry{
		String: `S' : Statements	<<  >>`,
		Id: "S'",
		NTType: 0,
		Index: 0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statements : Statements Statement	<< append(X[0].([]ast.Statement), X[1].(ast.Statement)), nil >>`,
		Id: "Statements",
		NTType: 1,
		Index: 1,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]ast.Statement), X[1].(ast.Statement)), nil
		},
	},
	ProdTabEntry{
		String: `Statements : Statement	<< []ast.Statement{X[0].(ast.Statement)}, nil >>`,
		Id: "Statements",
		NTType: 1,
		Index: 2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []ast.Statement{X[0].(ast.Statement)}, nil
		},
	},
	ProdTabEntry{
		String: `Statement : Let	<<  >>`,
		Id: "Statement",
		NTType: 2,
		Index: 3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : Import	<<  >>`,
		Id: "Statement",
		NTType: 2,
		Index: 4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : Println	<<  >>`,
		Id: "Statement",
		NTType: 2,
		Index: 5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : Expr	<<  >>`,
		Id: "Statement",
		NTType: 2,
		Index: 6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Let : "let" id "=" Expr	<< &ast.Let{ast.ToString(X[1]), X[3].(*ast.Expr)}, nil >>`,
		Id: "Let",
		NTType: 3,
		Index: 7,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Let{ast.ToString(X[1]), X[3].(*ast.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `Import : "import" id string_lit	<< &ast.Import{ast.ToString(X[1]), ast.ToString(X[2])}, nil >>`,
		Id: "Import",
		NTType: 4,
		Index: 8,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Import{ast.ToString(X[1]), ast.ToString(X[2])}, nil
		},
	},
	ProdTabEntry{
		String: `Println : "println(" Expr ")"	<< &ast.Println{X[1].(*ast.Expr)}, nil >>`,
		Id: "Println",
		NTType: 5,
		Index: 9,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Println{X[1].(*ast.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : id "(" ")"	<< &ast.Expr{ast.ToString(X[0]), []*ast.Param{}, nil}, nil >>`,
		Id: "Expr",
		NTType: 6,
		Index: 10,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Expr{ast.ToString(X[0]), []*ast.Param{}, nil}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : id	<< &ast.Expr{ast.ToString(X[0]), nil, nil}, nil >>`,
		Id: "Expr",
		NTType: 6,
		Index: 11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Expr{ast.ToString(X[0]), nil, nil}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : id "(" Params ")"	<< &ast.Expr{ast.ToString(X[0]), X[2].([]*ast.Param), nil}, nil >>`,
		Id: "Expr",
		NTType: 6,
		Index: 12,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Expr{ast.ToString(X[0]), X[2].([]*ast.Param), nil}, nil
		},
	},
	ProdTabEntry{
		String: `Params : Params "," Param	<< append(X[0].([]*ast.Param), X[2].(*ast.Param)), nil >>`,
		Id: "Params",
		NTType: 7,
		Index: 13,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]*ast.Param), X[2].(*ast.Param)), nil
		},
	},
	ProdTabEntry{
		String: `Params : Param	<< []*ast.Param{X[0].(*ast.Param)}, nil >>`,
		Id: "Params",
		NTType: 7,
		Index: 14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []*ast.Param{X[0].(*ast.Param)}, nil
		},
	},
	ProdTabEntry{
		String: `Param : id	<< &ast.Param{Id: ast.ToStringPtr(X[0])}, nil >>`,
		Id: "Param",
		NTType: 8,
		Index: 15,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Param{Id: ast.ToStringPtr(X[0])}, nil
		},
	},
	ProdTabEntry{
		String: `Param : string_lit	<< &ast.Param{Str: ast.ToStringPtr(X[0])}, nil >>`,
		Id: "Param",
		NTType: 8,
		Index: 16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Param{Str: ast.ToStringPtr(X[0])}, nil
		},
	},
	ProdTabEntry{
		String: `Param : float_lit	<< &ast.Param{Float: ast.ToFloat64(X[0])}, nil >>`,
		Id: "Param",
		NTType: 8,
		Index: 17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Param{Float: ast.ToFloat64(X[0])}, nil
		},
	},
	ProdTabEntry{
		String: `Param : singed_int_lit	<< &ast.Param{Float: ast.ToFloat64(X[0])}, nil >>`,
		Id: "Param",
		NTType: 8,
		Index: 18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Param{Float: ast.ToFloat64(X[0])}, nil
		},
	},
	ProdTabEntry{
		String: `Param : "{" Params "}"	<< &ast.Param{List: X[1].([]*ast.Param)}, nil >>`,
		Id: "Param",
		NTType: 8,
		Index: 19,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Param{List: X[1].([]*ast.Param)}, nil
		},
	},
	
}
