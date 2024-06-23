package main

import (
	"dgen/internal"
)

func main() {
	entireModel := internal.InitEntireModel()
	entireModel.Fill(entireModel.Export.ServerPackagePrefix)
	internal.ServerParser.Parse(entireModel)
}
