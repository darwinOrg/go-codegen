package dgen

import (
	"fmt"
	"github.com/darwinOrg/go-codegen/pkg"
	dgcoll "github.com/darwinOrg/go-common/collection"
	"github.com/darwinOrg/go-common/utils"
	"github.com/iancoleman/strcase"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var (
	productMap = map[string]string{
		"RPA机器人":  "nvwa_enum.Product.RPA",
		"AI智能面试":  "nvwa_enum.Product.AsyncInterview",
		"金牌面试官":   "nvwa_enum.Product.OnlineInterview",
		"需求沟通助手":  "nvwa_enum.Product.RequirementAssistant",
		"候选人沟通助手": "nvwa_enum.Product.RPA",
		"人才库":     "nvwa_enum.Product.CandidateAssistant",
	}
	logLevelMap = map[string]string{
		LogLevelAll:    "wrapper.LOG_LEVEL_ALL",
		LogLevelParam:  "wrapper.LOG_LEVEL_PARAM",
		LogLevelReturn: "wrapper.LOG_LEVEL_RETURN",
		LogLevelNone:   "wrapper.LOG_LEVEL_NONE",
	}
)

type EntireModel struct {
	Dbs        []*DbModelData        `json:"dbs,omitempty"`
	Enums      []*EnumModelData      `json:"enums,omitempty"`
	Requests   []*RequestModelData   `json:"requests,omitempty"`
	Responses  []*ResponseModelData  `json:"responses,omitempty"`
	Converters []*ConverterData      `json:"converters,omitempty"`
	Interfaces []*InterfaceModelData `json:"interfaces,omitempty"`
	Export     *ExportConfigData     `json:"export,omitempty"`

	PackagePrefix  string `json:"packagePrefix,omitempty"`
	HasDecimal     bool   `json:"hasDecimal,omitempty"`
	HasPage        bool   `json:"hasPage,omitempty"`
	HasQuery       bool   `json:"hasQuery,omitempty"`
	HasId          bool   `json:"hasId,omitempty"`
	HasModel       bool   `json:"hasModel,omitempty"`
	HasProducts    bool   `json:"hasProducts,omitempty"`
	UpperCamelName string `json:"upperCamelName,omitempty"`
}

type DbModel struct {
	DbFieldName    string `json:"dbFieldName,omitempty"`
	ModelFieldName string `json:"modelFieldName,omitempty"`
	DbDataType     string `json:"dbDataType,omitempty"`
	ModelDataType  string `json:"modelDataType,omitempty"`
	Nullable       bool   `json:"nullable"`
	DefaultValue   string `json:"defaultValue,omitempty"`
	Enumerable     bool   `json:"enumerable,omitempty"`
	Remark         string `json:"remark,omitempty"`
}

type DbModelData struct {
	Sql    string     `json:"sql"`
	Name   string     `json:"name"`
	Models []*DbModel `json:"models"`
}

type EnumModel struct {
	Code  string `json:"code,omitempty"`
	Value string `json:"value,omitempty"`
	Name  string `json:"name,omitempty"`
}

type EnumModelData struct {
	Name     string       `json:"name,omitempty"`
	Models   []*EnumModel `json:"models,omitempty"`
	DataType string       `json:"dataType,omitempty"`

	UpperCamelName string `json:"upperCamelName,omitempty"`
}

type RequestModel struct {
	FieldName   string `json:"fieldName,omitempty"`
	DataType    string `json:"dataType,omitempty"`
	IsPointer   bool   `json:"isPointer,omitempty"`
	IsArray     bool   `json:"isArray,omitempty"`
	Nullable    bool   `json:"nullable"`
	VerifyRules string `json:"verifyRules,omitempty"`
	EnumModel   string `json:"enumModel,omitempty"`
	Remark      string `json:"remark,omitempty"`

	UpperCamelName   string            `json:"upperCamelName,omitempty"`
	LowerCamelName   string            `json:"lowerCamelName,omitempty"`
	NullableString   bool              `json:"nullableString,omitempty"`
	RequestModelData *RequestModelData `json:"requestModelData,omitempty"`
}

type RequestModelData struct {
	Name       string          `json:"name,omitempty"`
	Models     []*RequestModel `json:"models,omitempty"`
	ExtendName string          `json:"extendName,omitempty"`

	UpperCamelName         string            `json:"upperCamelName,omitempty"`
	IsPage                 bool              `json:"isPage,omitempty"`
	IsPageOrList           bool              `json:"isPageOrList,omitempty"`
	InterfaceType          string            `json:"interfaceType,omitempty"`
	ExtendRequestModelData *RequestModelData `json:"extendRequestModelData,omitempty"`
}

type ResponseModel struct {
	FieldName  string `json:"fieldName,omitempty"`
	DataType   string `json:"dataType,omitempty"`
	IsPointer  bool   `json:"isPointer,omitempty"`
	IsArray    bool   `json:"isArray,omitempty"`
	Nullable   bool   `json:"nullable"`
	IsMediaUrl bool   `json:"isMediaUrl,omitempty"`
	EnumModel  string `json:"enumModel,omitempty"`
	Remark     string `json:"remark,omitempty"`

	UpperCamelName    string             `json:"upperCamelName,omitempty"`
	LowerCamelName    string             `json:"lowerCamelName,omitempty"`
	EnumTitle         string             `json:"enumTitle,omitempty"`
	EnumRemark        string             `json:"enumRemark,omitempty"`
	EnumFieldName     string             `json:"enumFieldName,omitempty"`
	NullableString    bool               `json:"nullableString,omitempty"`
	ResponseModelData *ResponseModelData `json:"responseModelData,omitempty"`
}

type ResponseModelData struct {
	Name       string           `json:"name,omitempty"`
	Models     []*ResponseModel `json:"models,omitempty"`
	ExtendName string           `json:"extendName,omitempty"`

	UpperCamelName          string             `json:"upperCamelName,omitempty"`
	InterfaceType           string             `json:"interfaceType,omitempty"`
	ExtendResponseModelData *ResponseModelData `json:"extendResponseModelData,omitempty"`
}

type InterfaceModel struct {
	InterfaceType     string   `json:"interfaceType,omitempty"`
	MethodType        string   `json:"methodType,omitempty"`
	RelativePath      string   `json:"relativePath,omitempty"`
	MethodName        string   `json:"methodName,omitempty"`
	DbModelName       string   `json:"dbModelName,omitempty"`
	RequestModelName  string   `json:"requestModelName,omitempty"`
	ResponseModelName string   `json:"responseModelName,omitempty"`
	NonLogin          bool     `json:"nonLogin,omitempty"`
	AllowRoles        []string `json:"allowRoles,omitempty"`
	AllowProducts     []string `json:"allowProducts,omitempty"`
	LogLevel          string   `json:"logLevel,omitempty"`
	NotLogSQL         bool     `json:"notLogSQL,omitempty"`
	Remark            string   `json:"remark,omitempty"`

	DbTableUpperCamel       string             `json:"dbTableUpperCamel,omitempty"`
	DbTableLowerCamel       string             `json:"dbTableLowerCamel,omitempty"`
	RequestModelData        *RequestModelData  `json:"requestModelData,omitempty"`
	RequestModelNameExp     string             `json:"requestModelNameExp,omitempty"`
	ResponseModelData       *ResponseModelData `json:"responseModelData,omitempty"`
	ResponseModelNameExp    string             `json:"responseModelNameExp,omitempty"`
	ResponseModelHasPointer bool               `json:"responseModelHasPointer,omitempty"`
	AllowProductsExp        string             `json:"allowProductsExp,omitempty"`
	LogLevelExp             string             `json:"logLevelExp,omitempty"`
	MethodNameExp           string             `json:"methodNameExp,omitempty"`
}

type InterfaceModelData struct {
	Group        string            `json:"group,omitempty"`
	RoutePrefix  string            `json:"routePrefix,omitempty"`
	ExportClient bool              `json:"exportClient"`
	Models       []*InterfaceModel `json:"models,omitempty"`

	PackagePrefix   string `json:"packagePrefix,omitempty"`
	GroupUpperCamel string `json:"groupUpperCamel,omitempty"`
	GroupLowerCamel string `json:"groupLowerCamel,omitempty"`
	GroupSnake      string `json:"groupSnake,omitempty"`
	HasPage         bool   `json:"hasPage,omitempty"`
	HasQuery        bool   `json:"hasQuery,omitempty"`
	HasId           bool   `json:"hasId,omitempty"`
	HasModel        bool   `json:"hasModel,omitempty"`
	HasProducts     bool   `json:"hasProducts,omitempty"`
	HasDbTable      bool   `json:"hasDbTable,omitempty"`
}

type ExportConfigData struct {
	FilePrefix          string `json:"filePrefix,omitempty"`
	ServerOutput        string `json:"serverOutput,omitempty"`
	ServerPackagePrefix string `json:"serverPackagePrefix,omitempty"`
	ClientOutput        string `json:"clientOutput,omitempty"`
	ClientPackagePrefix string `json:"clientPackagePrefix,omitempty"`
	ApifoxOutput        string `json:"apifoxOutput,omitempty"`
}

type ConverterData struct {
	PackagePrefix     string               `json:"packagePrefix,omitempty"`
	DbTableUpperCamel string               `json:"dbTableUpperCamel,omitempty"`
	DbTableLowerCamel string               `json:"dbTableLowerCamel,omitempty"`
	Requests          []*RequestModelData  `json:"requests,omitempty"`
	Responses         []*ResponseModelData `json:"responses,omitempty"`
	HasEnum           bool                 `json:"hasEnum,omitempty"`
}

func InitEntireModel() *EntireModel {
	if len(os.Args) < 2 {
		fmt.Println("Please input code generation design file")
		os.Exit(1)
	}
	inputFile := os.Args[len(os.Args)-1]
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Read the code generation design file error: %v", err)
		os.Exit(1)
	}

	entireModel, err := utils.ConvertJsonBytesToBean[EntireModel](data)
	if err != nil {
		fmt.Printf("Parse the code generation design file error: %v", err)
		os.Exit(1)
	}

	filenameWithExt := filepath.Base(inputFile)
	filename := strings.TrimSuffix(filenameWithExt, filepath.Ext(inputFile))
	entireModel.UpperCamelName = strcase.ToCamel(filename)

	entireModel.Export.FilePrefix = strcase.ToSnake(filename)

	if !strings.HasPrefix(entireModel.Export.ServerOutput, "/") {
		entireModel.Export.ServerOutput = filepath.Join(filepath.Dir(inputFile), entireModel.Export.ServerOutput)
	}
	if entireModel.Export.ServerOutput == "" {
		entireModel.Export.ServerOutput = "."
	}

	if !strings.HasPrefix(entireModel.Export.ClientOutput, "/") {
		entireModel.Export.ClientOutput = filepath.Join(filepath.Dir(inputFile), entireModel.Export.ClientOutput)
	}
	if entireModel.Export.ClientOutput == "" {
		entireModel.Export.ClientOutput = "."
	}

	return entireModel
}

