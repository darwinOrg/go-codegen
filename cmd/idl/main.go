package main

import (
	"dgen/internal"
	"dgen/internal/idl/transfer"
	nvwa_enum "e.globalpand.cn/libs/nvwa-sdk/api/nvwa/enum"
	dgcoll "github.com/darwinOrg/go-common/collection"
	"github.com/darwinOrg/go-common/utils"
	"github.com/iancoleman/strcase"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	inputFile := os.Args[1]

	def, err := transfer.ParseIdl(inputFile)
	if err != nil {
		panic(err)
	}

	createRelatedKeywords := []string{"Save", "Create", "New", "Insert", "Add", "save", "create", "new", "insert", "add"}
	modifyRelatedKeywords := []string{"Modify", "Update", "Alter", "modify", "update", "alter"}
	deleteRelatedKeywords := []string{"Delete", "Remove", "Drop", "delete", "remove", "drop"}
	queryRelatedKeywords := []string{"Query", "Search", "List", "query", "search", "list"}
	detailRelatedKeywords := []string{"Get", "Detail", "get", "detail"}

	productMap := map[int]string{
		nvwa_enum.Product.RPA:                  "RPA机器人",
		nvwa_enum.Product.AsyncInterview:       "AI智能面试",
		nvwa_enum.Product.OnlineInterview:      "金牌面试官",
		nvwa_enum.Product.RequirementAssistant: "需求沟通助手",
		nvwa_enum.Product.RPA:                  "候选人沟通助手",
		nvwa_enum.Product.CandidateAssistant:   "人才库",
	}

	entireModel := &internal.EntireModel{}

	for _, serviceDefine := range def.Services {
		inter := &internal.InterfaceModelData{
			Group:        serviceDefine.Name,
			RoutePrefix:  serviceDefine.RootUrl,
			ExportClient: strings.Contains(serviceDefine.RootUrl, "/internal/"),
		}
		entireModel.Interfaces = append(entireModel.Interfaces, inter)

		if len(serviceDefine.Posts) > 0 {
			for _, postMethod := range serviceDefine.Posts {
				interfaceModel := &internal.InterfaceModel{
					MethodType:   "Post",
					MethodName:   postMethod.Name,
					RelativePath: postMethod.Url,
					NonLogin:     postMethod.NotLogin,
					Remark:       postMethod.Description,
				}
				inter.Models = append(inter.Models, interfaceModel)

				if strings.Contains(postMethod.Description, "@prod=") {
					prodIndex := strings.Index(postMethod.Description, "@prod=")
					interfaceModel.Remark = postMethod.Description[:prodIndex]
					strProducts := postMethod.Description[prodIndex+len("@prod="):]
					if strProducts != "" {
						products := dgcoll.SplitToIntsByComma[int](strProducts)
						interfaceModel.AllowProducts = dgcoll.MapToList(products, func(product int) string {
							return productMap[product]
						})
					}
				}

				if postMethod.IsPager {
					interfaceModel.InterfaceType = "分页"
					interfaceModel.LogLevel = "请求参数"
				} else if postMethod.IsList {
					interfaceModel.InterfaceType = "列表"
					interfaceModel.LogLevel = "请求参数"
				} else {
					for _, createRelatedKeyword := range createRelatedKeywords {
						if strings.Contains(postMethod.Name, createRelatedKeyword) {
							interfaceModel.InterfaceType = "新建"
							interfaceModel.LogLevel = "全部"
							break
						}
					}

					if interfaceModel.InterfaceType == "" {
						for _, modifyRelatedKeyword := range modifyRelatedKeywords {
							if strings.Contains(postMethod.Name, modifyRelatedKeyword) {
								interfaceModel.InterfaceType = "修改"
								interfaceModel.LogLevel = "全部"
								break
							}
						}
					}

					if interfaceModel.InterfaceType == "" {
						for _, deleteRelatedKeyword := range deleteRelatedKeywords {
							if strings.Contains(postMethod.Name, deleteRelatedKeyword) {
								interfaceModel.InterfaceType = "删除"
								interfaceModel.LogLevel = "全部"
								break
							}
						}
					}

					if interfaceModel.InterfaceType == "" {
						for _, queryRelatedKeyword := range queryRelatedKeywords {
							if strings.Contains(postMethod.Name, queryRelatedKeyword) {
								interfaceModel.InterfaceType = "列表"
								interfaceModel.LogLevel = "请求参数"
								break
							}
						}
					}

					if interfaceModel.InterfaceType == "" {
						for _, detailRelatedKeyword := range detailRelatedKeywords {
							if strings.Contains(postMethod.Name, detailRelatedKeyword) {
								interfaceModel.InterfaceType = "详情"
								interfaceModel.LogLevel = "请求参数"
								break
							}
						}
					}

					if interfaceModel.InterfaceType == "" {
						interfaceModel.InterfaceType = "修改"
						interfaceModel.LogLevel = "全部"
					}
				}

				interfaceModel.RequestModelName = postMethod.Params.StructName
				interfaceModel.ResponseModelName = postMethod.MethodReturnType.TypeName
			}
		}

		if len(serviceDefine.Gets) > 0 {
			for _, getMethod := range serviceDefine.Gets {
				interfaceModel := &internal.InterfaceModel{
					MethodType:   "Get",
					MethodName:   getMethod.Name,
					RelativePath: getMethod.Url,
					NonLogin:     getMethod.NotLogin,
					Remark:       getMethod.Description,
				}
				inter.Models = append(inter.Models, interfaceModel)

				if strings.Contains(getMethod.Description, "@prod=") {
					prodIndex := strings.Index(getMethod.Description, "@prod=")
					interfaceModel.Remark = getMethod.Description[:prodIndex]
					strProducts := getMethod.Description[prodIndex+len("@prod="):]
					if strProducts != "" {
						products := dgcoll.SplitToIntsByComma[int](strProducts)
						interfaceModel.AllowProducts = dgcoll.MapToList(products, func(product int) string {
							return productMap[product]
						})
					}
				}

				if getMethod.IsPager {
					interfaceModel.InterfaceType = "分页"
					interfaceModel.LogLevel = "请求参数"
				} else if getMethod.IsList {
					interfaceModel.InterfaceType = "列表"
					interfaceModel.LogLevel = "请求参数"
				} else {
					for _, createRelatedKeyword := range createRelatedKeywords {
						if strings.Contains(getMethod.Name, createRelatedKeyword) {
							interfaceModel.InterfaceType = "新建"
							interfaceModel.LogLevel = "全部"
							break
						}
					}

					if interfaceModel.InterfaceType == "" {
						for _, modifyRelatedKeyword := range modifyRelatedKeywords {
							if strings.Contains(getMethod.Name, modifyRelatedKeyword) {
								interfaceModel.InterfaceType = "修改"
								interfaceModel.LogLevel = "全部"
								break
							}
						}
					}

					if interfaceModel.InterfaceType == "" {
						for _, deleteRelatedKeyword := range deleteRelatedKeywords {
							if strings.Contains(getMethod.Name, deleteRelatedKeyword) {
								interfaceModel.InterfaceType = "删除"
								interfaceModel.LogLevel = "全部"
								break
							}
						}
					}

					if interfaceModel.InterfaceType == "" {
						for _, queryRelatedKeyword := range queryRelatedKeywords {
							if strings.Contains(getMethod.Name, queryRelatedKeyword) {
								interfaceModel.InterfaceType = "列表"
								interfaceModel.LogLevel = "请求参数"
								break
							}
						}
					}

					if interfaceModel.InterfaceType == "" {
						for _, detailRelatedKeyword := range detailRelatedKeywords {
							if strings.Contains(getMethod.Name, detailRelatedKeyword) {
								interfaceModel.InterfaceType = "详情"
								interfaceModel.LogLevel = "请求参数"
								break
							}
						}
					}

					if interfaceModel.InterfaceType == "" {
						interfaceModel.InterfaceType = "详情"
						interfaceModel.LogLevel = "请求参数"
					}
				}

				if getMethod.Params.StructName != "" {
					interfaceModel.RequestModelName = getMethod.Params.StructName
				} else if len(getMethod.Params.BasicParams) > 0 {
					interfaceModel.RequestModelName = strcase.ToCamel(getMethod.Name)
					sd := &transfer.StructDefine{
						Name: interfaceModel.RequestModelName,
					}
					def.Structs = append(def.Structs, sd)

					for _, param := range getMethod.Params.BasicParams {
						field := &transfer.Field{
							Name: param.ParamName,
							Tp: &transfer.FieldType{
								IsBasic:  true,
								IsList:   param.IsList,
								TypeName: param.TypeName,
							},
							Annotations: param.Annotations,
						}

						if param.IsList {
							field.Tp.ValueType = &transfer.FieldType{
								IsBasic:  true,
								TypeName: param.TypeName,
							}
						}

						if dgcoll.NoneMatch(field.Annotations, func(annotation *transfer.Annotation) bool {
							return annotation.Key == "binding"
						}) {
							field.Annotations = append(field.Annotations, &transfer.Annotation{
								Key:   "binding",
								Value: "required",
							})
						}

						sd.Fields = append(sd.Fields, field)
					}
				}
				interfaceModel.ResponseModelName = getMethod.MethodReturnType.TypeName
			}
		}

		if len(inter.Models) > 0 {
			for _, interfaceModel := range inter.Models {
				if interfaceModel.RequestModelName != "" {
					convertRequest(def, entireModel, interfaceModel.RequestModelName)
				}

				if interfaceModel.ResponseModelName != "" {
					convertResponse(def, entireModel, interfaceModel.ResponseModelName)
				}
			}
		}
	}

	utils.WriteFileWithString("output/"+def.Namespace+".cui", utils.MustConvertBeanToJsonString(entireModel))
}

