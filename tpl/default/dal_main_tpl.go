package _default

import "strings"

var DalMainTpl = `package dal

import (
	daogext "github.com/darwinOrg/daog-ext"
    "github.com/rolandhe/daog"
    {{if .HasType}}"github.com/rolandhe/daog/ttypes"{{end}}
    {{if .HasDecimal}}"github.com/shopspring/decimal"{{end}}
)

var {{.GoTable}}Fields = struct {
   {{range .Columns}}{{.GoName}} string
   {{end}}
}{
    {{range .Columns}}"{{.DbName}}",
    {{end}}
}

var  {{.GoTable}}Meta = &daog.TableMeta[{{.GoTable}}]{
    Table: "{{.TableName}}",
    Columns: []string {
        {{range .Columns}}"{{.DbName}}",
        {{end}}
    },
    AutoColumn: "{{.AutoColumn}}",
    LookupFieldFunc: func(columnName string,ins *{{.GoTable}},point bool) any {
        {{range .Columns}}if "{{.DbName}}" == columnName {
            if point {
                 return &ins.{{.GoName}}
            }
            return ins.{{.GoName}}
        }
        {{end}}
        return nil
    },
}

var {{.GoTable}}Dao daog.QuickDao[{{.GoTable}}] = &struct {
	daog.QuickDao[{{.GoTable}}]
}{
	daogext.NewBaseQuickDao({{.GoTable}}Meta),
}

type {{.GoTable}} struct {
    {{range .Columns}}{{.GoName}} {{.DbType}} ###json:"{{.LowerCamelName}}" remark:"{{.Comment}}"###
    {{end}}
}
`

func init() {
	DalMainTpl = strings.ReplaceAll(DalMainTpl, "###", "`")
}
