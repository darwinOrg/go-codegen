package dgen

import (
	dgcoll "github.com/darwinOrg/go-common/collection"
	"github.com/darwinOrg/go-common/result"
	"github.com/darwinOrg/go-common/utils"
	dgcfg "github.com/darwinOrg/go-config"
	"github.com/darwinOrg/go-swagger"
	"github.com/darwinOrg/go-web/wrapper"
	"github.com/go-openapi/spec"
	"os"
)

type pageList struct {
	PageNo     int   `json:"pageNo" binding:"required" remark:"页码"`
	PageSize   int   `json:"pageSize" binding:"required" remark:"每页记录数"`
	TotalCount int   `json:"totalCount" remark:"总记录数"`
	TotalPages int   `json:"totalPages" remark:"总页数"`
	List       []any `json:"list" remark:"列表记录"`
}

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
			ra := &wrapper.RequestApi{
				Method:        im.MethodType,
				BasePath:      imd.RoutePrefix,
				RelativePath:  im.RelativePath,
				Remark:        im.Remark,
				RequestObject: im.RequestModelObject,
			}

			if im.ResponseApiObject != nil {
				ra.ResponseObject = im.ResponseApiObject
			} else if im.ResponseApiObject == nil && im.ResponseModelObject != nil {
				ra.ResponseObject = result.Success(&pageList{
					List: []any{im.ResponseModelObject},
				})
			}

			return ra
		})
	})

	return swagger.BuildSwaggerProps(req)
}
