package main

import (
	"dgen/internal"
	"dgen/pkg"
	"flag"
	"fmt"
	_ "github.com/darwinOrg/daog-ext"
	dgcoll "github.com/darwinOrg/go-common/collection"
	"github.com/darwinOrg/go-common/utils"
	_ "github.com/darwinOrg/go-logger"
	_ "github.com/darwinOrg/go-web/wrapper"
	_ "github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
	_ "github.com/rolandhe/daog"
	_ "github.com/shopspring/decimal"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	inputFile  string
	productMap = map[string]int{
		"RPA机器人":  10,
		"AI智能面试":  11,
		"金牌面试官":   12,
		"需求沟通助手":  13,
		"候选人沟通助手": 14,
		"人才库":     15,
	}
	logLevelMap = map[string]string{
		"全部":   "wrapper.LOG_LEVEL_ALL",
		"请求参数": "wrapper.LOG_LEVEL_PARAM",
		"返回响应": "wrapper.LOG_LEVEL_RETURN",
		"无":    "wrapper.LOG_LEVEL_NONE",
	}
)

func init() {
	flag.StringVar(&inputFile, "i", "./scripts/test.cui", "the code generation design file")
}

func main() {
	flag.Parse()

	if inputFile == "" {
		if len(os.Args) < 2 {
			fmt.Println("Please input code generation design file")
			os.Exit(1)
		}
		inputFile = os.Args[1]
	}
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Please input code generation design file")
		os.Exit(1)
	}

	entireModel, err := utils.ConvertJsonBytesToBean[internal.EntireModel](data)
	if err != nil {
		fmt.Println("Please input correct code generation design file")
		os.Exit(1)
	}

	if len(entireModel.Requests) > 0 {
		for _, request := range entireModel.Requests {
			for _, model := range request.Models {
				model.LowerCamelName = strcase.ToLowerCamel(model.FieldName)

				if !model.Nullable {
					if model.VerifyRules != "" {
						model.VerifyRules = "required," + model.VerifyRules
					} else {
						model.VerifyRules = "required"
					}
				}

				if strings.Contains(model.DataType, "ttypes") {
					entireModel.HasDaogType = true
				} else if model.DataType == "decimal.Decimal" {
					entireModel.HasDecimal = true
				}

				if model.DataType == "Id" {
					entireModel.HasId = true
				}
			}
		}
	}

	if len(entireModel.Responses) > 0 {
		for _, response := range entireModel.Responses {
			for _, model := range response.Models {
				model.LowerCamelName = strcase.ToLowerCamel(model.FieldName)

				columnComment := model.Remark
				columnComment = strings.ReplaceAll(columnComment, "（", "(")
				columnComment = strings.ReplaceAll(columnComment, "）", ")")
				columnComment = strings.ReplaceAll(columnComment, "，", ",")
				columnComment = strings.ReplaceAll(columnComment, "：", ":")
				if model.EnumModel != "" && pkg.HasEnum(columnComment) {
					model.EnumTitle = columnComment[:strings.Index(columnComment, "(")]
					model.EnumRemark = columnComment[strings.Index(columnComment, "(")+1 : len(columnComment)-1]
				}

				if strings.Contains(model.DataType, "ttypes") {
					entireModel.HasDaogType = true
				} else if model.DataType == "decimal.Decimal" {
					entireModel.HasDecimal = true
				}
			}
		}
	}

	if len(entireModel.Interfaces) > 0 {
		for _, inter := range entireModel.Interfaces {
			inter.GroupUpperCamel = strcase.ToCamel(inter.Group)
			inter.GroupLowerCamel = strcase.ToLowerCamel(inter.Group)

			for _, model := range inter.Models {
				if model.InterfaceType == "分页" && model.RequestModelName != "" {
					for _, request := range entireModel.Requests {
						if request.Name == model.RequestModelName {
							request.IsPage = true
							entireModel.HasPage = true
							break
						}
					}
				}

				if model.RequestModelName == "Id" {
					entireModel.HasId = true
					model.RequestModelNameExp = "cm.IdReq"
				} else if model.RequestModelName != "" {
					model.RequestModelNameExp = "model." + model.RequestModelName
				} else {
					model.RequestModelNameExp = "result.Void"
				}

				if model.ResponseModelName != "" {
					model.ResponseModelNameExp = model.ResponseModelName
					for _, response := range entireModel.Responses {
						if model.ResponseModelName == response.Name {
							model.ResponseModelHasPointer = true
							if model.InterfaceType == "分页" {
								model.ResponseModelNameExp = "*page.PageList[model." + model.ResponseModelName + "]"
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
					products := dgcoll.MapToList(model.AllowProducts, func(name string) int {
						return productMap[name]
					})
					model.AllowProductsExp = "[]int{" + dgcoll.JoinIntsByComma(products) + "}"
				}

				if model.LogLevel != "" {
					model.LogLevelExp = logLevelMap[model.LogLevel]
				}

				model.MethodNameExp = strcase.ToCamel(model.MethodName)
			}
		}
	}

	if entireModel.Enums == nil {
		entireModel.Enums = []*internal.EnumModelData{}
	}
	if entireModel.Requests == nil {
		entireModel.Requests = []*internal.RequestModelData{}
	}
	if entireModel.Responses == nil {
		entireModel.Responses = []*internal.ResponseModelData{}
	}
	if entireModel.Interfaces == nil {
		entireModel.Interfaces = []*internal.InterfaceModelData{}
	}

	entireModel.HasModel = len(entireModel.Requests) > 0 || len(entireModel.Responses) > 0

	filenameWithExt := filepath.Base(inputFile)
	filename := strings.TrimSuffix(filenameWithExt, filepath.Ext(inputFile))
	entireModel.UpperCamelName = strcase.ToCamel(filename)
	internal.ServerParser.Parse(entireModel, filename)

	outputPath := entireModel.Export.ServerOutput
	if outputPath == "./" {
		outputPath, _ = os.Getwd()
	}

	dirs := []string{"model", "handler", "router"}
	for _, dir := range dirs {
		cmd := exec.Command("go", "fmt", outputPath+"/"+dir)
		err = cmd.Run()
		if err != nil {
			fmt.Printf("go fmt %s error: %v", dir, err)
		}
	}
}
