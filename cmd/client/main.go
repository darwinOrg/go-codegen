package main

import (
	"github.com/darwinOrg/go-codegen/parser"
	dgcoll "github.com/darwinOrg/go-common/collection"
)

func main() {
	entireModel := parser.InitEntireModel()

	interfaces := dgcoll.FilterList(entireModel.Interfaces, func(item *parser.InterfaceModelData) bool {
		return item.ExportClient
	})
	if len(interfaces) == 0 {
		return
	}
	entireModel.Interfaces = interfaces

	requestModelNames := dgcoll.FlatMapToSet(interfaces, func(inter *parser.InterfaceModelData) []string {
		return dgcoll.MapToSet(inter.Models, func(interfaceModel *parser.InterfaceModel) string {
			return interfaceModel.RequestModelName
		})
	})
	requestModelNames = dgcoll.FilterList(requestModelNames, func(requestModelName string) bool {
		return requestModelName != "" && requestModelName != "Id"
	})
	if len(requestModelNames) > 0 {
		entireModel.Requests = dgcoll.FilterList(entireModel.Requests, func(requestModel *parser.RequestModelData) bool {
			return dgcoll.Contains(requestModelNames, requestModel.Name)
		})
	} else {
		entireModel.Requests = []*parser.RequestModelData{}
	}

	responseModelNames := dgcoll.FlatMapToSet(interfaces, func(inter *parser.InterfaceModelData) []string {
		return dgcoll.MapToSet(inter.Models, func(interfaceModel *parser.InterfaceModel) string {
			return interfaceModel.ResponseModelName
		})
	})
	responseModelNames = dgcoll.FilterList(responseModelNames, func(responseModelName string) bool {
		return responseModelName != ""
	})
	if len(responseModelNames) > 0 {
		entireModel.Responses = dgcoll.FilterList(entireModel.Responses, func(responseModel *parser.ResponseModelData) bool {
			return dgcoll.Contains(responseModelNames, responseModel.Name)
		})
	} else {
		entireModel.Responses = []*parser.ResponseModelData{}
	}

	entireModel.Fill(entireModel.Export.ClientPackagePrefix)
	_ = parser.ClientParser.Parse(entireModel)
}
