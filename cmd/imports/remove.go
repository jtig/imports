package main

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"io/ioutil"
)

func Remove(filePath string, importPath string) error {
	parserMode := parser.Mode(0)
	parserMode |= parser.ParseComments
	parserMode |= parser.AllErrors
	fset := token.NewFileSet() // positions are relative to fset

	f, err := parser.ParseFile(fset, filePath, nil, parserMode)
	if err != nil {
		return err
	}
	astutil.DeleteImport(fset, f, importPath)

	printerMode := printer.UseSpaces
	printConfig := &printer.Config{Mode: printerMode, Tabwidth: 4}

	var buf bytes.Buffer
	err = printConfig.Fprint(&buf, fset, f)
	if err != nil {
		return err
	}
	out := buf.Bytes()

	out, err = format.Source(out)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, out, 0)
	if err != nil {
		return err
	}
	return nil
}
