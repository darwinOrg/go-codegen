package internal

type EntireModel struct {
	Dbs        []*DbModelData        `json:"dbs,omitempty"`
	Enums      []*EnumModelData      `json:"enums,omitempty"`
	Requests   []*RequestModelData   `json:"requests,omitempty"`
	Responses  []*ResponseModelData  `json:"responses,omitempty"`
	Interfaces []*InterfaceModelData `json:"interfaces,omitempty"`
	Export     *ExportConfigData     `json:"export,omitempty"`

	HasDaogType    bool   `json:"hasDaogType,omitempty"`
	HasDecimal     bool   `json:"hasDecimal,omitempty"`
	HasPage        bool   `json:"hasPage,omitempty"`
	HasId          bool   `json:"hasId,omitempty"`
	HasModel       bool   `json:"hasModel,omitempty"`
	UpperCamelName string `json:"upperCamelName,omitempty"`
}

type DbModel struct {
	DbFieldName    string `json:"dbFieldName,omitempty"`
	ModelFieldName string `json:"modelFieldName,omitempty"`
	DbDataType     string `json:"dbDataType,omitempty"`
	ModelDataType  string `json:"modelDataType,omitempty"`
	Nullable       bool   `json:"nullable,omitempty"`
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

	LowerCamelName string `json:"lowerCamelName,omitempty"`
}

type RequestModelData struct {
	Name       string          `json:"name,omitempty"`
	Models     []*RequestModel `json:"models,omitempty"`
	ExtendName string          `json:"extendName,omitempty"`

	IsPage bool `json:"isPage,omitempty"`
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

	LowerCamelName string `json:"lowerCamelName,omitempty"`
	EnumTitle      string `json:"enumTitle,omitempty"`
	EnumRemark     string `json:"enumRemark,omitempty"`
}

type ResponseModelData struct {
	Name       string           `json:"name,omitempty"`
	Models     []*ResponseModel `json:"models,omitempty"`
	ExtendName string           `json:"extendName,omitempty"`
}

type InterfaceModel struct {
	InterfaceType     string   `json:"interfaceType,omitempty"`
	RelativePath      string   `json:"relativePath,omitempty"`
	MethodName        string   `json:"methodName,omitempty"`
	RequestModelName  string   `json:"requestModelName,omitempty"`
	ResponseModelName string   `json:"responseModelName,omitempty"`
	NonLogin          bool     `json:"nonLogin,omitempty"`
	AllowRoles        []string `json:"allowRoles,omitempty"`
	AllowProducts     []string `json:"allowProducts,omitempty"`
	LogLevel          string   `json:"logLevel,omitempty"`
	NotLogSQL         bool     `json:"notLogSQL,omitempty"`
	Remark            string   `json:"remark,omitempty"`

	ResponseModelHasPointer bool   `json:"responseModelHasPointer,omitempty"`
	AllowProductsExp        string `json:"allowProductsExp,omitempty"`
	LogLevelExp             string `json:"logLevelExp,omitempty"`
	MethodNameExp           string `json:"methodNameExp,omitempty"`
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
