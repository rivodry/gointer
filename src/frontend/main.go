package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func ParseImport(args []string) [][]*ast.ImportSpec {
	var asts [][]*ast.ImportSpec

	for _, file := range args {
		dat, _ := os.ReadFile("./" + file)
		a, _ := parser.ParseFile(&token.FileSet{}, file, dat, 0)

		asts = append(asts, a.Imports)

	}
	return asts
}
func main() {
	args := os.Args[1:]
	asts := ParseImport(args)

	ast.Print(&token.FileSet{}, asts)
}
