package dgen

import (
	dgcoll "github.com/darwinOrg/go-common/collection"
	"github.com/darwinOrg/go-common/utils"
	dgcfg "github.com/darwinOrg/go-config"
	"github.com/darwinOrg/go-swagger"
	"github.com/darwinOrg/go-web/wrapper"
	"github.com/go-openapi/spec"
	"os"
)

func SyncToApifoxDefault(entireModel *EntireModel) {
	dgcfg.LoadDotEnv(dgcfg.MustConfRoot())

	SyncToApifox(entireModel, os.Getenv("APIFOX_ACCESS_TOKEN"), os.Getenv("APIFOX_PROJECT_ID"))
}

func SyncToApifox(entireModel *EntireModel, accessToken, projectId string) {
	swaggerProps := BuildSwaggerProps(entireModel)
	swaggerJsonBytes := utils.MustConvertBeanToJsonStringPretty(swaggerProps)

	swagger.SyncSwaggerJsonBytesToApifox(&swagger.SyncToApifoxRequest{
		AccessToken:         accessToken,
		ProjectId:           projectId,
		ApiOverwriteMode:    swagger.ApiOverwriteModeMethodAndPath,
		SchemaOverwriteMode: swagger.SchemaOverwriteModeBoth,
		SyncApiFolder:       false,
		ImportBasePath:      false,
		ApiFolderPath:       entireModel.Export.ApifoxOutput,
	}, []byte(swaggerJsonBytes))
}

func BuildSwaggerProps(entireModel *EntireModel) spec.SwaggerProps {
	req := &swagger.ExportSwaggerRequest{
		Title:       "接口文档",
		Description: "接口描述",
		Version:     "v1.0.0",
	}

	req.RequestApis = dgcoll.FlatMapToList(entireModel.Interfaces, func(imd *InterfaceModelData) []*wrapper.RequestApi {
		return dgcoll.MapToList(imd.Models, func(im *InterfaceModel) *wrapper.RequestApi {
			return &wrapper.RequestApi{
				Method:         im.MethodType,
				BasePath:       imd.Group,
				RelativePath:   im.RelativePath,
				Remark:         im.Remark,
				RequestObject:  im.RequestModelObject,
				ResponseObject: im.ResponseModelObject,
			}
		})
	})

	return swagger.BuildSwaggerProps(req)
}