func (m *EntireModel) Fill(packagePrefix string) {
	if m.Enums == nil {
		m.Enums = []*EnumModelData{}
	}
	if m.Requests == nil {
		m.Requests = []*RequestModelData{}
	}
	if m.Responses == nil {
		m.Responses = []*ResponseModelData{}
	}
	if m.Interfaces == nil {
		m.Interfaces = []*InterfaceModelData{}
	}
	m.HasModel = len(m.Requests) > 0 || len(m.Responses) > 0

	m.FillEnums()
	m.FillRequests()
	m.FillResponses()
	m.FillInterfaces()
	m.FillConverters()
	m.FillPackagePrefix(packagePrefix)
}

func (m *EntireModel) FillEnums() {
	if len(m.Enums) == 0 {
		return
	}

	for _, enum := range m.Enums {
		enum.UpperCamelName = strcase.ToCamel(enum.Name)
	}
}

func (m *EntireModel) FillRequests() {
	if len(m.Requests) == 0 {
		return
	}

	for _, request := range m.Requests {
		request.UpperCamelName = strcase.ToCamel(request.Name)

		for _, model := range request.Models {
			model.UpperCamelName = strcase.ToCamel(model.FieldName)
			model.LowerCamelName = strcase.ToLowerCamel(model.FieldName)
			model.NullableString = model.Nullable && model.DataType == "string"

			if !model.Nullable {
				if model.VerifyRules != "" {
					model.VerifyRules = "required," + model.VerifyRules
				} else {
					model.VerifyRules = "required"
				}
			}

			if model.DataType == "decimal.Decimal" {
				m.HasDecimal = true
			}

			if model.DataType == "Id" {
				m.HasId = true
			}

			for _, req := range m.Requests {
				if req.Name == model.DataType {
					model.RequestModelData = request
				}
			}
		}

		if request.ExtendName != "" {
			for _, req := range m.Requests {
				if req.Name == request.ExtendName {
					req.ExtendRequestModelData = request
				}
			}
		}
	}
}

