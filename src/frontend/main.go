package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	args := os.Args[1:]
	var asts []*ast.File
	for _, file := range args {
		dat, _ := os.ReadFile("./" + file)
		a, _ := parser.ParseFile(&token.FileSet{}, file, dat, 0)
		asts = append(asts, a)
	}
	ast.Print(&token.FileSet{}, asts)
}
