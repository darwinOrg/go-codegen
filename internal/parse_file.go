package internal

import (
	"github.com/darwinOrg/go-common/utils"
	"os"
	"strings"
	"text/template"
)

var (
	// 创建函数映射
	funcs = template.FuncMap{
		"contains": strings.Contains,
	}
)

func parseFile(fileName string, tplName string, tplText string, data any) error {
	if utils.ExistsFile(fileName) {
		return nil
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	t, err := template.New(tplName).Funcs(funcs).Parse(tplText)
	if err != nil {
		return err
	}

	return t.Execute(file, data)
}