func (m *EntireModel) FillResponses() {
	if len(m.Responses) == 0 {
		return
	}

	for _, response := range m.Responses {
		response.UpperCamelName = strcase.ToCamel(response.Name)

		for _, model := range response.Models {
			model.UpperCamelName = strcase.ToCamel(model.FieldName)
			model.LowerCamelName = strcase.ToLowerCamel(model.FieldName)
			model.NullableString = model.Nullable && model.DataType == "string"

			columnComment := model.Remark
			columnComment = strings.ReplaceAll(columnComment, "（", "(")
			columnComment = strings.ReplaceAll(columnComment, "）", ")")
			columnComment = strings.ReplaceAll(columnComment, "，", ",")
			columnComment = strings.ReplaceAll(columnComment, "：", ":")
			if model.EnumModel != "" && pkg.HasEnum(columnComment) {
				model.EnumTitle = columnComment[:strings.Index(columnComment, "(")]
				model.EnumRemark = columnComment[strings.Index(columnComment, "(")+1 : len(columnComment)-1]
			}

			if model.DataType == "decimal.Decimal" {
				m.HasDecimal = true
			}

			for _, resp := range m.Responses {
				if resp.Name == model.DataType {
					model.ResponseModelData = resp
				}
			}
		}

		if response.ExtendName != "" {
			for _, resp := range m.Responses {
				if resp.Name == response.ExtendName {
					resp.ExtendResponseModelData = response
				}
			}
		}
	}
}

