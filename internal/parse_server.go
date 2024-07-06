package internal

import (
	_default "dgen/tpl/default"
	_server "dgen/tpl/server"
	"fmt"
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

	err = p.parseRouter(entireModel)
	if err != nil {
		return err
	}

	err = p.parseHandler(entireModel)
	if err != nil {
		return err
	}

	err = p.parseService(entireModel)
	if err != nil {
		return err
	}

	err = p.parseConverter(entireModel)
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
		err := parseNewFile(dalMain, "dal-main", _default.DalMainTpl, meta)
		if err != nil {
			return err
		}

		dalExt := filepath.Join(dalDir, meta.GoTable+"-ext.go")
		if utils.ExistsFile(dalExt) {
			continue
		}

		err = parseNewFile(dalExt, "dal-ext", _default.DalExtTpl, meta)
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

	err := parseNewFile(enum, "enum", _server.EnumTpl, entireModel)
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
	err := parseNewFile(model, "model", _server.ModelTpl, entireModel)
	if err != nil {
		return err
	}

	//request := filepath.Join(modelDir, entireModel.FilePrefix+"_request.go")
	//err := parseNewFile(request, "request", _server.RequestTpl, entireModel)
	//if err != nil {
	//	return err
	//}
	//
	//response := filepath.Join(modelDir, entireModel.FilePrefix+"_response.go")
	//err = parseNewFile(response, "response", _server.ResponseTpl, entireModel)
	//if err != nil {
	//	return err
	//}

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
			fileBytes, err := os.ReadFile(converter)
			if err != nil {
				return err
			}
			fileString := string(fileBytes)

			nc, err := utils.ConvertToNewBeanByJson[ConverterData](c)
			if err != nil {
				return err
			}

			nc.Requests = dgcoll.FilterList(nc.Requests, func(rmd *RequestModelData) bool {
				return !strings.Contains(fileString, fmt.Sprintf("%sConverter) %s2Entity(", c.DbTableLowerCamel, rmd.UpperCamelName)) &&
					!strings.Contains(fileString, fmt.Sprintf("%sConverter) %s2Param(", c.DbTableLowerCamel, rmd.UpperCamelName)) &&
					!strings.Contains(fileString, fmt.Sprintf("%sConverter) FillEntityWith%s(", c.DbTableLowerCamel, rmd.UpperCamelName))
			})
			nc.Requests = dgcoll.FilterList(nc.Requests, func(rmd *RequestModelData) bool {
				return dgcoll.AnyMatch(entireModel.Interfaces, func(inter *InterfaceModelData) bool {
					service := filepath.Join(entireModel.Export.ServerOutput, "service", strcase.ToSnake(inter.Group)+"_service.go")
					serviceBytes, _ := os.ReadFile(service)
					serviceString := string(serviceBytes)

					return strings.Contains(serviceString, fmt.Sprintf("converter.%sConverter.%s2Entity(", c.DbTableUpperCamel, rmd.UpperCamelName)) ||
						strings.Contains(serviceString, fmt.Sprintf("converter.%sConverter.%s2Param(", c.DbTableUpperCamel, rmd.UpperCamelName)) ||
						strings.Contains(serviceString, fmt.Sprintf("converter.%sConverter.FillEntityWith%s(", c.DbTableUpperCamel, rmd.UpperCamelName))
				})
			})

			nc.Responses = dgcoll.FilterList(nc.Responses, func(rmd *ResponseModelData) bool {
				return !strings.Contains(fileString, fmt.Sprintf("converter.%sConverter.Entity2%s(", c.DbTableLowerCamel, rmd.UpperCamelName))
			})
			nc.Responses = dgcoll.FilterList(nc.Responses, func(rmd *ResponseModelData) bool {
				return dgcoll.AnyMatch(entireModel.Interfaces, func(inter *InterfaceModelData) bool {
					service := filepath.Join(entireModel.Export.ServerOutput, "service", strcase.ToSnake(inter.Group)+"_service.go")
					serviceBytes, _ := os.ReadFile(service)
					serviceString := string(serviceBytes)

					return strings.Contains(serviceString, fmt.Sprintf("converter.%sConverter.Entity2%s(", c.DbTableUpperCamel, rmd.UpperCamelName))
				})
			})

			if len(nc.Requests) == 0 && len(nc.Responses) == 0 {
				return nil
			}

			err = parseAppendFile(converter, "converter", _server.ConverterAppendTpl, nc)
			if err != nil {
				return err
			}
		} else {
			err := parseNewFile(converter, "converter", _server.ConverterTpl, c)
			if err != nil {
				return err
			}
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
			fileBytes, err := os.ReadFile(service)
			if err != nil {
				return err
			}
			fileString := string(fileBytes)

			newInter, err := utils.ConvertToNewBeanByJson[InterfaceModelData](inter)
			if err != nil {
				return err
			}
			newInter.Models = dgcoll.FilterList(newInter.Models, func(interfaceModel *InterfaceModel) bool {
				return !strings.Contains(fileString, fmt.Sprintf("%sService) %s(", inter.GroupLowerCamel, interfaceModel.MethodNameExp))
			})

			if len(newInter.Models) == 0 {
				return nil
			}

			err = parseAppendFile(service, "service", _server.ServiceAppendTpl, newInter)
			if err != nil {
				return err
			}
		} else {
			err := parseNewFile(service, "service", _server.ServiceTpl, inter)
			if err != nil {
				return err
			}
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

	for _, inter := range entireModel.Interfaces {
		handler := filepath.Join(handlerDir, strcase.ToSnake(inter.Group)+"_handler.go")
		if utils.ExistsFile(handler) {
			fileBytes, err := os.ReadFile(handler)
			if err != nil {
				return err
			}
			fileString := string(fileBytes)

			newInter, err := utils.ConvertToNewBeanByJson[InterfaceModelData](inter)
			if err != nil {
				return err
			}
			newInter.Models = dgcoll.FilterList(newInter.Models, func(interfaceModel *InterfaceModel) bool {
				return !strings.Contains(fileString, fmt.Sprintf("%sHandler) %s(", inter.GroupLowerCamel, interfaceModel.MethodNameExp))
			})

			if len(newInter.Models) == 0 {
				return nil
			}

			err = parseAppendFile(handler, "handler", _server.HandlerAppendTpl, newInter)
			if err != nil {
				return err
			}
		} else {
			err := parseNewFile(handler, "handler", _server.HandlerTpl, inter)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *serverParser) parseRouter(entireModel *EntireModel) error {
	if len(entireModel.Interfaces) == 0 {
		return nil
	}

	routerDir := filepath.Join(entireModel.Export.ServerOutput, "router")
	_ = os.MkdirAll(routerDir, fs.ModeDir|fs.ModePerm)

	for _, inter := range entireModel.Interfaces {
		router := filepath.Join(routerDir, strcase.ToSnake(inter.Group)+"_router.go")

		err := parseNewFile(router, "router", _server.RouterTpl, inter)
		if err != nil {
			return err
		}
	}

	return nil
}
