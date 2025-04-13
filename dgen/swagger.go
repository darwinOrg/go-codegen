package dgen

import (
	"fmt"
	"github.com/darwinOrg/go-common/utils"
	"github.com/go-openapi/spec"
	"net/http"
	"strings"
)

const (
	contentTypeJson = "application/json"
)

func BuildSwaggerProps(entireModel *EntireModel) spec.SwaggerProps {
	return spec.SwaggerProps{
		Swagger:             "2.0",
		Definitions:         spec.Definitions{},
		SecurityDefinitions: spec.SecurityDefinitions{},
		Info: &spec.Info{
			InfoProps: spec.InfoProps{
				Title:       "接口文档",
				Description: "接口描述",
				Version:     "v1.0.0",
			},
		},
		Paths: buildPaths(entireModel),
	}
}

func buildPaths(entireModel *EntireModel) *spec.Paths {
	paths := map[string]spec.PathItem{}

	for _, g := range entireModel.Interfaces {
		for _, m := range g.Models {
			url := fmt.Sprintf("%s/%s", g.RoutePrefix, m.RelativePath)
			url = strings.ReplaceAll(url, "//", "/")

			operation := &spec.Operation{
				OperationProps: spec.OperationProps{
					Summary:     m.Remark,
					Description: m.Remark,
					ID:          g.GroupUpperCamel + "_" + m.MethodNameExp + "_",
					Consumes:    []string{contentTypeJson},
					Produces:    []string{contentTypeJson},
					Parameters: utils.IfReturn(strings.EqualFold(m.MethodType, http.MethodGet),
						buildGetParameters(m), buildPostParameters(m)),
					Responses: buildResponses(m),
				},
			}

			itemProps := spec.PathItemProps{}
			if strings.EqualFold(m.MethodType, http.MethodGet) {
				itemProps.Get = operation
			} else {
				itemProps.Post = operation
			}

			paths[url] = spec.PathItem{
				PathItemProps: itemProps,
			}
		}
	}

	return &spec.Paths{
		Paths: paths,
	}
}

func buildGetParameters(interfaceModel *InterfaceModel) []spec.Parameter {
	var parameters []spec.Parameter

	if interfaceModel.InterfaceType == InterfaceTypePage {
		pageNoParam := *spec.QueryParam("pageNo")
		pageNoParam.Required = true
		pageNoParam.Description = "页码"

		pageSizeParam := *spec.QueryParam("pageSize")
		pageSizeParam.Required = true
		pageSizeParam.Description = "每页记录数"

		parameters = append(parameters, pageNoParam, pageSizeParam)
	}

	if interfaceModel.RequestModelName == "Id" {
		p := *spec.QueryParam("id")
		p.Required = true
		p.Description = "id"

		parameters = append(parameters, p)
		return parameters
	}

	if interfaceModel.RequestModelData == nil || len(interfaceModel.RequestModelData.Models) == 0 {
		return parameters
	}

	for _, requestModel := range interfaceModel.RequestModelData.Models {
		p := *spec.QueryParam(requestModel.LowerCamelName)
		if requestModel.IsArray {
			p.Type = "array"
		} else {
			switch requestModel.DataType {
			case "int", "int64", "uint64", "int8", "uint8", "int16", "uint16", "int32", "uint32":
				p.Type = "integer"
			case "float32", "float64", "decimal.Decimal":
				p.Type = "number"
			default:
				p.Type = "string"
			}
		}

		p.Required = !requestModel.Nullable
		p.Description = requestModel.Remark

		parameters = append(parameters, p)
	}

	return parameters
}

func buildPostParameters(interfaceModel *InterfaceModel) []spec.Parameter {
	schema := createPostRequestSchemaForInterface(interfaceModel)
	bodyParam := *spec.BodyParam("body", schema)
	bodyParam.Required = true
	return []spec.Parameter{bodyParam}
}

