package compile

import (
	"fmt"
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

func BuildTableMata(sql string, baseRoot string, pkg string) error {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		fmt.Println(err)
		return err
	}

	targetPath := filepath.Join(baseRoot, pkg)

	err = os.Mkdir(targetPath, fs.ModeDir|fs.ModePerm)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, node := range stmtNodes {
		root, ok := node.(*ast.CreateTableStmt)
		if !ok {
			continue
		}
		meta := &Meta{Package: pkg}
		root.Accept(meta)
		if compile(targetPath, meta) != nil {
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

	t := template.New(tplName)
	t, err = t.Parse(tplText)
	if err != nil {
		return err
	}

	return t.Execute(file, meta)
}
