package main

import (
	"flag"
	"fmt"
	_ "github.com/darwinOrg/daog-ext"
	"github.com/darwinOrg/go-codegen/compile"
	_ "github.com/darwinOrg/go-logger"
	_ "github.com/gin-gonic/gin"
	_ "github.com/rolandhe/daog"
	_ "github.com/shopspring/decimal"
	"os"
	"os/exec"
)

var (
	projectPath string
	inputFile   string
	outputPath  string
)

func init() {
	flag.StringVar(&projectPath, "p", "github.com/darwinOrg/go-codegen/output", "project path for every go file")
	flag.StringVar(&inputFile, "i", "./scripts/test.sql", "the create tables file")
	flag.StringVar(&outputPath, "o", "./output", "output directory, default is current directory")
}

func main() {
	flag.Parse()

	if projectPath == "" {
		fmt.Println("Please input project path")
		os.Exit(1)
	}
	fmt.Println("Using project path: ", projectPath)

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

	if outputPath == "./" {
		fmt.Println("Using current directory as output: ./")
	} else {
		fmt.Println("Using output file directory is: ", outputPath)
	}
	_ = compile.BuildTableMata(sql, projectPath, outputPath)

	if outputPath == "./" {
		outputPath, _ = os.Getwd()
	}
	cmd := exec.Command("go", "fmt", outputPath)
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
