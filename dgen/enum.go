package dgen

import "regexp"

const (
	InterfaceTypePage     = "分页"
	InterfaceTypeList     = "列表"
	InterfaceTypeCreate   = "新建"
	InterfaceTypeModify   = "修改"
	InterfaceTypeSave     = "保存"
	InterfaceTypeDelete   = "删除"
	InterfaceTypeDetail   = "详情"
	InterfaceTypeReadOnly = "只读"
	InterfaceTypeOther    = "其它"
)

const (
	MethodTypePost = "Post"
	MethodTypeGet  = "Get"
)

const (
	LogLevelAll    = "全部"
	LogLevelParam  = "请求参数"
	LogLevelReturn = "返回响应"
	LogLevelNone   = "无"
)

const (
	RequestModelId  = "Id"
	RequestModelIds = "Ids"
)

var enumRegexp = regexp.MustCompile(`\(([^)]+)\)`)

func HasEnum(str string) bool {
	matches := enumRegexp.FindStringSubmatch(str)
	return len(matches) > 0
}
