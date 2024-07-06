package internal

import (
	_client "dgen/tpl/client"
	_server "dgen/tpl/server"
	"io/fs"
	"os"
	"path/filepath"
)

var ClientParser = &clientParser{}

type clientParser struct {
}

func (p *clientParser) Parse(entireModel *EntireModel) error {
	err := p.parseModel(entireModel)
	if err != nil {
		return err
	}

	err = p.parseClient(entireModel)
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
	err := parseNewFile(model, "model", _server.ModelTpl, entireModel)
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
	err := parseNewFile(client, "client", _client.ClientTpl, entireModel)
	if err != nil {
		return err
	}

	return nil
}
