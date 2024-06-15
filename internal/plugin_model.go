package internal

type EntireModel struct {
	Enums      []*EnumModelData      `json:"enums,omitempty"`
	Requests   []*RequestModelData   `json:"requests,omitempty"`
	Responses  []*ResponseModelData  `json:"responses,omitempty"`
	Interfaces []*InterfaceModelData `json:"interfaces,omitempty"`
	Export     *ExportConfigData     `json:"export,omitempty"`
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
}

type RequestModel struct {
	FieldName   string `json:"fieldName,omitempty"`
	DataType    string `json:"dataType,omitempty"`
	IsPointer   bool   `json:"isPointer,omitempty"`
	IsArray     bool   `json:"isArray,omitempty"`
	Nullable    bool   `json:"nullable,omitempty"`
	VerifyRules string `json:"verifyRules,omitempty"`
	EnumModel   string `json:"enumModel,omitempty"`
	Remark      string `json:"remark,omitempty"`
}

type RequestModelData struct {
	Name       string          `json:"name,omitempty"`
	Models     []*RequestModel `json:"models,omitempty"`
	ExtendName string          `json:"extendName,omitempty"`
}

type ResponseModel struct {
	FieldName  string `json:"fieldName,omitempty"`
	DataType   string `json:"dataType,omitempty"`
	IsPointer  bool   `json:"isPointer,omitempty"`
	IsArray    bool   `json:"isArray,omitempty"`
	Nullable   bool   `json:"nullable,omitempty"`
	IsMediaUrl bool   `json:"isMediaUrl,omitempty"`
	EnumModel  string `json:"enumModel,omitempty"`
	Remark     string `json:"remark,omitempty"`
}

type ResponseModelData struct {
	Name       string           `json:"name,omitempty"`
	Models     []*ResponseModel `json:"models,omitempty"`
	ExtendName string           `json:"extendName,omitempty"`
}

type InterfaceModel struct {
	InterfaceType     string   `json:"interfaceType,omitempty"`
	RelativePath      string   `json:"relativePath,omitempty"`
	RequestModelName  string   `json:"requestModelName,omitempty"`
	ResponseModelName string   `json:"responseModelName,omitempty"`
	NonLogin          bool     `json:"nonLogin,omitempty"`
	AllowRoles        []string `json:"allowRoles,omitempty"`
	AllowProducts     []string `json:"allowProducts,omitempty"`
	LogLevel          string   `json:"logLevel,omitempty"`
	NotLogSQL         bool     `json:"notLogSQL,omitempty"`
	Remark            string   `json:"remark,omitempty"`
}

type InterfaceModelData struct {
	Group       string            `json:"group,omitempty"`
	RoutePrefix string            `json:"routePrefix,omitempty"`
	Models      []*InterfaceModel `json:"models,omitempty"`
}

type ExportConfigData struct {
	ServerOutput  string `json:"serverOutput,omitempty"`
	ClientOutput  string `json:"clientOutput,omitempty"`
	PackagePrefix string `json:"packagePrefix,omitempty"`
	ApifoxOutput  string `json:"apifoxOutput,omitempty"`
}
