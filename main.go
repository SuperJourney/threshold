package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	expr := "a + b + c"

	// Parse the expression into an AST
	astTree, err := parseExpr(expr)
	if err != nil {
		fmt.Println("Error parsing expression:", err)
		return
	}

	// Print the AST
	ast.Print(nil, astTree)
}

func parseExpr(expr string) (ast.Node, error) {
	// Create a file set
	fset := token.NewFileSet()

	// Parse the expression into an AST
	exprFile, err := parser.ParseExprFrom(fset, "expr", expr, 0)
	if err != nil {
		return nil, err
	}

	return exprFile, nil
}
