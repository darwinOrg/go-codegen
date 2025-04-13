package dgen

import (
	"fmt"
	dgcoll "github.com/darwinOrg/go-common/collection"
	"github.com/darwinOrg/go-common/model"
	"github.com/darwinOrg/go-common/utils"
	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	_ "github.com/pingcap/tidb/parser/test_driver"
	"io/fs"
	"os"
	"strings"
)

func ParseSql(sql string, packagePrefix string, outputPath string) error {
	metas, err := BuildTableMetas(sql)
	if err != nil {
		return err
	}
	if len(metas) == 0 {
		return nil
	}

	_ = os.Mkdir(outputPath, fs.ModeDir|fs.ModePerm)

	for _, meta := range metas {
		meta.PackagePrefix = packagePrefix
		if parseMeta(outputPath, meta) != nil {
			return err
		}
	}

	return nil
}

func ParseToDbModelDataList(sqlFile string) ([]*DbModelData, error) {
	sql := utils.MustReadFileString(sqlFile)
	metas, err := BuildTableMetas(sql)
	if err != nil {
		return nil, err
	}

	return dgcoll.MapToList(metas, func(meta *Meta) *DbModelData {
		return meta.ConvertToDbModelData(sql)
	}), nil
}

func BuildTableMetas(sql string) ([]*Meta, error) {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var metas []*Meta

	for _, node := range stmtNodes {
		root, ok := node.(*ast.CreateTableStmt)
		if !ok {
			continue
		}
		meta := &Meta{}
		root.Accept(meta)
		metas = append(metas, meta)

		for _, option := range root.Options {
			if option.Tp == ast.TableOptionComment && option.StrValue != "" {
				meta.TableComment = option.StrValue
			}
		}

		meta.QueryColumns = dgcoll.FilterList(meta.Columns, func(c *Column) bool {
			return !dgcoll.Contains(ignoreQueryModelFieldNames, c.DbName)
		})

		meta.ModifyColumns = dgcoll.FilterList(meta.Columns, func(c *Column) bool {
			return !dgcoll.Contains(ignoreModifyModelFieldNames, c.DbName) &&
				!strings.Contains(c.DbName, "status") &&
				!strings.Contains(c.DbName, "state")
		})

		meta.CreateColumns = dgcoll.FilterList(meta.Columns, func(c *Column) bool {
			return !dgcoll.Contains(ignoreCreateModelFieldNames, c.DbName)
		})

		meta.EnumMap = map[*Column][]*model.KeyValuePair[string, string]{}

		for _, column := range meta.Columns {
			columnComment := strings.TrimSpace(column.Comment)
			if columnComment == "" {
				continue
			}
			columnComment = strings.ReplaceAll(columnComment, "（", "(")
			columnComment = strings.ReplaceAll(columnComment, "）", ")")
			columnComment = strings.ReplaceAll(columnComment, "，", ",")
			columnComment = strings.ReplaceAll(columnComment, "：", ":")

			matches := enumRegexp.FindStringSubmatch(columnComment)
			if len(matches) == 0 {
				continue
			}

			pairsStr := matches[1]
			pairParts := strings.Split(strings.TrimSpace(pairsStr), ",")

			for _, pairPart := range pairParts {
				kvs := strings.Split(strings.TrimSpace(pairPart), ":")

				if len(kvs) == 2 {
					key := strings.TrimSpace(kvs[0])
					value := strings.TrimSpace(kvs[1])
					meta.HasEnum = true
					column.HasEnum = true
					enumName := columnComment[:strings.Index(columnComment, "(")]
					enumRemark := columnComment[strings.Index(columnComment, "(")+1 : len(columnComment)-1]
					column.EnumName = enumName
					column.EnumRemark = enumRemark

					meta.EnumMap[column] = append(meta.EnumMap[column], &model.KeyValuePair[string, string]{
						Key:   key,
						Value: value,
					})
				}
			}
		}
	}

	return metas, nil
}
