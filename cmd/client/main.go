package main

import (
	"github.com/darwinOrg/go-codegen/dgen"
	dgcoll "github.com/darwinOrg/go-common/collection"
)

func main() {
	entireModel := dgen.InitEntireModel()

	interfaces := dgcoll.FilterList(entireModel.Interfaces, func(item *dgen.InterfaceModelData) bool {
		return item.ExportClient
	})
	if len(interfaces) == 0 {
		return
	}
	entireModel.Interfaces = interfaces

	requestModelNames := dgcoll.FlatMapToSet(interfaces, func(inter *dgen.InterfaceModelData) []string {
		return dgcoll.MapToSet(inter.Models, func(interfaceModel *dgen.InterfaceModel) string {
			return interfaceModel.RequestModelName
		})
	})
	requestModelNames = dgcoll.FilterList(requestModelNames, func(requestModelName string) bool {
		return requestModelName != "" && requestModelName != dgen.RequestModelId && requestModelName != dgen.RequestModelIds
	})
	if len(requestModelNames) > 0 {
		entireModel.Requests = dgcoll.FilterList(entireModel.Requests, func(requestModel *dgen.RequestModelData) bool {
			return dgcoll.Contains(requestModelNames, requestModel.Name)
		})
	} else {
		entireModel.Requests = []*dgen.RequestModelData{}
	}

	responseModelNames := dgcoll.FlatMapToSet(interfaces, func(inter *dgen.InterfaceModelData) []string {
		return dgcoll.MapToSet(inter.Models, func(interfaceModel *dgen.InterfaceModel) string {
			return interfaceModel.ResponseModelName
		})
	})
	responseModelNames = dgcoll.FilterList(responseModelNames, func(responseModelName string) bool {
		return responseModelName != ""
	})
	if len(responseModelNames) > 0 {
		entireModel.Responses = dgcoll.FilterList(entireModel.Responses, func(responseModel *dgen.ResponseModelData) bool {
			return dgcoll.Contains(responseModelNames, responseModel.Name)
		})
	} else {
		entireModel.Responses = []*dgen.ResponseModelData{}
	}

	entireModel.Fill(entireModel.Export.ClientPackagePrefix)
	_ = dgen.ClientParser.Parse(entireModel)
}
