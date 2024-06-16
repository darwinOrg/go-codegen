package main

import (
	"dgen/internal"
	"flag"
	"fmt"
	_ "github.com/darwinOrg/daog-ext"
	"github.com/darwinOrg/go-common/utils"
	_ "github.com/darwinOrg/go-logger"
	_ "github.com/darwinOrg/go-web/wrapper"
	_ "github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
	_ "github.com/rolandhe/daog"
	_ "github.com/shopspring/decimal"
	"os"
	"os/exec"
	"strings"
)

var (
	inputFile string
)

func init() {
	flag.StringVar(&inputFile, "i", "", "the code generation design file")
}

func main() {
	flag.Parse()

	if inputFile == "" {
		fmt.Println("Please input code generation design file")
		os.Exit(1)
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
			}
		}
	}

	if len(entireModel.Responses) > 0 {
		for _, response := range entireModel.Responses {
			for _, model := range response.Models {
				model.LowerCamelName = strcase.ToLowerCamel(model.FieldName)

				if model.EnumModel != "" {
					columnComment := model.Remark
					columnComment = strings.ReplaceAll(columnComment, "（", "(")
					columnComment = strings.ReplaceAll(columnComment, "）", ")")
					columnComment = strings.ReplaceAll(columnComment, "，", ",")
					columnComment = strings.ReplaceAll(columnComment, "：", ":")
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
			for _, model := range inter.Models {
				if model.InterfaceType == "分页" && model.ResponseModelName != "" {
					for _, request := range entireModel.Requests {
						if request.Name == model.ResponseModelName {
							request.IsPage = true
							entireModel.HasPage = true
							break
						}
					}
				}
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

	internal.ServerGenerator.Generate(entireModel)

	outputPath := entireModel.Export.ServerOutput
	if outputPath == "./" {
		outputPath, _ = os.Getwd()
	}

	dirs := []string{"model"}
	for _, dir := range dirs {
		cmd := exec.Command("go", "fmt", outputPath+"/"+dir)
		err = cmd.Run()
		if err != nil {
			fmt.Printf("go fmt %s error: %v", dir, err)
		}
	}
}
