package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func List(filePath string) ([]*ast.ImportSpec, error) {
	fset := token.NewFileSet() // positions are relative to fset

	// Parse the file containing this very example
	// but stop after processing the imports.
	f, err := parser.ParseFile(fset, filePath, nil, parser.ImportsOnly)
	if err != nil {
		return nil, err
	}

	return f.Imports, nil
}

func BasicImportString(i *ast.ImportSpec) string {
	if i.Name != nil {
		return fmt.Sprintf("%s %s", i.Name.Name, i.Path.Value)
	} else {
		return i.Path.Value
	}
}
