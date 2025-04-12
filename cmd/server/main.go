package main

import (
	"github.com/darwinOrg/go-codegen/parser"
)

func main() {
	entireModel := parser.InitEntireModel()
	entireModel.Fill(entireModel.Export.ServerPackagePrefix)
	_ = parser.ServerParser.ParseAll(entireModel)
}