func createPostRequestSchemaForInterface(interfaceModel *InterfaceModel) *spec.Schema {
	schema := &spec.Schema{}
	schema.Type = []string{"object"}
	schema.Properties = make(map[string]spec.Schema)

	if interfaceModel.InterfaceType == InterfaceTypePage {
		pageNoProperty := &spec.Schema{}
		pageNoProperty.Type = []string{"integer"}
		pageNoProperty.Description = "页码"
		schema.Properties["pageNo"] = *pageNoProperty

		pageSizeProperty := &spec.Schema{}
		pageSizeProperty.Type = []string{"integer"}
		pageSizeProperty.Description = "每页记录数"
		schema.Properties["pagSize"] = *pageSizeProperty

		schema.Required = append(schema.Required, "pageNo", "pageSize")
	}

	if interfaceModel.RequestModelName == "Id" {
		idProperty := &spec.Schema{}
		idProperty.Type = []string{"integer"}
		idProperty.Description = "id"
		schema.Properties["id"] = *idProperty
		schema.Required = append(schema.Required, "id")

		return schema
	}

	if interfaceModel.RequestModelName == "" || interfaceModel.RequestModelData == nil {
		return schema
	}

	for _, requestModelData := range []*RequestModelData{interfaceModel.RequestModelData, interfaceModel.RequestModelData.ExtendRequestModelData} {
		if requestModelData == nil {
			continue
		}

		fillPropertyForRequest(schema, requestModelData, 0)
	}

	return schema
}

func fillPropertyForRequest(schema *spec.Schema, requestModelData *RequestModelData, depth int) {
	// 限制递归深度最大为8
	if depth > 8 {
		return
	}

	for _, requestModel := range requestModelData.Models {
		var property *spec.Schema
		modelSchema := createSchemaForRequest(requestModel, depth)

		if requestModel.IsArray {
			property = &spec.Schema{}
			property.Type = []string{"array"}
			property.Items = &spec.SchemaOrArray{Schema: modelSchema}
			property.Description = requestModel.Remark
		} else {
			property = modelSchema
		}

		schema.Properties[requestModel.LowerCamelName] = *property
		if !requestModel.Nullable {
			schema.Required = append(schema.Required, requestModel.LowerCamelName)
		}
	}
}

func createSchemaForRequest(requestModel *RequestModel, depth int) *spec.Schema {
	schema := &spec.Schema{}
	schema.Description = requestModel.Remark

	switch requestModel.DataType {
	case "string":
		schema.Type = []string{"string"}
	case "int", "int64", "uint64", "int8", "uint8", "int16", "uint16", "int32", "uint32":
		schema.Type = []string{"integer"}
	case "float32", "float64", "decimal.Decimal":
		schema.Type = []string{"number"}
	case "bool":
		schema.Type = []string{"boolean"}
	case "any":
	default:
		schema.Type = []string{"object"}
		schema.Properties = make(map[string]spec.Schema)
		schema.Required = make([]string, 0)

		if requestModel.RequestModelData != nil {
			fillPropertyForRequest(schema, requestModel.RequestModelData, depth+1)
		}
	}

	return schema
}

func buildResponses(interfaceModel *InterfaceModel) *spec.Responses {
	return &spec.Responses{
		ResponsesProps: spec.ResponsesProps{
			StatusCodeResponses: map[int]spec.Response{
				http.StatusOK: {
					ResponseProps: spec.ResponseProps{
						Description: "成功",
						Schema:      createResponseSchemaForInterface(interfaceModel),
					},
				},
			},
		},
	}
}

