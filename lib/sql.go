package lib

import (
	"bytes"
	"os"
	"text/template"
)

// Table 表
type Table struct {
	Package           string  `json:"package"`             // 包名
	Name              string  `json:"name"`                // 结构体名
	TableName         string  `json:"table_name"`          // 表名
	Comment           string  `json:"comment"`             // 表注释
	Fields            []Field `json:"fields"`              // 列
	ContainsTimeField bool    `json:"contains_time_field"` // 是否存在时间字段
}

// Field 字段
type Field struct {
	IsPk          bool `json:"is_pk"`           // 是否为主键
	IsAuto        bool `json:"is_auto"`         // 是否自增长
	IsUnique      bool `json:"is_unique"`       // 是否唯一键
	IsNotNull     bool `json:"is_not_null"`     // 是否不允许为空
	IsIndex       bool `json:"is_index"`        // 是否为索引字段
	IsUniqueIndex bool `json:"is_unique_index"` // 是否创建唯一索引

	Name       string `json:"name"`        // 结构体属性名称
	FieldName  string `json:"field_name"`  // 列名
	FieldType  string `json:"field_type"`  // 列字段类型
	Comment    string `json:"commont"`     // 字段注释
	DefaultVal string `json:"default_val"` // 字段默认值
	FieldTag   string `json:"field_tag"`   // 标签
}

// ToCode 表转结构体
func (t *Table) ToCode() (string, error) {
	tpl, err := os.ReadFile("../tpl/struct.tpl")
	if err != nil {
		return "", err
	}
	for i := range t.Fields {
		field := t.Fields[i]
		tag := "`" + field.GetGormTag() + "`"
		t.Fields[i].FieldTag = tag
	}

	tl := template.Must(template.New("tpl").Parse(string(tpl)))
	var res bytes.Buffer
	if err := tl.Execute(&res, t); err != nil {
		return "", err
	}
	return CommonInitialisms(res.String()), nil
}

func ReadSql(val string) (*Table, error) {
	return nil, nil
}
