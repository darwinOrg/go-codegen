package internal

import (
	"os"
	"os/exec"
	"strings"
	"text/template"
)

var (
	// 创建函数映射
	funcs = template.FuncMap{
		"contains": strings.Contains,
	}
)

func parseNewFile(fileName string, tplName string, tplText string, data any) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return parseFile(fileName, file, tplName, tplText, data)
}

func parseAppendFile(fileName string, tplName string, tplText string, data any) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return parseFile(fileName, file, tplName, tplText, data)
}

func parseFile(fileName string, file *os.File, tplName string, tplText string, data any) error {
	t, err := template.New(tplName).Funcs(funcs).Parse(tplText)
	if err != nil {
		return err
	}

	err = t.Execute(file, data)
	if err != nil {
		return err
	}

	cmd := exec.Command("gofmt", "-w", fileName)
	_ = cmd.Run()

	return nil
}