func convertRequest(def *transfer.Definition, entireModel *internal.EntireModel, requestModelName string) {
	if len(entireModel.Requests) > 0 {
		for _, request := range entireModel.Requests {
			if request.Name == requestModelName {
				return
			}
		}
	}

	for _, structDefine := range def.Structs {
		if structDefine.Name == requestModelName {
			request := &internal.RequestModelData{
				Name: structDefine.Name,
			}
			entireModel.Requests = append(entireModel.Requests, request)

			if structDefine.Extends != "page.PageParam" {
				request.ExtendName = structDefine.Extends
			}

			if len(structDefine.Fields) > 0 {
				for _, field := range structDefine.Fields {
					rm := &internal.RequestModel{
						FieldName: field.Name,
						IsArray:   field.Tp.IsList,
						Nullable:  true,
					}

					if field.Tp.IsList {
						rm.DataType = field.Tp.ValueType.TypeName
						rm.IsPointer = field.Tp.ValueType.IsStruct
					} else {
						rm.DataType = field.Tp.TypeName
						rm.IsPointer = field.Tp.IsStruct
					}

					if rm.IsPointer {
						convertRequest(def, entireModel, rm.DataType)
					}

					if len(field.Annotations) > 0 {
						for _, annotation := range field.Annotations {
							annotationValue := strings.TrimPrefix(annotation.Value, `"`)
							annotationValue = strings.TrimSuffix(annotationValue, `"`)

							if annotation.Key == "binding" {
								if strings.Contains(annotationValue, "required") {
									rm.Nullable = false
									rm.VerifyRules = strings.Trim(rm.VerifyRules, "required,")
									rm.VerifyRules = strings.Trim(rm.VerifyRules, "required")
								}
							} else if annotation.Key == "remark" {
								rm.Remark = annotationValue
							}
						}
					}

					request.Models = append(request.Models, rm)
				}
			}

			break
		}
	}
}