func (m *EntireModel) FillInterfaces() {
	if len(m.Interfaces) == 0 {
		return
	}

	for _, inter := range m.Interfaces {
		inter.GroupUpperCamel = strcase.ToCamel(inter.Group)
		inter.GroupLowerCamel = strcase.ToLowerCamel(inter.Group)
		inter.GroupSnake = strcase.ToSnake(inter.Group)

		for _, model := range inter.Models {
			model.RelativePath = strings.TrimPrefix(model.RelativePath, "/")

			if model.MethodType == "" {
				model.MethodType = MethodTypePost
			}

			if model.InterfaceType == InterfaceTypePage {
				inter.HasPage = true
			}

			if model.InterfaceType == InterfaceTypePage || model.InterfaceType == InterfaceTypeList {
				m.HasQuery = true
				inter.HasQuery = true
			}

			if model.RequestModelName == "Id" {
				m.HasId = true
				inter.HasId = true
				model.RequestModelNameExp = "cm.IdReq"
			} else if model.RequestModelName != "" {
				model.RequestModelNameExp = "model." + strcase.ToCamel(model.RequestModelName)
			} else {
				model.RequestModelNameExp = "result.Void"
			}

			if model.RequestModelName != "" {
				for _, request := range m.Requests {
					if request.Name == model.RequestModelName {
						model.RequestModelData = request

						if model.InterfaceType == InterfaceTypePage {
							request.IsPage = true
							request.IsPageOrList = true
							m.HasPage = true
							inter.HasPage = true
						} else if model.InterfaceType == InterfaceTypeList {
							request.IsPageOrList = true
						}
					}
				}
			}

			if model.ResponseModelName != "" {
				model.ResponseModelNameExp = model.ResponseModelName
				for _, response := range m.Responses {
					if model.ResponseModelName == response.Name {
						model.ResponseModelData = response
						model.ResponseModelHasPointer = true
						if model.InterfaceType == InterfaceTypePage {
							model.ResponseModelNameExp = "*page.PageList[model." + model.ResponseModelName + "]"
						} else if model.InterfaceType == InterfaceTypeList {
							model.ResponseModelNameExp = "[]*model." + model.ResponseModelName
						} else {
							model.ResponseModelNameExp = "*model." + model.ResponseModelName
						}
						break
					}
				}
			} else {
				model.ResponseModelNameExp = "*result.Void"
			}

			if len(model.AllowProducts) > 0 {
				products := dgcoll.MapToList(model.AllowProducts, func(name string) string {
					return productMap[name]
				})
				model.AllowProductsExp = "[]int{" + strings.Join(products, ",") + "}"
				inter.HasProducts = true
				m.HasProducts = true
			}

			if model.LogLevel != "" {
				model.LogLevelExp = logLevelMap[model.LogLevel]
			}

			if model.MethodName == "" {
				parts := strings.Split(model.RelativePath, "/")
				model.MethodName = parts[len(parts)-1]
			}

			model.MethodNameExp = strcase.ToCamel(model.MethodName)
			model.DbTableUpperCamel = strcase.ToCamel(model.DbModelName)
			model.DbTableLowerCamel = strcase.ToLowerCamel(model.DbModelName)

			if (model.RequestModelName != "" && model.RequestModelName != "Id") || model.ResponseModelName != "" {
				inter.HasModel = true
			}
			if model.DbModelName != "" {
				inter.HasDbTable = true
			}
		}
	}
}

