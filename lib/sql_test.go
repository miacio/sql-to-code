package lib_test

import (
	"fmt"
	"testing"

	"github.com/miacio/sql-to-code/lib"
)

func TestTo(t *testing.T) {
	tb := lib.Table{
		Package:   "model",
		Name:      "User",
		TableName: "user",
		Comment:   "用户表",
		Fields: []lib.Field{
			{
				IsPk:      true,
				IsAuto:    true,
				Name:      "Id",
				FieldName: "id",
				FieldType: "int",
				Comment:   "自增长",
			},
			{
				Name:      "Name",
				FieldName: "name",
				FieldType: "string",
				Comment:   "用户名",
			},
			{
				Name:      "Age",
				FieldName: "age",
				FieldType: "int",
				Comment:   "年龄",
			},
			{
				Name:      "CreateBy",
				FieldName: "create_by",
				FieldType: "time.Time",
			},
		},
		ContainsTimeField: true,
	}
	val, err := tb.ToCode()
	if err != nil {
		fmt.Printf("table to code fail: %v", err)
		return
	}
	fmt.Println(val)
}
