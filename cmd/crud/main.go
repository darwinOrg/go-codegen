package main

import (
	"dgen/internal"
	"flag"
	"fmt"
	_ "github.com/darwinOrg/daog-ext"
	dgcoll "github.com/darwinOrg/go-common/collection"
	"github.com/darwinOrg/go-common/model"
	"github.com/darwinOrg/go-common/utils"
	_ "github.com/darwinOrg/go-logger"
	_ "github.com/darwinOrg/go-web/wrapper"
	_ "github.com/gin-gonic/gin"
	_ "github.com/rolandhe/daog"
	_ "github.com/shopspring/decimal"
	"os"
)

var (
	inputFile string
)

func init() {
	flag.StringVar(&inputFile, "i", "./scripts/test.sql", "the create tables file")
}

func main() {
	flag.Parse()

	if inputFile == "" {
		if len(os.Args) < 2 {
			os.Exit(1)
		}
		inputFile = os.Args[1]
	}
	data, err := os.ReadFile(inputFile)
	if err != nil {
		os.Exit(1)
	}
	sql := string(data)

	metas, err := internal.BuildTableMetas(sql)
	if err != nil {
		os.Exit(1)
	}
	if len(metas) == 0 {
		os.Exit(1)
	}

	entireModel := &internal.EntireModel{}

	for _, meta := range metas {
		enumMap := meta.EnumMap

		for column, keyValues := range enumMap {
			entireModel.Enums = append(entireModel.Enums, &internal.EnumModelData{
				Name:     meta.GoTable + column.GoName,
				DataType: column.DbType,
				Models: dgcoll.MapToList(keyValues, func(keyValue *model.KeyValuePair[string, string]) *internal.EnumModel {
					return &internal.EnumModel{
						Code:  meta.GoTable + column.GoName + "_" + keyValue.Key,
						Value: keyValue.Key,
						Name:  keyValue.Value,
					}
				}),
			})
		}

		entireModel.Requests = append(entireModel.Requests, &internal.RequestModelData{
			Name: "Create" + meta.GoTable + "Req",
			Models: dgcoll.MapToList(meta.CreateColumns, func(column *internal.Column) *internal.RequestModel {
				model := &internal.RequestModel{
					FieldName: column.GoName,
					DataType:  column.DbType,
					Nullable:  column.IsNull,
					Remark:    column.Comment,
				}

				if column.HasEnum {
					model.EnumModel = meta.GoTable + column.GoName
				}

				return model
			}),
		})

		entireModel.Requests = append(entireModel.Requests, &internal.RequestModelData{
			Name: "Modify" + meta.GoTable + "Req",
			Models: dgcoll.MapToList(meta.ModifyColumns, func(column *internal.Column) *internal.RequestModel {
				model := &internal.RequestModel{
					FieldName: column.GoName,
					DataType:  column.DbType,
					Nullable:  column.IsNull,
					Remark:    column.Comment,
				}

				if column.HasEnum {
					model.EnumModel = meta.GoTable + column.GoName
				}

				return model
			}),
		})

		entireModel.Requests = append(entireModel.Requests, &internal.RequestModelData{
			Name: "Query" + meta.GoTable + "Req",
			Models: dgcoll.MapToList(meta.QueryColumns, func(column *internal.Column) *internal.RequestModel {
				return &internal.RequestModel{
					FieldName: column.GoName,
					DataType:  column.DbType,
					Remark:    column.Comment,
				}
			}),
		})

		entireModel.Responses = append(entireModel.Responses, &internal.ResponseModelData{
			Name: meta.GoTable + "ListResp",
			Models: dgcoll.FlatMapToList(meta.QueryColumns, func(column *internal.Column) []*internal.ResponseModel {
				model := &internal.ResponseModel{
					FieldName: column.GoName,
					DataType:  column.DbType,
					Remark:    column.Comment,
				}

				if !column.HasEnum {
					return []*internal.ResponseModel{model}
				}

				return []*internal.ResponseModel{model, {
					FieldName: column.GoName + "Name",
					DataType:  "string",
					Remark:    column.EnumName + "名称",
				}}
			}),
		})

		entireModel.Responses = append(entireModel.Responses, &internal.ResponseModelData{
			Name: meta.GoTable + "DetailResp",
			Models: dgcoll.FlatMapToList(meta.QueryColumns, func(column *internal.Column) []*internal.ResponseModel {
				model := &internal.ResponseModel{
					FieldName: column.GoName,
					DataType:  column.DbType,
					Remark:    column.Comment,
				}

				if !column.HasEnum {
					return []*internal.ResponseModel{model}
				}

				return []*internal.ResponseModel{model, {
					FieldName: column.GoName + "Name",
					DataType:  "string",
					Remark:    column.EnumName + "名称",
				}}
			}),
		})

		entireModel.Interfaces = append(entireModel.Interfaces, &internal.InterfaceModelData{
			Group:       meta.GoTable,
			RoutePrefix: "/api/" + meta.KebabName + "/v1",
			Models: []*internal.InterfaceModel{
				{
					InterfaceType:    "新建",
					RelativePath:     "/create",
					DbModelName:      meta.TableName,
					RequestModelName: "Create" + meta.GoTable + "Req",
					LogLevel:         "全部",
					Remark:           meta.TableComment + " - 新建",
				},
				{
					InterfaceType:    "修改",
					RelativePath:     "/modify",
					DbModelName:      meta.TableName,
					RequestModelName: "Modify" + meta.GoTable + "Req",
					LogLevel:         "全部",
					Remark:           meta.TableComment + " - 修改",
				},
				{
					InterfaceType:    "删除",
					RelativePath:     "/delete",
					DbModelName:      meta.TableName,
					RequestModelName: "Id",
					LogLevel:         "全部",
					Remark:           meta.TableComment + " - 删除",
				},
				{
					InterfaceType:     "分页",
					RelativePath:      "/page",
					DbModelName:       meta.TableName,
					RequestModelName:  "Query" + meta.GoTable + "Req",
					ResponseModelName: meta.GoTable + "ListResp",
					LogLevel:          "请求参数",
					Remark:            meta.TableComment + " - 分页",
				},
				{
					InterfaceType:     "列表",
					RelativePath:      "/list",
					DbModelName:       meta.TableName,
					RequestModelName:  "Query" + meta.GoTable + "Req",
					ResponseModelName: meta.GoTable + "ListResp",
					LogLevel:          "请求参数",
					Remark:            meta.TableComment + " - 列表",
				},
				{
					InterfaceType:     "详情",
					RelativePath:      "/detail",
					DbModelName:       meta.TableName,
					RequestModelName:  "Id",
					ResponseModelName: meta.GoTable + "DetailResp",
					LogLevel:          "请求参数",
					Remark:            meta.TableComment + " - 详情",
				},
			},
		})

	}

	outputJson := utils.MustConvertBeanToJsonString(entireModel)
	fmt.Print(outputJson)
}