func (m *EntireModel) FillConverters() {
	interfaceModels := dgcoll.FlatMapToList(m.Interfaces, func(inter *InterfaceModelData) []*InterfaceModel {
		return inter.Models
	})
	dbTableNames := dgcoll.FilterAndMapToSet(interfaceModels, func(interfaceModel *InterfaceModel) bool {
		return interfaceModel.DbModelName != ""
	}, func(interfaceModel *InterfaceModel) string {
		return interfaceModel.DbModelName
	})
	if len(dbTableNames) == 0 {
		return
	}
	dbTableName2RequestsMap := dgcoll.Extract2KeySetMap(interfaceModels, func(interfaceModel *InterfaceModel) string {
		return interfaceModel.DbModelName
	}, func(interfaceModel *InterfaceModel) *RequestModelData {
		if interfaceModel.RequestModelName != "" {
			for _, request := range m.Requests {
				if interfaceModel.RequestModelName == request.Name {
					request.InterfaceType = interfaceModel.InterfaceType
					return request
				}
			}
		}

		return nil
	})
	dbTableName2ResponsesMap := dgcoll.Extract2KeySetMap(interfaceModels, func(interfaceModel *InterfaceModel) string {
		return interfaceModel.DbModelName
	}, func(interfaceModel *InterfaceModel) *ResponseModelData {
		if interfaceModel.ResponseModelName != "" {
			for _, response := range m.Responses {
				if interfaceModel.ResponseModelName == response.Name {
					response.InterfaceType = interfaceModel.InterfaceType
					return response
				}
			}
		}

		return nil
	})

	converters := dgcoll.MapToList(dbTableNames, func(dbTableName string) *ConverterData {
		requests := dbTableName2RequestsMap[dbTableName]
		requests = dgcoll.FilterList(requests, func(request *RequestModelData) bool {
			return request != nil
		})

		responses := dbTableName2ResponsesMap[dbTableName]
		responses = dgcoll.FilterList(responses, func(response *ResponseModelData) bool {
			return response != nil
		})

		hasEnum := false
		for _, response := range responses {
			for _, responseModel := range response.Models {
				if responseModel.EnumModel != "" && strings.HasSuffix(responseModel.FieldName, "Name") {
					hasEnum = true
					responseModel.EnumFieldName = strings.TrimSuffix(responseModel.FieldName, "Name")
					break
				}
			}
		}

		return &ConverterData{
			DbTableUpperCamel: strcase.ToCamel(dbTableName),
			DbTableLowerCamel: strcase.ToLowerCamel(dbTableName),
			Requests:          requests,
			Responses:         responses,
			HasEnum:           hasEnum,
		}
	})

	m.Converters = converters
}

