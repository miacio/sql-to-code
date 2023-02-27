package main_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/miacio/sql-to-code/lib"
	"github.com/miacio/sql-to-code/sqltools"
)

func TestReadSql(t *testing.T) {
	sql, err := os.ReadFile("./testfile/load.input")
	if err != nil {
		t.Fatal("read testfile/load.input failed", err)
	}

	var fieldOtherTypes []lib.FieldOtherType
	importOtherTypeFile, err := os.ReadFile("./testfile/fieldOtherType.json")
	if err != nil {
		t.Fatal("read param importOtherType file fail:", err)
	}
	if err := json.Unmarshal(importOtherTypeFile, &fieldOtherTypes); err != nil {
		t.Fatal("importOtherType to json fail:", err)
	}

	needTag := []string{"gorm", "json", "toml"}
	upperFirstLetter := false
	humpNaming := true

	dir := "model"
	runPath, _ := os.Getwd()
	outFilePath := path.Join(runPath, dir)
	if err := sqltools.GenerateCodeFile(outFilePath, "model", string(sql), needTag, upperFirstLetter, humpNaming, fieldOtherTypes); err != nil {
		t.Fatal("generate code file fail:", err)
	}
}

func TestGetTables(t *testing.T) {
	result, err := lib.GetTableNames()
	if err != nil {
		fmt.Printf("get table names fail: %v", err)
		return
	}
	for _, res := range result {
		fmt.Println(res)
	}
}
