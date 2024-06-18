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
	err := p.parseModel(entireModel, mark)
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

func (g *serverParser) parseHandler(entireModel *EntireModel, mark string) error {
	modelDir := filepath.Join(entireModel.Export.ServerOutput, "handler")
	_ = os.MkdirAll(modelDir, fs.ModeDir|fs.ModePerm)

	handler := filepath.Join(modelDir, strcase.ToSnake(mark)+"_handler.go")
	err := parseFile(handler, "handler", _server.HandlerExtTpl, entireModel)
	if err != nil {
		return err
	}

	return nil
}

func (g *serverParser) parseService(entireModel *EntireModel, mark string) error {
	modelDir := filepath.Join(entireModel.Export.ServerOutput, "service")
	_ = os.MkdirAll(modelDir, fs.ModeDir|fs.ModePerm)

	service := filepath.Join(modelDir, strcase.ToSnake(mark)+"_service.go")
	err := parseFile(service, "service", _server.ServiceExtTpl, entireModel)
	if err != nil {
		return err
	}

	return nil
}

func (g *serverParser) parseRouter(entireModel *EntireModel, mark string) error {
	modelDir := filepath.Join(entireModel.Export.ServerOutput, "router")
	_ = os.MkdirAll(modelDir, fs.ModeDir|fs.ModePerm)

	router := filepath.Join(modelDir, strcase.ToSnake(mark)+"_router.go")
	err := parseFile(router, "router", _server.RouterExtTpl, entireModel)
	if err != nil {
		return err
	}

	return nil
}
