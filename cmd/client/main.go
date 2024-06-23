package main

import (
	"dgen/internal"
)

func main() {
	entireModel := internal.InitEntireModel()
	entireModel.Fill(entireModel.Export.ClientPackagePrefix)
	internal.ClientParser.Parse(entireModel)
}
