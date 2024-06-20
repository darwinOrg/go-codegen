package internal

import (
	_server "dgen/tpl/server"
	"github.com/iancoleman/strcase"
	"io/fs"
	"os"
	"path/filepath"
)

var ServerParser = &serverParser{}

type serverParser struct {
}

func (p *serverParser) Parse(entireModel *EntireModel, mark string) error {
	err := p.parseEnum(entireModel, mark)
	if err != nil {
		return err
	}

	err = p.parseModel(entireModel, mark)
	if err != nil {
		return err
	}

	err = p.parseService(entireModel, mark)
	if err != nil {
		return err
	}

	err = p.parseHandler(entireModel, mark)
	if err != nil {
		return err
	}

	err = p.parseRouter(entireModel, mark)
	if err != nil {
		return err
	}

	return nil
}

func (g *serverParser) parseEnum(entireModel *EntireModel, mark string) error {
	enumDir := filepath.Join(entireModel.Export.ServerOutput, "enum")
	_ = os.MkdirAll(enumDir, fs.ModeDir|fs.ModePerm)

	enum := filepath.Join(enumDir, strcase.ToSnake(mark)+"_enum.go")
	err := parseFile(enum, "enum", _server.EnumTpl, entireModel)
	if err != nil {
		return err
	}

	return nil
}

func (g *serverParser) parseModel(entireModel *EntireModel, mark string) error {
	modelDir := filepath.Join(entireModel.Export.ServerOutput, "model")
	_ = os.MkdirAll(modelDir, fs.ModeDir|fs.ModePerm)

	model := filepath.Join(modelDir, strcase.ToSnake(mark)+"_model.go")
	err := parseFile(model, "model", _server.ModelTpl, entireModel)
	if err != nil {
		return err
	}

	return nil
}

func (g *serverParser) parseService(entireModel *EntireModel, mark string) error {
	serviceDir := filepath.Join(entireModel.Export.ServerOutput, "service")
	_ = os.MkdirAll(serviceDir, fs.ModeDir|fs.ModePerm)

	for _, inter := range entireModel.Interfaces {
		service := filepath.Join(serviceDir, strcase.ToSnake(inter.Group)+"_service.go")
		err := parseFile(service, "service", _server.ServiceExtTpl, inter)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *serverParser) parseHandler(entireModel *EntireModel, mark string) error {
	handlerDir := filepath.Join(entireModel.Export.ServerOutput, "handler")
	_ = os.MkdirAll(handlerDir, fs.ModeDir|fs.ModePerm)

	handler := filepath.Join(handlerDir, strcase.ToSnake(mark)+"_handler.go")
	err := parseFile(handler, "handler", _server.HandlerExtTpl, entireModel)
	if err != nil {
		return err
	}

	return nil
}

func (g *serverParser) parseRouter(entireModel *EntireModel, mark string) error {
	routerDir := filepath.Join(entireModel.Export.ServerOutput, "router")
	_ = os.MkdirAll(routerDir, fs.ModeDir|fs.ModePerm)

	router := filepath.Join(routerDir, strcase.ToSnake(mark)+"_router.go")
	err := parseFile(router, "router", _server.RouterExtTpl, entireModel)
	if err != nil {
		return err
	}

	return nil
}
