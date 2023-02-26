package sqltools

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/miacio/sql-to-code/lib"
	"github.com/miacio/sql-to-code/tpl"
)

// Table 表
type Table struct {
	Package   string   `json:"package"`    // 包名(用于设定当前表转结构体时所在包名称)
	Name      string   `json:"name"`       // 表名(基于结构体名称-对表名进行驼峰等操作后结果)
	TableName string   `json:"table_name"` // 表名(原数据库表名)
	Comment   string   `json:"comment"`    // 表注释
	Fields    []Field  `json:"fields"`     // 列
	Imports   []string `json:"imports"`    // 引入的包名
}

// ToCode 表转结构体
func (t *Table) ToCode() (string, error) {
	tl := template.Must(template.New("tpl").Parse(tpl.StructTpl))
	var res bytes.Buffer
	if err := tl.Execute(&res, t); err != nil {
		return "", err
	}
	return lib.CommonInitialisms(res.String()), nil
}

// Field 字段
type Field struct {
	IsPrimaryKey     bool `json:"is_primary_key"`      // 是否为主键
	IsAutoIncrement  bool `json:"is_auto_increment"`   // 是否自增长
	IsAutoCreateTime bool `json:"is_auto_create_time"` // 是否创建时追踪当前时间
	IsAutoUpdateTime bool `json:"is_auto_update_time"` // 是否创建/修改时追踪当前时间
	IsNotNull        bool `json:"is_not_null"`         // 是否不允许为空
	IsUnsigned       bool `json:"is_unsigned"`         // 是否使用了无符号参数类型

	Name       string `json:"name"`        // 结构体属性名称
	Type       string `json:"type"`        // 结构体字段类型
	FieldName  string `json:"field_name"`  // 列名
	FieldType  string `json:"field_type"`  // 列字段类型
	Comment    string `json:"commont"`     // 字段注释
	DefaultVal string `json:"default_val"` // 字段默认值
	FieldTag   string `json:"field_tag"`   // 标签
}

// GetGormTag 获取gorm tag
func (field *Field) GetGormTag() string {
	gormTag := make([]string, 0)
	gormTag = append(gormTag, "column:"+field.FieldName)

	if field.IsPrimaryKey {
		gormTag = append(gormTag, "primaryKey")
		// 判断是否为自增长 (仅主键字段可用)
		if field.IsAutoIncrement {
			gormTag = append(gormTag, "autoIncrement")
		}
	}
	if field.IsAutoCreateTime {
		gormTag = append(gormTag, "autoCreateTime")
	}
	if field.IsAutoUpdateTime {
		gormTag = append(gormTag, "autoUpdateTime")
	}

	if field.IsNotNull {
		gormTag = append(gormTag, "not null")
	}
	if field.DefaultVal != "" {
		gormTag = append(gormTag, "default:"+field.DefaultVal)
	}
	return fmt.Sprintf("gorm:\"%s\"", strings.Join(gormTag, ";"))
}

// GenerateTag 生成标签 默认值以字段数据库名称为准
// upperFirstLetter 是否要求首字母大写
// humpNaming 是否要求驼峰命名
func (field *Field) GenerateTag(tagName string, upperFirstLetter, humpNaming bool) string {
	var tag string
	if humpNaming {
		tag = HumpNaming(field.FieldName)
	}
	if !upperFirstLetter {
		tag = strings.ToLower(tag[:1]) + tag[1:]
	}
	return fmt.Sprintf("%s:\"%s\"", tagName, tag)
}

// FieldOtherType 字段其它类型引入
type FieldOtherType struct {
	ImportPath string `json:"importPath" toml:"importPath"` // 引用的包地址
	FieldType  string `json:"fieldType" toml:"fieldType"`   // 代码中字段类型
	DbType     string `json:"dbType" toml:"dbType"`         // 数据库对应的类型
}
