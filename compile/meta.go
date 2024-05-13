package compile

import (
	"fmt"
	dgcoll "github.com/darwinOrg/go-common/collection"
	"github.com/darwinOrg/go-common/model"
	"github.com/iancoleman/strcase"
	"github.com/pingcap/tidb/parser/ast"
	"github.com/pingcap/tidb/parser/mysql"
	"github.com/pingcap/tidb/parser/test_driver"
	"log"
	"strings"
)

var keywordsMap = map[string]int{}

var (
	ignoreQueryModelFieldNames  = []string{"mask", "deleted", "test", "vsn", "op_id", "created_by", "created_at", "modified_by", "modified_at"}
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
	LowerCamelName string
	IsNull         bool
	HasEnum        bool
	Comment        string
}

type Meta struct {
	ProjectPath     string
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

	dbType := toDbTypeString(def.Tp.GetType(), def.Tp.GetFlag(), isNull)
	if dbType == "" {
		log.Printf("%s不能推断出数据库类型\n", c.DbName)
		return nil
	}
	c.DbType = dbType

	return c
}

func toDbTypeString(tp byte, flag uint, isNull bool) string {
	switch tp {
	case mysql.TypeTiny:
		if mysql.HasUnsignedFlag(flag) {
			return "uint8"
		}
		return "int8"
	case mysql.TypeShort:
		if mysql.HasUnsignedFlag(flag) {
			return "uint16"
		}
		return "int16"
	case mysql.TypeLong:
		if mysql.HasUnsignedFlag(flag) {
			return "uint32"
		}
		return "int32"
	case mysql.TypeFloat:
		return "float32"
	case mysql.TypeDouble:
		return "float64"
	case mysql.TypeTimestamp:
		if isNull {
			return "ttypes.NilableDatetime"
		}
		return "ttypes.NormalDatetime"
	case mysql.TypeDate:
		if isNull {
			return "ttypes.NilableDate"
		}
		return "ttypes.NormalDate"
	case mysql.TypeDatetime:
		if isNull {
			return "ttypes.NilableDatetime"
		}
		return "ttypes.NormalDatetime"
	case mysql.TypeNewDate:
		if isNull {
			return "ttypes.NilableDate"
		}
		return "ttypes.NormalDate"
	case mysql.TypeInt24:
		if mysql.HasUnsignedFlag(flag) {
			return "uint32"
		}
		return "int32"
	case mysql.TypeLonglong:
		if mysql.HasUnsignedFlag(flag) {
			return "uint64"
		}
		return "int64"
	case mysql.TypeVarchar:
		if isNull {
			return "ttypes.NilableString"
		}
		return "string"
	case mysql.TypeJSON:
		if isNull {
			return "ttypes.NilableString"
		}
		return "string"
	case mysql.TypeBit:
		return "bool"
	case mysql.TypeVarString:
		if isNull {
			return "ttypes.NilableString"
		}
		return "string"
	case mysql.TypeString:
		if isNull {
			return "ttypes.NilableString"
		}
		return "string"
	case mysql.TypeTinyBlob:
		if mysql.HasBinaryFlag(flag) {
			return "[]byte"
		}
		if isNull {
			return "ttypes.NilableString"
		}
		return "string"
	case mysql.TypeMediumBlob:
		if mysql.HasBinaryFlag(flag) {
			return "[]byte"
		}
		if isNull {
			return "ttypes.NilableString"
		}
		return "string"
	case mysql.TypeLongBlob:
		if mysql.HasBinaryFlag(flag) {
			return "[]byte"
		}
		if isNull {
			return "ttypes.NilableString"
		}
		return "string"
	case mysql.TypeBlob:
		if mysql.HasBinaryFlag(flag) {
			return "[]byte"
		}
		if isNull {
			return "ttypes.NilableString"
		}
		return "string"
	case mysql.TypeNewDecimal:
		return "decimal.Decimal"
	}

	return ""
}
