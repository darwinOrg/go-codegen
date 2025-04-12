package main

import (
	"github.com/darwinOrg/go-codegen/dgen"
)

func main() {
	entireModel := dgen.InitEntireModel()
	entireModel.Fill(entireModel.Export.ServerPackagePrefix)
	_ = dgen.ServerParser.ParseAll(entireModel)
}
