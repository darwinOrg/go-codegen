package compile

import (
	"dgen/tpl"
	"fmt"
	dgcoll "github.com/darwinOrg/go-common/collection"
	"github.com/darwinOrg/go-common/model"
	"github.com/iancoleman/strcase"
	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	_ "github.com/pingcap/tidb/parser/test_driver"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

var (
	// 创建函数映射
	funcs = template.FuncMap{
		"contains": strings.Contains,
	}

	enumRegexp = regexp.MustCompile(`\(([^)]+)\)`)
)

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

		meta.QueryColumns = dgcoll.FilterList(meta.Columns, func(c *Column) bool {
			return !dgcoll.Contains(ignoreQueryModelFieldNames, c.DbName)
		})

		meta.ModifyColumns = dgcoll.FilterList(meta.Columns, func(c *Column) bool {
			return !dgcoll.Contains(ignoreModifyModelFieldNames, c.DbName) &&
				!strings.Contains(c.DbName, "status") &&
				!strings.Contains(c.DbName, "state")
		})

		meta.CreateColumns = dgcoll.FilterList(meta.Columns, func(c *Column) bool {
			return !dgcoll.Contains(ignoreCreateModelFieldNames, c.DbName)
		})

		meta.EnumMap = map[*Column][]*model.KeyValuePair[string, string]{}

		for _, column := range meta.Columns {
			columnComment := strings.TrimSpace(column.Comment)
			if columnComment == "" {
				continue
			}
			columnComment = strings.ReplaceAll(columnComment, "（", "(")
			columnComment = strings.ReplaceAll(columnComment, "）", ")")
			columnComment = strings.ReplaceAll(columnComment, "，", ",")
			columnComment = strings.ReplaceAll(columnComment, "：", ":")

			matches := enumRegexp.FindStringSubmatch(columnComment)
			if len(matches) == 0 {
				continue
			}

			pairsStr := matches[1]
			pairParts := strings.Split(strings.TrimSpace(pairsStr), ",")

			for _, pairPart := range pairParts {
				kvs := strings.Split(strings.TrimSpace(pairPart), ":")

				if len(kvs) == 2 {
					key := strings.TrimSpace(kvs[0])
					value := strings.TrimSpace(kvs[1])
					meta.HasEnum = true
					column.HasEnum = true
					enumName := columnComment[:strings.Index(columnComment, "(")]
					enumRemark := columnComment[strings.Index(columnComment, "(")+1 : len(columnComment)-1]
					column.EnumName = enumName
					column.EnumRemark = enumRemark

					meta.EnumMap[column] = append(meta.EnumMap[column], &model.KeyValuePair[string, string]{
						Key:   key,
						Value: value,
					})
				}
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

	err = compileEnum(targetPath, meta)
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

func compileEnum(targetPath string, meta *Meta) error {
	if len(meta.EnumMap) == 0 {
		return nil
	}

	enumDir := filepath.Join(targetPath, "enum")
	_ = os.MkdirAll(enumDir, fs.ModeDir|fs.ModePerm)

	enum := filepath.Join(enumDir, strcase.ToSnake(meta.GoTable)+"_enum.go")
	err := compileFile(enum, "enum", tpl.EnumTpl, meta)
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