func convertResponse(def *transfer.Definition, entireModel *internal.EntireModel, responseModelName string) {
	if len(entireModel.Responses) > 0 {
		for _, response := range entireModel.Responses {
			if response.Name == responseModelName {
				return
			}
		}
	}

	for _, structDefine := range def.Structs {
		if structDefine.Name == responseModelName {
			response := &internal.ResponseModelData{
				Name:       structDefine.Name,
				ExtendName: structDefine.Extends,
			}
			entireModel.Responses = append(entireModel.Responses, response)

			if len(structDefine.Fields) > 0 {
				for _, field := range structDefine.Fields {
					rm := &internal.ResponseModel{
						FieldName: field.Name,
						IsArray:   field.Tp.IsList,
						Nullable:  true,
					}

					if field.Tp.IsList {
						rm.DataType = field.Tp.ValueType.TypeName
						rm.IsPointer = field.Tp.ValueType.IsStruct
					} else {
						rm.DataType = field.Tp.TypeName
						rm.IsPointer = field.Tp.IsStruct
					}

					if rm.IsPointer {
						convertResponse(def, entireModel, rm.DataType)
					}

					if len(field.Annotations) > 0 {
						for _, annotation := range field.Annotations {
							annotationValue := strings.TrimPrefix(annotation.Value, `"`)
							annotationValue = strings.TrimSuffix(annotationValue, `"`)

							if annotation.Key == "remark" {
								rm.Remark = annotationValue
							} else if annotation.Key == "appendUid" {
								rm.IsMediaUrl = true
							}
						}
					}

					response.Models = append(response.Models, rm)
				}
			}

			break
		}
	}
}