func (m *EntireModel) FillPackagePrefix(packagePrefix string) {
	m.PackagePrefix = packagePrefix

	if len(m.Interfaces) > 0 {
		for _, inter := range m.Interfaces {
			inter.PackagePrefix = packagePrefix
		}
	}

	if len(m.Converters) > 0 {
		for _, converter := range m.Converters {
			converter.PackagePrefix = packagePrefix
		}
	}
}

func BatchReflectToRequestModelData(objs []any) []*RequestModelData {
	return dgcoll.MapToList(objs, func(obj any) *RequestModelData {
		return ReflectToRequestModelData(obj)
	})
}

func ReflectToRequestModelData(obj any) *RequestModelData {
	tpe := reflect.TypeOf(obj)
	for tpe.Kind() == reflect.Pointer {
		tpe = tpe.Elem()
	}
	cnt := tpe.NumField()

	rmd := &RequestModelData{
		Name:           tpe.Name(),
		UpperCamelName: strcase.ToCamel(tpe.Name()),
	}

	for i := 0; i < cnt; i++ {
		field := tpe.Field(i)
		tag := field.Tag

		rm := &RequestModel{
			FieldName:      field.Name,
			DataType:       field.Type.Name(),
			IsPointer:      field.Type.Kind() == reflect.Pointer,
			IsArray:        field.Type.Kind() == reflect.Slice,
			Nullable:       strings.Contains(tag.Get("binding"), "required"),
			VerifyRules:    tag.Get("binding"),
			Remark:         tag.Get("remark"),
			UpperCamelName: strcase.ToCamel(field.Name),
			LowerCamelName: strcase.ToLowerCamel(field.Name),
		}

		rmd.Models = append(rmd.Models, rm)
	}

	return rmd
}

func BatchReflectToResponseModelData(objs []any) []*ResponseModelData {
	return dgcoll.MapToList(objs, func(obj any) *ResponseModelData {
		return ReflectToResponseModelData(obj)
	})
}

func ReflectToResponseModelData(obj any) *ResponseModelData {
	tpe := reflect.TypeOf(obj)
	for tpe.Kind() == reflect.Pointer {
		tpe = tpe.Elem()
	}
	cnt := tpe.NumField()

	rmd := &ResponseModelData{
		Name:           tpe.Name(),
		UpperCamelName: strcase.ToCamel(tpe.Name()),
	}

	for i := 0; i < cnt; i++ {
		field := tpe.Field(i)
		tag := field.Tag

		rm := &ResponseModel{
			FieldName:  field.Name,
			DataType:   field.Type.Name(),
			IsPointer:  field.Type.Kind() == reflect.Pointer,
			IsArray:    field.Type.Kind() == reflect.Slice,
			Nullable:   strings.Contains(tag.Get("binding"), "required"),
			IsMediaUrl: tag.Get("appendUid") != "",
			Remark:     tag.Get("remark"),
		}

		rmd.Models = append(rmd.Models, rm)
	}

	return rmd
}
