package pkg

import (
	"github.com/pingcap/tidb/parser/mysql"
	"strings"
)

func ToDbTypeString(tp byte, flag uint, isNull bool) string {
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

	return "string"
}

func AdjustDbType(dbType string) string {
	if strings.HasPrefix(dbType, "ttypes") {
		return "string"
	}

	return dbType
}