func createResponseSchemaForInterface(interfaceModel *InterfaceModel) *spec.Schema {
	successProperty := &spec.Schema{}
	successProperty.Type = []string{"bool"}
	successProperty.Description = "是否成功"

	codeProperty := newIntegerSchema("消息编码")

	messageProperty := &spec.Schema{}
	messageProperty.Type = []string{"string"}
	messageProperty.Description = "消息内容"

	schema := &spec.Schema{}
	schema.Type = []string{"object"}
	schema.Required = []string{"success", "code"}
	schema.Properties = map[string]spec.Schema{
		"success": *successProperty,
		"code":    *codeProperty,
		"message": *messageProperty,
	}

	if interfaceModel.ResponseModelName == "" || interfaceModel.ResponseModelData == nil {
		return schema
	}

	var dataProperty *spec.Schema
	modelSchema := createResponseModelSchema(interfaceModel)

	if interfaceModel.InterfaceType == InterfaceTypePage {
		dataProperty = &spec.Schema{}
		dataProperty.Type = []string{"object"}

		pageNoProperty := newIntegerSchema("页码")
		pageSizeProperty := newIntegerSchema("每页记录数")
		totalCountProperty := newIntegerSchema("总记录数")
		totalPagesProperty := newIntegerSchema("总页数")

		listProperty := &spec.Schema{}
		listProperty.Type = []string{"array"}
		listProperty.Description = "列表数据"
		listProperty.Items = &spec.SchemaOrArray{Schema: modelSchema}

		dataProperty.Properties = map[string]spec.Schema{
			"pageNo":     *pageNoProperty,
			"pageSize":   *pageSizeProperty,
			"totalCount": *totalCountProperty,
			"totalPages": *totalPagesProperty,
			"list":       *listProperty,
		}

		dataProperty.Required = []string{"pageNo", "pageSize", "totalCount", "totalPages"}
	} else if interfaceModel.InterfaceType == InterfaceTypeList {
		dataProperty = &spec.Schema{}
		dataProperty.Type = []string{"array"}
		dataProperty.Items = &spec.SchemaOrArray{Schema: modelSchema}
	} else {
		dataProperty = modelSchema
	}

	dataProperty.Description = "数据"
	schema.Properties["data"] = *dataProperty
	return schema
}

func createResponseModelSchema(interfaceModel *InterfaceModel) *spec.Schema {
	schema := &spec.Schema{}
	schema.Type = []string{"object"}
	schema.Properties = make(map[string]spec.Schema)
	schema.Required = make([]string, 0)
	depth := 0

	for _, responseModelData := range []*ResponseModelData{interfaceModel.ResponseModelData, interfaceModel.ResponseModelData.ExtendResponseModelData} {
		if responseModelData == nil {
			continue
		}

		fillPropertyForResponse(schema, responseModelData, depth)
	}

	return schema
}

func fillPropertyForResponse(schema *spec.Schema, responseModelData *ResponseModelData, depth int) {
	// 限制递归深度最大为8
	if depth > 8 {
		return
	}

	for _, responseModel := range responseModelData.Models {
		var property *spec.Schema
		modelSchema := createSchemaForResponse(responseModel, depth)

		if responseModel.IsArray {
			property = &spec.Schema{}
			property.Type = []string{"array"}
			property.Items = &spec.SchemaOrArray{Schema: modelSchema}
		} else {
			property = modelSchema
		}

		schema.Properties[responseModel.LowerCamelName] = *property
		if !responseModel.Nullable {
			schema.Required = append(schema.Required, responseModel.LowerCamelName)
		}
	}
}

func createSchemaForResponse(responseModel *ResponseModel, depth int) *spec.Schema {
	schema := &spec.Schema{}

	if responseModel.EnumTitle != "" {
		schema.Title = responseModel.EnumTitle
	}

	if responseModel.EnumRemark != "" {
		schema.Description = responseModel.EnumRemark
	} else {
		schema.Description = responseModel.Remark
	}

	switch responseModel.DataType {
	case "string":
		schema.Type = []string{"string"}
	case "int", "int64", "uint64", "int8", "uint8", "int16", "uint16", "int32", "uint32":
		schema.Type = []string{"integer"}
	case "float32", "float64", "decimal.Decimal":
		schema.Type = []string{"number"}
	case "bool":
		schema.Type = []string{"boolean"}
	case "any":
	default:
		schema.Type = []string{"object"}
		schema.Properties = make(map[string]spec.Schema)
		schema.Required = make([]string, 0)

		if responseModel.ResponseModelData != nil {
			fillPropertyForResponse(schema, responseModel.ResponseModelData, depth)
		}
	}

	return schema
}

func newIntegerSchema(desc string) *spec.Schema {
	codeProperty := &spec.Schema{}
	codeProperty.Type = []string{"integer"}
	codeProperty.Description = desc
	return codeProperty
}
