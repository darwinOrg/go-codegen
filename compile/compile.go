package compile

import (
	"fmt"
	"github.com/darwinOrg/go-codegen/functions"
	"github.com/darwinOrg/go-codegen/tpl"
	"github.com/iancoleman/strcase"
	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	_ "github.com/pingcap/tidb/parser/test_driver"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"
)

// 创建函数映射
var funcs = template.FuncMap{
	"contains": functions.Contains,
}

func BuildTableMata(sql string, projectPath string, outputPath string) error {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = os.Mkdir(outputPath, fs.ModeDir|fs.ModePerm)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, node := range stmtNodes {
		root, ok := node.(*ast.CreateTableStmt)
		if !ok {
			continue
		}
		meta := &Meta{ProjectPath: projectPath}
		root.Accept(meta)
		if compile(outputPath, meta) != nil {
			return err
		}
	}

	return nil
}

func compile(targetPath string, meta *Meta) error {
	err := compileDAL(targetPath, meta)
	if err != nil {
		return err
	}

	err = compileModel(targetPath, meta)
	if err != nil {
		return err
	}

	return nil
}

func compileDAL(targetPath string, meta *Meta) error {
	dalDir := filepath.Join(targetPath, "dal")
	_ = os.MkdirAll(dalDir, fs.ModeDir|fs.ModePerm)

	dalMain := filepath.Join(dalDir, strcase.ToSnake(meta.GoTable)+".go")
	err := compileFile(dalMain, "dal-main", tpl.DalMainTpl, meta)
	if err != nil {
		return err
	}

	dalExt := filepath.Join(dalDir, strcase.ToSnake(meta.GoTable)+"-ext.go")
	err = compileFile(dalExt, "dal-ext", tpl.DalExtTpl, meta)
	if err != nil {
		return err
	}

	return nil
}

func compileModel(targetPath string, meta *Meta) error {
	modelDir := filepath.Join(targetPath, "model")
	_ = os.MkdirAll(modelDir, fs.ModeDir|fs.ModePerm)

	model := filepath.Join(modelDir, strcase.ToSnake(meta.GoTable)+"_model.go")
	err := compileFile(model, "model", tpl.ModelTpl, meta)
	if err != nil {
		return err
	}

	return nil
}

func compileFile(fileName string, tplName string, tplText string, meta *Meta) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	t, err := template.New(tplName).Funcs(funcs).Parse(tplText)
	if err != nil {
		return err
	}

	return t.Execute(file, meta)
}
