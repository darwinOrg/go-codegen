package main

import (
	"flag"
	"fmt"
	"github.com/darwinOrg/go-codegen/compile"
	"os"
	"os/exec"
	"path/filepath"
)

var inputFile string
var oPath string
var pkgName string

func init() {
	flag.StringVar(&inputFile, "i", "", "the create tables file")
	flag.StringVar(&oPath, "o", "./", "output directory, default is current directory")
	flag.StringVar(&pkgName, "pkg", "", "package name for every go file")
}

func main() {
	flag.Parse()
	if inputFile == "" {
		fmt.Println("Please input create tables file")
		os.Exit(1)
	}
	fmt.Println("Using create table files is: ", inputFile)
	if pkgName == "" {
		fmt.Println("Please input package name")
		os.Exit(1)
	}
	fmt.Println("Using package name: ", pkgName)
	if oPath == "./" {
		fmt.Println("Using current directory as output: ./", pkgName)
	} else {
		fmt.Println("Using output file directory is: ", filepath.Join(oPath, pkgName))
	}
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Read input create tables file error.", err)
		os.Exit(1)
	}

	sql := string(data)

	_ = compile.BuildTableMata(sql, oPath, pkgName)

	targetPath := oPath
	if oPath == "./" {
		str, _ := os.Getwd()
		targetPath = str + "/" + pkgName
	}

	cmd := exec.Command("go", "fmt", targetPath)
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
