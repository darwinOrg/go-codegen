package internal

import (
	_default "dgen/tpl/default"
	_server "dgen/tpl/server"
	dgcoll "github.com/darwinOrg/go-common/collection"
	"github.com/darwinOrg/go-common/utils"
	"github.com/iancoleman/strcase"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var ServerParser = &serverParser{}

type serverParser struct {
}

func (p *serverParser) Parse(entireModel *EntireModel) error {
	err := p.parseDal(entireModel)
	if err != nil {
		return err
	}

	err = p.parseEnum(entireModel)
	if err != nil {
		return err
	}

	err = p.parseModel(entireModel)
	if err != nil {
		return err
	}

	err = p.parseConverter(entireModel)
	if err != nil {
		return err
	}

	err = p.parseService(entireModel)
	if err != nil {
		return err
	}

	err = p.parseHandler(entireModel)
	if err != nil {
		return err
	}

	err = p.parseRouter(entireModel)
	if err != nil {
		return err
	}

	return nil
}

func (g *serverParser) parseDal(entireModel *EntireModel) error {
	if len(entireModel.Dbs) == 0 {
		return nil
	}

	sqls := dgcoll.MapToList(entireModel.Dbs, func(db *DbModelData) string { return db.Sql })
	dbSql := strings.Join(sqls, "\n")

	metas, err := BuildTableMetas(dbSql)
	if err != nil {
		return err
	}
	if len(metas) == 0 {
		return nil
	}

	dalDir := filepath.Join(entireModel.Export.ServerOutput, "dal")
	_ = os.MkdirAll(dalDir, fs.ModeDir|fs.ModePerm)

	for _, meta := range metas {
		meta.PackagePrefix = entireModel.Export.ServerPackagePrefix

		dalMain := filepath.Join(dalDir, meta.GoTable+".go")
		err := parseFile(dalMain, "dal-main", _default.DalMainTpl, meta)
		if err != nil {
			return err
		}

		dalExt := filepath.Join(dalDir, meta.GoTable+"-ext.go")
		if utils.ExistsFile(dalExt) {
			continue
		}

		err = parseFile(dalExt, "dal-ext", _default.DalExtTpl, meta)
		if err != nil {
			return err
		}
	}

	exec.Command("go", "fmt", dalDir)
	return nil
}

func (g *serverParser) parseEnum(entireModel *EntireModel) error {
	if len(entireModel.Enums) == 0 {
		return nil
	}

	enumDir := filepath.Join(entireModel.Export.ServerOutput, "enum")
	_ = os.MkdirAll(enumDir, fs.ModeDir|fs.ModePerm)

	enum := filepath.Join(enumDir, entireModel.FilePrefix+"_enum.go")
	if utils.ExistsFile(enum) {
		return nil
	}

	err := parseFile(enum, "enum", _server.EnumTpl, entireModel)
	if err != nil {
		return err
	}

	return nil
}

func (g *serverParser) parseModel(entireModel *EntireModel) error {
	if len(entireModel.Requests) == 0 && len(entireModel.Responses) == 0 {
		return nil
	}

	modelDir := filepath.Join(entireModel.Export.ServerOutput, "model")
	_ = os.MkdirAll(modelDir, fs.ModeDir|fs.ModePerm)

	model := filepath.Join(modelDir, entireModel.FilePrefix+"_model.go")
	err := parseFile(model, "model", _server.ModelTpl, entireModel)
	if err != nil {
		return err
	}

	return nil
}

func (g *serverParser) parseConverter(entireModel *EntireModel) error {
	if len(entireModel.Converters) == 0 {
		return nil
	}

	converterDir := filepath.Join(entireModel.Export.ServerOutput, "converter")
	_ = os.MkdirAll(converterDir, fs.ModeDir|fs.ModePerm)

	for _, c := range entireModel.Converters {
		converter := filepath.Join(converterDir, strcase.ToSnake(c.DbTableUpperCamel)+"_converter.go")
		if utils.ExistsFile(converter) {
			continue
		}

		err := parseFile(converter, "converter", _server.ConverterTpl, c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *serverParser) parseService(entireModel *EntireModel) error {
	if len(entireModel.Interfaces) == 0 {
		return nil
	}

	serviceDir := filepath.Join(entireModel.Export.ServerOutput, "service")
	_ = os.MkdirAll(serviceDir, fs.ModeDir|fs.ModePerm)

	for _, inter := range entireModel.Interfaces {
		service := filepath.Join(serviceDir, strcase.ToSnake(inter.Group)+"_service.go")
		if utils.ExistsFile(service) {
			continue
		}

		err := parseFile(service, "service", _server.ServiceTpl, inter)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *serverParser) parseHandler(entireModel *EntireModel) error {
	if len(entireModel.Interfaces) == 0 {
		return nil
	}

	handlerDir := filepath.Join(entireModel.Export.ServerOutput, "handler")
	_ = os.MkdirAll(handlerDir, fs.ModeDir|fs.ModePerm)

	handler := filepath.Join(handlerDir, entireModel.FilePrefix+"_handler.go")
	err := parseFile(handler, "handler", _server.HandlerTpl, entireModel)
	if err != nil {
		return err
	}

	return nil
}

func (g *serverParser) parseRouter(entireModel *EntireModel) error {
	if len(entireModel.Interfaces) == 0 {
		return nil
	}

	routerDir := filepath.Join(entireModel.Export.ServerOutput, "router")
	_ = os.MkdirAll(routerDir, fs.ModeDir|fs.ModePerm)

	router := filepath.Join(routerDir, entireModel.FilePrefix+"_router.go")
	err := parseFile(router, "router", _server.RouterTpl, entireModel)
	if err != nil {
		return err
	}

	return nil
}
