package lib

import (
	"fmt"
	"strings"
)

func (field *Field) GetGormTag() string {
	gormTag := make([]string, 0)
	gormTag = append(gormTag, "column:"+field.FieldName)

	if field.IsPk {
		gormTag = append(gormTag, "primaryKey")
		// 判断是否为自增长 (仅主键字段可用)
		if field.IsAuto {
			gormTag = append(gormTag, "autoIncrement")
		}
	}
	if field.IsUnique {
		gormTag = append(gormTag, "unique")
	}
	if field.IsNotNull {
		gormTag = append(gormTag, "not null")
	}
	if field.IsIndex {
		gormTag = append(gormTag, "index")
	}
	if field.IsUniqueIndex {
		gormTag = append(gormTag, "uniqueIndex")
	}
	if field.DefaultVal != "" {
		gormTag = append(gormTag, "default:"+field.DefaultVal)
	}
	return fmt.Sprintf("gorm:\"%s\"", strings.Join(gormTag, ";"))
}
