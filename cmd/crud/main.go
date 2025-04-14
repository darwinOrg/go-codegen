package main

import (
	"fmt"
	_ "github.com/darwinOrg/daog-ext"
	"github.com/darwinOrg/go-codegen/dgen"
	dgcoll "github.com/darwinOrg/go-common/collection"
	"github.com/darwinOrg/go-common/model"
	"github.com/darwinOrg/go-common/utils"
	_ "github.com/darwinOrg/go-logger"
	_ "github.com/darwinOrg/go-web/wrapper"
	_ "github.com/gin-gonic/gin"
	_ "github.com/rolandhe/daog"
	_ "github.com/shopspring/decimal"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	inputFile := os.Args[1]
	data, err := os.ReadFile(inputFile)
	if err != nil {
		os.Exit(1)
	}
	sql := string(data)

	metas, err := dgen.BuildTableMetas(sql)
	if err != nil {
		os.Exit(1)
	}
	if len(metas) == 0 {
		os.Exit(1)
	}

	entireModel := &dgen.EntireModel{}

	for _, meta := range metas {
		enumMap := meta.EnumMap

		for column, keyValues := range enumMap {
			entireModel.Enums = append(entireModel.Enums, &dgen.EnumModelData{
				Name:     column.GoName,
				DataType: adjustDbType(column.DbType),
				Models: dgcoll.MapToList(keyValues, func(keyValue *model.KeyValuePair[string, string]) *dgen.EnumModel {
					return &dgen.EnumModel{
						Code:  column.GoName + "_" + keyValue.Key,
						Value: keyValue.Key,
						Name:  keyValue.Value,
					}
				}),
			})
		}

		entireModel.Requests = append(entireModel.Requests, &dgen.RequestModelData{
			Name: "Create" + meta.GoTable + "Req",
			Models: dgcoll.MapToList(meta.CreateColumns, func(column *dgen.Column) *dgen.RequestModel {
				md := &dgen.RequestModel{
					FieldName: column.GoName,
					DataType:  adjustDbType(column.DbType),
					Nullable:  column.IsNull,
					Remark:    column.Comment,
				}

				if column.HasEnum {
					md.EnumModel = meta.GoTable + column.GoName
				}

				return md
			}),
		})

		entireModel.Requests = append(entireModel.Requests, &dgen.RequestModelData{
			Name: "Modify" + meta.GoTable + "Req",
			Models: dgcoll.MapToList(meta.ModifyColumns, func(column *dgen.Column) *dgen.RequestModel {
				requestModel := &dgen.RequestModel{
					FieldName: column.GoName,
					DataType:  adjustDbType(column.DbType),
					Nullable:  column.IsNull,
					Remark:    column.Comment,
				}

				if column.HasEnum {
					requestModel.EnumModel = meta.GoTable + column.GoName
				}

				return requestModel
			}),
		})

		entireModel.Requests = append(entireModel.Requests, &dgen.RequestModelData{
			Name: "Query" + meta.GoTable + "Req",
			Models: dgcoll.FilterAndMapToList(meta.QueryColumns, func(column *dgen.Column) bool {
				return column.DbName != "id" && !column.IsNull
			}, func(column *dgen.Column) *dgen.RequestModel {
				requestModel := &dgen.RequestModel{
					FieldName: column.GoName,
					DataType:  adjustDbType(column.DbType),
					Nullable:  true,
					Remark:    column.Comment,
				}

				if column.HasEnum {
					requestModel.EnumModel = meta.GoTable + column.GoName
				}

				return requestModel
			}),
		})

		entireModel.Responses = append(entireModel.Responses, &dgen.ResponseModelData{
			Name: meta.GoTable + "ListResp",
			Models: dgcoll.FlatMapToList(meta.QueryColumns, func(column *dgen.Column) []*dgen.ResponseModel {
				requestModel := &dgen.ResponseModel{
					FieldName: column.GoName,
					DataType:  adjustDbType(column.DbType),
					Nullable:  column.IsNull,
					Remark:    column.Comment,
				}

				if !column.HasEnum {
					return []*dgen.ResponseModel{requestModel}
				}

				return []*dgen.ResponseModel{requestModel, {
					FieldName: column.GoName + "Name",
					DataType:  "string",
					Nullable:  column.IsNull,
					EnumModel: meta.GoTable + column.GoName,
					Remark:    column.EnumName + "名称",
				}}
			}),
		})

		entireModel.Responses = append(entireModel.Responses, &dgen.ResponseModelData{
			Name: meta.GoTable + "DetailResp",
			Models: dgcoll.FlatMapToList(meta.QueryColumns, func(column *dgen.Column) []*dgen.ResponseModel {
				requestModel := &dgen.ResponseModel{
					FieldName: column.GoName,
					DataType:  adjustDbType(column.DbType),
					Nullable:  column.IsNull,
					Remark:    column.Comment,
				}

				if !column.HasEnum {
					return []*dgen.ResponseModel{requestModel}
				}

				return []*dgen.ResponseModel{requestModel, {
					FieldName: column.GoName + "Name",
					DataType:  "string",
					EnumModel: meta.GoTable + column.GoName,
					Remark:    column.EnumName + "名称",
				}}
			}),
		})

		entireModel.Interfaces = append(entireModel.Interfaces, &dgen.InterfaceModelData{
			Group:       meta.GoTable,
			RoutePrefix: "/api/" + meta.KebabName + "/v1",
			Models: []*dgen.InterfaceModel{
				{
					InterfaceType:    dgen.InterfaceTypeCreate,
					MethodType:       "Post",
					RelativePath:     "create",
					MethodName:       "create",
					DbModelName:      meta.TableName,
					RequestModelName: "Create" + meta.GoTable + "Req",
					LogLevel:         dgen.LogLevelAll,
					Remark:           meta.TableComment + " - " + dgen.InterfaceTypeCreate,
				},
				{
					InterfaceType:    dgen.InterfaceTypeModify,
					MethodType:       "Post",
					RelativePath:     "modify",
					MethodName:       "modify",
					DbModelName:      meta.TableName,
					RequestModelName: "Modify" + meta.GoTable + "Req",
					LogLevel:         dgen.LogLevelAll,
					Remark:           meta.TableComment + " - " + dgen.InterfaceTypeModify,
				},
				{
					InterfaceType:    dgen.InterfaceTypeDelete,
					MethodType:       "Post",
					RelativePath:     "delete",
					MethodName:       "delete",
					DbModelName:      meta.TableName,
					RequestModelName: dgen.RequestModelId,
					LogLevel:         dgen.LogLevelAll,
					Remark:           meta.TableComment + " - " + dgen.InterfaceTypeDelete,
				},
				{
					InterfaceType:     dgen.InterfaceTypePage,
					MethodType:        "Post",
					RelativePath:      "page",
					MethodName:        "page",
					DbModelName:       meta.TableName,
					RequestModelName:  "Query" + meta.GoTable + "Req",
					ResponseModelName: meta.GoTable + "ListResp",
					LogLevel:          dgen.LogLevelParam,
					Remark:            meta.TableComment + " - " + dgen.InterfaceTypePage,
				},
				{
					InterfaceType:     dgen.InterfaceTypeList,
					MethodType:        "Post",
					RelativePath:      "list",
					MethodName:        "list",
					DbModelName:       meta.TableName,
					RequestModelName:  "Query" + meta.GoTable + "Req",
					ResponseModelName: meta.GoTable + "ListResp",
					LogLevel:          dgen.LogLevelParam,
					Remark:            meta.TableComment + " - " + dgen.InterfaceTypeList,
				},
				{
					InterfaceType:     dgen.InterfaceTypeDetail,
					MethodType:        "Get",
					RelativePath:      "detail",
					MethodName:        "detail",
					DbModelName:       meta.TableName,
					RequestModelName:  dgen.RequestModelId,
					ResponseModelName: meta.GoTable + "DetailResp",
					LogLevel:          dgen.LogLevelParam,
					Remark:            meta.TableComment + " - " + dgen.InterfaceTypeDetail,
				},
			},
		})

	}

	outputJson := utils.MustConvertBeanToJsonString(entireModel)
	fmt.Print(outputJson)
}

func adjustDbType(dbType string) string {
	if strings.HasPrefix(dbType, "ttypes") {
		return "string"
	}

	return dbType
}
