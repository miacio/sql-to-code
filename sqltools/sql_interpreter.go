package sqltools

import (
	"fmt"
	"regexp"
	"strings"
)

// GetByPattern 按模板获取数据
func GetByPattern(val, patt string) []string {
	reg := regexp.MustCompile(patt)
	res := reg.FindStringSubmatch(val)
	return res
}

// GetByIndexPattern 按模板获取指定下标位数据
func GetByIndexPattern(val, patt string, index int) string {
	vals := GetByPattern(val, patt)
	if len(vals) >= index {
		return vals[index]
	}
	return ""
}

// GetTableName 获取表名
func GetTableName(line string) string {
	result := GetByIndexPattern(line, "(?i)CREATE TABLE `(.*)`", 1)
	if result == "" {
		result = strings.Split(GetByIndexPattern(line, "(?i)CREATE TABLE (.*)", 1), " ")[0]
	}
	return result
}

// GetTableComment 获取表注释
func GetTableComment(line string) string {
	result := GetByIndexPattern(line, "(?i)\\).*COMMENT='(.*)'", 1)
	if result == "" {
		result = GetByIndexPattern(line, "(?i)\\).*COMMENT = '(.*)'", 1)
	}
	return result
}

// GetPrimaryKey 获取表主键
func GetPrimaryKey(line string) string {
	return GetByIndexPattern(line, "(?i)PRIMARY KEY \\(`(.*)`\\)", 1)
}

// GetFieldSql 获取字段sql
func GetFieldSql(sql string) []string {
	result := make([]string, 0)
	lines := strings.Split(sql, "\n")
	start := false
	for _, line := range lines {
		upLine := strings.ToUpper(line)
		if strings.Contains(upLine, "CREATE TABLE ") {
			start = true
			continue
		}
		if strings.Contains(upLine, "PRIMARY KEY (`") || strings.Contains(upLine, "ENGINE =") || strings.Contains(upLine, "ENGINE=") {
			start = false
		}
		if start {
			result = append(result, line)
		}
	}
	return result
}

// FieldSqlToStruct 字段sql转结构体对象
func FieldSqlToStruct(line string) Field {
	line = strings.TrimSpace(line)
	var fieldName, fieldType, comment, defaultVal string
	var isNotNull, isUnsigned, isPrimaryKey, isAutoIncrement, isAutoCreateTime, isAutoUpdateTime bool

	if line[0] == '`' {
		fieldName = GetByIndexPattern(line, "`(.*)`", 1)
		fieldType = GetByIndexPattern(line, fmt.Sprintf("`%s` ([A-Za-z]*)", fieldName), 1)
	} else {
		fieldName = strings.Split(line, " ")[0]
		fieldType = GetByIndexPattern(line, fmt.Sprintf("%s ([A-Za-z]*)", fieldName), 1)
	}

	comment = GetByIndexPattern(line, "(?i)COMMENT '(.*)'", 1)
	defaultVal = strings.Split(GetByIndexPattern(line, "(?i)DEFAULT (.*) ", 1), " ")[0]
	if defaultVal == "NULL" {
		defaultVal = ""
	}

	upLine := strings.ToUpper(line)
	if strings.Contains(upLine, "UNSIGNED") {
		isUnsigned = true
	}
	if strings.Contains(upLine, "NOT NULL") {
		isNotNull = true
	}
	if strings.Contains(upLine, "PRIMARY KEY") {
		isPrimaryKey = true
	}

	if strings.Contains(upLine, "AUTO_INCREMENT") {
		isAutoIncrement = true
	}

	if strings.Contains(upLine, "UPDATE CURRENT_TIMESTAMP") {
		defaultVal = ""
		isAutoUpdateTime = true
	} else if strings.Contains(upLine, "CURRENT_TIMESTAMP") {
		defaultVal = ""
		isAutoCreateTime = true
	}

	return Field{
		IsUnsigned:       isUnsigned,
		IsNotNull:        isNotNull,
		IsPrimaryKey:     isPrimaryKey,
		IsAutoIncrement:  isAutoIncrement,
		IsAutoCreateTime: isAutoCreateTime,
		IsAutoUpdateTime: isAutoUpdateTime,

		FieldName:  fieldName,
		FieldType:  fieldType,
		Comment:    comment,
		DefaultVal: defaultVal,
	}
}
