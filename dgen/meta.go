package dgen

import (
	"fmt"
	"github.com/darwinOrg/go-codegen/pkg"
	dgcoll "github.com/darwinOrg/go-common/collection"
	"github.com/darwinOrg/go-common/model"
	"github.com/iancoleman/strcase"
	"github.com/pingcap/tidb/parser/ast"
	"github.com/pingcap/tidb/parser/test_driver"
	"log"
	"strings"
)

var keywordsMap = map[string]int{}

var (
	ignoreQueryModelFieldNames  = []string{"company_id", "org_id", "mask", "deleted", "test", "vsn", "op_id", "created_by", "created_at", "modified_by", "modified_at"}
	ignoreModifyModelFieldNames = append(ignoreQueryModelFieldNames, "status", "state")
	ignoreCreateModelFieldNames = append(ignoreModifyModelFieldNames, "id")
)

func init() {
	keywordsMap["case"] = 1
	keywordsMap["key"] = 1
	keywordsMap["table"] = 1
	keywordsMap["column"] = 1
}

type Column struct {
	GoName         string
	DbName         string
	DbType         string
	ModelType      string
	LowerCamelName string
	IsNull         bool
	HasEnum        bool
	EnumName       string
	EnumRemark     string
	Comment        string
}

type Meta struct {
	PackagePrefix   string
	GoTable         string
	TableName       string
	TableComment    string
	LowerCamelName  string
	KebabName       string
	Columns         []*Column
	QueryColumns    []*Column
	CreateColumns   []*Column
	ModifyColumns   []*Column
	EnumMap         map[*Column][]*model.KeyValuePair[string, string]
	AutoColumn      string
	HasType         bool
	HasEnum         bool
	ModelHasType    bool
	HasDecimal      bool
	ModelHasDecimal bool
}

func (meta *Meta) Enter(in ast.Node) (ast.Node, bool) {
	switch in.(type) {
	case *ast.ColumnDef:
		def := in.(*ast.ColumnDef)
		c := toColumn(def)
		if c != nil {
			if strings.HasPrefix(c.DbType, "ttypes.") {
				meta.HasType = true
			}
			if strings.HasPrefix(c.DbType, "ttypes.") && !dgcoll.Contains(ignoreQueryModelFieldNames, c.DbName) {
				meta.ModelHasType = true
			}
			if isAuto(def) {
				meta.AutoColumn = c.DbName
			}
			if c.DbType == "decimal.Decimal" {
				meta.HasDecimal = true
			}
			if c.DbType == "decimal.Decimal" && !dgcoll.Contains(ignoreQueryModelFieldNames, c.DbName) {
				meta.ModelHasDecimal = true
			}
			meta.Columns = append(meta.Columns, toColumn(def))
		}
	case *ast.TableName:
		name := in.(*ast.TableName)
		meta.TableName = escapeKeyName(name.Name.O)
		meta.GoTable = strcase.ToCamel(name.Name.O)
		meta.LowerCamelName = strcase.ToLowerCamel(name.Name.O)
		meta.KebabName = strcase.ToKebab(name.Name.O)
	}

	return in, false
}

func escapeKeyName(name string) string {
	lcName := strings.ToLower(name)
	if keywordsMap[lcName] == 0 {
		return name
	}
	return fmt.Sprintf("`%s`", name)
}

func (meta *Meta) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func (meta *Meta) ConvertToDbModelData(sql string) *DbModelData {
	return &DbModelData{
		Sql:  sql,
		Name: meta.TableName,
		Models: dgcoll.MapToList(meta.Columns, func(column *Column) *DbModel {
			return &DbModel{
				DbFieldName:    column.DbName,
				ModelFieldName: column.GoName,
				DbDataType:     column.DbType,
				ModelDataType:  column.ModelType,
				Nullable:       column.IsNull,
				Remark:         column.Comment,
			}
		}),
	}
}

func isAuto(def *ast.ColumnDef) bool {
	for _, op := range def.Options {
		if ast.ColumnOptionAutoIncrement == op.Tp {
			return true
		}
	}
	return false
}

func toColumn(def *ast.ColumnDef) *Column {
	c := &Column{
		DbName: escapeKeyName(def.Name.Name.O),
		GoName: strcase.ToCamel(def.Name.Name.O),
	}
	c.LowerCamelName = strcase.ToLowerCamel(c.GoName)

	for _, op := range def.Options {
		if op.Tp == ast.ColumnOptionComment {
			c.Comment = op.Expr.(*test_driver.ValueExpr).Datum.GetString()
			break
		}
	}

	isNull := true
	for _, op := range def.Options {
		if op.Tp == ast.ColumnOptionNotNull {
			isNull = false
			break
		}
	}
	c.IsNull = isNull

	dbType := pkg.ToDbTypeString(def.Tp.GetType(), def.Tp.GetFlag(), isNull)
	if dbType == "" {
		log.Printf("%s不能推断出数据库类型\n", c.DbName)
		return nil
	}
	c.DbType = dbType
	c.ModelType = pkg.AdjustDbType(dbType)

	return c
}
