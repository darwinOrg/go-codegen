package main

import (
	"flag"
	"fmt"
	_ "github.com/darwinOrg/daog-ext"
	"github.com/darwinOrg/go-codegen/parser"
	_ "github.com/darwinOrg/go-httpclient"
	_ "github.com/darwinOrg/go-logger"
	_ "github.com/darwinOrg/go-web/wrapper"
	_ "github.com/gin-gonic/gin"
	_ "github.com/rolandhe/daog"
	_ "github.com/shopspring/decimal"
	"os"
)

var (
	inputFile     string
	packagePrefix string
	outputPath    string
)

func init() {
	flag.StringVar(&inputFile, "i", "./scripts/test.sql", "the create tables file")
	flag.StringVar(&packagePrefix, "p", "github.com/darwinOrg/go-codegen/output", "package prefix for every go file")
	flag.StringVar(&outputPath, "o", "./output", "output directory")
}

func main() {
	flag.Parse()

	if inputFile == "" {
		fmt.Println("Please input create tables file")
		os.Exit(1)
	}
	fmt.Println("Using create table files is: ", inputFile)
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Read input create tables file error.", err)
		os.Exit(1)
	}
	sql := string(data)

	if packagePrefix == "" {
		fmt.Println("Please input package prefix")
		os.Exit(1)
	}
	fmt.Println("Using package prefix: ", packagePrefix)

	if outputPath == "./" {
		fmt.Println("Using current directory as output: ./")
	} else {
		fmt.Println("Using output file directory is: ", outputPath)
	}
	_ = parser.ParseSql(sql, packagePrefix, outputPath)
}
