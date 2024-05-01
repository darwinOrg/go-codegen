package main

import (
	"dgen/compile"
	"flag"
	"fmt"
	_ "github.com/darwinOrg/daog-ext"
	_ "github.com/darwinOrg/go-logger"
	_ "github.com/darwinOrg/go-web/wrapper"
	_ "github.com/gin-gonic/gin"
	_ "github.com/rolandhe/daog"
	_ "github.com/shopspring/decimal"
	"os"
	"os/exec"
)

var (
	inputFile   string
	projectPath string
	outputPath  string
)

func init() {
	flag.StringVar(&inputFile, "i", "./scripts/test.sql", "the create tables file")
	flag.StringVar(&projectPath, "p", "dgen/output", "project path for every go file")
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

	if projectPath == "" {
		fmt.Println("Please input project path")
		os.Exit(1)
	}
	fmt.Println("Using project path: ", projectPath)

	if outputPath == "./" {
		fmt.Println("Using current directory as output: ./")
	} else {
		fmt.Println("Using output file directory is: ", outputPath)
	}
	_ = compile.BuildTableMata(sql, projectPath, outputPath)

	if outputPath == "./" {
		outputPath, _ = os.Getwd()
	}

	dirs := []string{"dal", "enum", "model", "converter", "service", "handler", "router"}
	for _, dir := range dirs {
		cmd := exec.Command("go", "fmt", outputPath+"/"+dir)
		err = cmd.Run()
		if err != nil {
			fmt.Printf("go fmt %s error: %v", dir, err)
		}
	}
}
