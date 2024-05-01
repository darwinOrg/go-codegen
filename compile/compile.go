package compile

import (
	"dgen/tpl"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	_ "github.com/pingcap/tidb/parser/test_driver"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// 创建函数映射
var funcs = template.FuncMap{
	"contains": strings.Contains,
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

		for _, option := range root.Options {
			if option.Tp == ast.TableOptionComment && option.StrValue != "" {
				meta.TableComment = option.StrValue
			}
		}

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

	err = compileConverter(targetPath, meta)
	if err != nil {
		return err
	}

	err = compileService(targetPath, meta)
	if err != nil {
		return err
	}

	err = compileHandler(targetPath, meta)
	if err != nil {
		return err
	}

	err = compileRouter(targetPath, meta)
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

func compileConverter(targetPath string, meta *Meta) error {
	converterDir := filepath.Join(targetPath, "converter")
	_ = os.MkdirAll(converterDir, fs.ModeDir|fs.ModePerm)

	converter := filepath.Join(converterDir, strcase.ToSnake(meta.GoTable)+"_converter.go")
	err := compileFile(converter, "converter", tpl.ConverterExtTpl, meta)
	if err != nil {
		return err
	}

	return nil
}

func compileService(targetPath string, meta *Meta) error {
	serviceDir := filepath.Join(targetPath, "service")
	_ = os.MkdirAll(serviceDir, fs.ModeDir|fs.ModePerm)

	service := filepath.Join(serviceDir, strcase.ToSnake(meta.GoTable)+"_service.go")
	err := compileFile(service, "service", tpl.ServiceExtTpl, meta)
	if err != nil {
		return err
	}

	return nil
}

func compileHandler(targetPath string, meta *Meta) error {
	handlerDir := filepath.Join(targetPath, "handler")
	_ = os.MkdirAll(handlerDir, fs.ModeDir|fs.ModePerm)

	handler := filepath.Join(handlerDir, strcase.ToSnake(meta.GoTable)+"_handler.go")
	err := compileFile(handler, "handler", tpl.HandlerExtTpl, meta)
	if err != nil {
		return err
	}

	return nil
}

func compileRouter(targetPath string, meta *Meta) error {
	routerDir := filepath.Join(targetPath, "router")
	_ = os.MkdirAll(routerDir, fs.ModeDir|fs.ModePerm)

	router := filepath.Join(routerDir, strcase.ToSnake(meta.GoTable)+"_router.go")
	err := compileFile(router, "router", tpl.RouterExtTpl, meta)
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
