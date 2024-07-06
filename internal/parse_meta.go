package internal

import (
	_default "dgen/tpl/default"
	"github.com/iancoleman/strcase"
	"io/fs"
	"os"
	"path/filepath"
)

func parseMeta(targetPath string, meta *Meta) error {
	err := parseMetaDAL(targetPath, meta)
	if err != nil {
		return err
	}

	err = parseMetaEnum(targetPath, meta)
	if err != nil {
		return err
	}

	err = parseMetaModel(targetPath, meta)
	if err != nil {
		return err
	}

	err = parseMetaConverter(targetPath, meta)
	if err != nil {
		return err
	}

	err = parseMetaService(targetPath, meta)
	if err != nil {
		return err
	}

	err = parseMetaHandler(targetPath, meta)
	if err != nil {
		return err
	}

	err = parseMetaRouter(targetPath, meta)
	if err != nil {
		return err
	}

	return nil
}

func parseMetaDAL(targetPath string, meta *Meta) error {
	dalDir := filepath.Join(targetPath, "dal")
	_ = os.MkdirAll(dalDir, fs.ModeDir|fs.ModePerm)

	dalMain := filepath.Join(dalDir, meta.GoTable+".go")
	err := parseNewFile(dalMain, "dal-main", _default.DalMainTpl, meta)
	if err != nil {
		return err
	}

	dalExt := filepath.Join(dalDir, meta.GoTable+"-ext.go")
	err = parseNewFile(dalExt, "dal-ext", _default.DalExtTpl, meta)
	if err != nil {
		return err
	}

	return nil
}

func parseMetaEnum(targetPath string, meta *Meta) error {
	if len(meta.EnumMap) == 0 {
		return nil
	}

	enumDir := filepath.Join(targetPath, "enum")
	_ = os.MkdirAll(enumDir, fs.ModeDir|fs.ModePerm)

	enum := filepath.Join(enumDir, strcase.ToSnake(meta.GoTable)+"_enum.go")
	err := parseNewFile(enum, "enum", _default.EnumTpl, meta)
	if err != nil {
		return err
	}

	return nil
}

func parseMetaModel(targetPath string, meta *Meta) error {
	modelDir := filepath.Join(targetPath, "model")
	_ = os.MkdirAll(modelDir, fs.ModeDir|fs.ModePerm)

	model := filepath.Join(modelDir, strcase.ToSnake(meta.GoTable)+"_model.go")
	err := parseNewFile(model, "model", _default.ModelTpl, meta)
	if err != nil {
		return err
	}

	return nil
}

func parseMetaConverter(targetPath string, meta *Meta) error {
	converterDir := filepath.Join(targetPath, "converter")
	_ = os.MkdirAll(converterDir, fs.ModeDir|fs.ModePerm)

	converter := filepath.Join(converterDir, strcase.ToSnake(meta.GoTable)+"_converter.go")
	err := parseNewFile(converter, "converter", _default.ConverterTpl, meta)
	if err != nil {
		return err
	}

	return nil
}

func parseMetaService(targetPath string, meta *Meta) error {
	serviceDir := filepath.Join(targetPath, "service")
	_ = os.MkdirAll(serviceDir, fs.ModeDir|fs.ModePerm)

	service := filepath.Join(serviceDir, strcase.ToSnake(meta.GoTable)+"_service.go")
	err := parseNewFile(service, "service", _default.ServiceTpl, meta)
	if err != nil {
		return err
	}

	return nil
}

func parseMetaHandler(targetPath string, meta *Meta) error {
	handlerDir := filepath.Join(targetPath, "handler")
	_ = os.MkdirAll(handlerDir, fs.ModeDir|fs.ModePerm)

	handler := filepath.Join(handlerDir, strcase.ToSnake(meta.GoTable)+"_handler.go")
	err := parseNewFile(handler, "handler", _default.HandlerTpl, meta)
	if err != nil {
		return err
	}

	return nil
}

func parseMetaRouter(targetPath string, meta *Meta) error {
	routerDir := filepath.Join(targetPath, "router")
	_ = os.MkdirAll(routerDir, fs.ModeDir|fs.ModePerm)

	router := filepath.Join(routerDir, strcase.ToSnake(meta.GoTable)+"_router.go")
	err := parseNewFile(router, "router", _default.RouterTpl, meta)
	if err != nil {
		return err
	}

	return nil
}
