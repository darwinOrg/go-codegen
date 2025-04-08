package main

import (
	"dgen/internal"
	"encoding/json"
	"fmt"
	dglogger "github.com/darwinOrg/go-logger"
	"github.com/darwinOrg/go-swagger"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Please input correct params")
		os.Exit(1)
	}
	dglogger.GlobalDgLogger = dglogger.NewDgLogger(dglogger.WarnLevel, dglogger.DefaultTimestampFormat, os.Stdout)

	entireModel := internal.InitEntireModel()
	entireModel.Fill(entireModel.Export.ServerPackagePrefix)
	swaggerProps := internal.BuildSwaggerProps(entireModel)
	swaggerJsonBytes, err := json.MarshalIndent(swaggerProps, "", "  ")
	if err != nil {
		log.Printf("json marshal error: %v", err)
	}

	swagger.SyncSwaggerJsonBytesToApifox(&swagger.SyncToApifoxRequest{
		AccessToken:         os.Args[1],
		ProjectId:           os.Args[2],
		ApiOverwriteMode:    swagger.ApiOverwriteModeMethodAndPath,
		SchemaOverwriteMode: swagger.SchemaOverwriteModeBoth,
		SyncApiFolder:       false,
		ImportBasePath:      false,
		ApiFolderPath:       entireModel.Export.ApifoxOutput,
	}, swaggerJsonBytes)
}
