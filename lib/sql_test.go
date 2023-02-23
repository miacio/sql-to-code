package lib_test

import (
	"fmt"
	"os"
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

func TestReadSQL(t *testing.T) {
	sql, err := os.ReadFile("../testfile/load.input")
	if err != nil {
		t.Fatal("read testfile/load.input failed", err)
	}

	err = lib.GenerateCodeFile("../model", "model", string(sql))
	if err != nil {
		t.Fatal("generate code file failed", err)
	}

}
