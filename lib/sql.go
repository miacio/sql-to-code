package lib

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/miacio/sql-to-code/tpl"
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
	// tpl, err := os.ReadFile("../tpl/struct.tpl")
	// if err != nil {
	// 	return "", err
	// }

	for i := range t.Fields {
		field := t.Fields[i]
		tag := "`" + field.GetGormTag() + "`"
		t.Fields[i].FieldTag = tag
	}

	tl := template.Must(template.New("tpl").Parse(tpl.StructTpl))
	var res bytes.Buffer
	if err := tl.Execute(&res, t); err != nil {
		return "", err
	}
	return CommonInitialisms(res.String()), nil
}

// ReadSql 读sql
func ReadSql(pkName, sql string) Table {
	tb := Table{Package: pkName}
	tb.Fields = make([]Field, 0)

	lines := strings.Split(sql, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		// 如果当前行存在CREATE TABLE 字段
		if strings.Contains(line, "CREATE TABLE") {
			tableName := strings.Split(line, "`")[1]
			tb.TableName = tableName
			tb.Name = HumpNaming(tableName)
		}
		if strings.Contains(line, "ENGINE") && strings.Contains(line, "COMMENT=") {
			tb.Comment = strings.Trim(strings.Split(line, "COMMENT='")[1], "'")
			continue
		}
		if line[0] == '`' {
			field := Field{}
			fieldName := strings.Split(line, "`")[1]
			field.FieldName = fieldName
			field.Name = HumpNaming(fieldName)
			// 字段类型判定 Start
			fieldType := strings.TrimRightFunc(strings.Split(line, " ")[1], func(r rune) bool {
				return r < 'a' || r > 'z'
			})
			var tp_success = false
			if strings.Contains(line, "UNSIGNED") || strings.Contains(line, "unsigned") {
				if ft, ok := MySqlUIntType[fieldType]; ok {
					fieldType = ft
					tp_success = true
				}
			}
			if !tp_success {
				tf := MySql2Type[fieldType]
				fieldType = tf
			}
			field.FieldType = fieldType

			if strings.Contains(fieldType, "time") {
				tb.ContainsTimeField = true
			}
			// 字段类型判定 End
			// 字段注释
			if strings.Contains(line, "COMMENT") {
				field.Comment = strings.Trim(strings.Split(line, "COMMENT '")[1], "',")
			}
			// 字段默认值
			if strings.Contains(line, "DEFAULT'") {
				field.DefaultVal = strings.Split(line, "DEFAULT ")[1]
			}
			// 是否为主键 - 如果主键同时判定是否自增
			if strings.Contains(line, "PRIMARY KEY") {
				field.IsPk = true
				if strings.Contains(line, "AUTO_INCREMENT") {
					field.IsAuto = true
				}
			}
			// 是否非空字段
			if strings.Contains(line, "NOT NULL") {
				field.IsNotNull = true
			}
			tb.Fields = append(tb.Fields, field)
		}
	}
	return tb
}

func GetSQL(tableNames ...string) ([]string, error) {
	var result []string

	for _, tableName := range tableNames {
		rows, err := DB.Query("show create table " + tableName)
		if err != nil {
			return nil, err
		}

		if rows.Next() {
			var r string
			err := rows.Scan(&tableName, &r)
			if err != nil {
				return nil, err
			}
			result = append(result, r)
		}
	}
	return result, nil
}
