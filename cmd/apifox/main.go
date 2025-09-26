package main

import (
	"fmt"
	"os"

	"github.com/darwinOrg/go-codegen/dgen"
	dglogger "github.com/darwinOrg/go-logger"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Please input correct params")
		os.Exit(1)
	}
	dglogger.GlobalDgLogger = dglogger.NewDgLogger(dglogger.WarnLevel, dglogger.DefaultTimestampFormat, os.Stdout)

	entireModel := dgen.InitEntireModel()
	entireModel.Fill(entireModel.Export.ServerPackagePrefix)

	dgen.SyncToApifox(entireModel, os.Args[1], os.Args[2])
}
