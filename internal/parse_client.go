package internal

import (
	_client "dgen/tpl/client"
	_server "dgen/tpl/server"
	"github.com/darwinOrg/go-common/utils"
	"io/fs"
	"os"
	"path/filepath"
)

var ClientParser = &clientParser{}

type clientParser struct {
}

func (p *clientParser) Parse(entireModel *EntireModel) error {
	err := p.parseEnum(entireModel)
	if err != nil {
		return err
	}

	err = p.parseModel(entireModel)
	if err != nil {
		return err
	}

	err = p.parseClient(entireModel)
	if err != nil {
		return err
	}

	return nil
}

func (g *clientParser) parseEnum(entireModel *EntireModel) error {
	if len(entireModel.Enums) == 0 {
		return nil
	}

	enumDir := filepath.Join(entireModel.Export.ClientOutput, "enum")
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

func (g *clientParser) parseModel(entireModel *EntireModel) error {
	if len(entireModel.Requests) == 0 && len(entireModel.Responses) == 0 {
		return nil
	}

	modelDir := filepath.Join(entireModel.Export.ClientOutput, "model")
	_ = os.MkdirAll(modelDir, fs.ModeDir|fs.ModePerm)

	model := filepath.Join(modelDir, entireModel.FilePrefix+"_model.go")
	err := parseFile(model, "model", _server.ModelTpl, entireModel)
	if err != nil {
		return err
	}

	return nil
}

func (g *clientParser) parseClient(entireModel *EntireModel) error {
	if len(entireModel.Interfaces) == 0 {
		return nil
	}

	clientDir := filepath.Join(entireModel.Export.ClientOutput, "client")
	_ = os.MkdirAll(clientDir, fs.ModeDir|fs.ModePerm)

	client := filepath.Join(clientDir, entireModel.FilePrefix+"_client.go")
	err := parseFile(client, "client", _client.ClientTpl, entireModel)
	if err != nil {
		return err
	}

	return nil
}
