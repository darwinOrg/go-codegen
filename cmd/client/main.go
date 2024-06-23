package main

import (
	"dgen/internal"
	dgcoll "github.com/darwinOrg/go-common/collection"
)

func main() {
	entireModel := internal.InitEntireModel()

	interfaces := dgcoll.FilterList(entireModel.Interfaces, func(item *internal.InterfaceModelData) bool {
		return item.ExportClient
	})
	if len(interfaces) == 0 {
		return
	}
	entireModel.Interfaces = interfaces

	requestModelNames := dgcoll.FlatMapToSet(interfaces, func(inter *internal.InterfaceModelData) []string {
		return dgcoll.MapToSet(inter.Models, func(interfaceModel *internal.InterfaceModel) string {
			return interfaceModel.RequestModelName
		})
	})
	requestModelNames = dgcoll.FilterList(requestModelNames, func(requestModelName string) bool {
		return requestModelName != "" && requestModelName != "Id"
	})
	if len(requestModelNames) > 0 {
		entireModel.Requests = dgcoll.FilterList(entireModel.Requests, func(requestModel *internal.RequestModelData) bool {
			return dgcoll.Contains(requestModelNames, requestModel.Name)
		})
	} else {
		entireModel.Requests = []*internal.RequestModelData{}
	}

	responseModelNames := dgcoll.FlatMapToSet(interfaces, func(inter *internal.InterfaceModelData) []string {
		return dgcoll.MapToSet(inter.Models, func(interfaceModel *internal.InterfaceModel) string {
			return interfaceModel.ResponseModelName
		})
	})
	responseModelNames = dgcoll.FilterList(responseModelNames, func(responseModelName string) bool {
		return responseModelName != ""
	})
	if len(responseModelNames) > 0 {
		entireModel.Responses = dgcoll.FilterList(entireModel.Responses, func(responseModel *internal.ResponseModelData) bool {
			return dgcoll.Contains(responseModelNames, responseModel.Name)
		})
	} else {
		entireModel.Responses = []*internal.ResponseModelData{}
	}

	entireModel.Fill(entireModel.Export.ClientPackagePrefix)
	internal.ClientParser.Parse(entireModel)
}
