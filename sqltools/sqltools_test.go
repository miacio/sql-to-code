package sqltools_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/miacio/sql-to-code/sqltools"
)

func TestReadSql(t *testing.T) {
	sql, err := os.ReadFile("../testfile/load.input")
	if err != nil {
		t.Fatal("read testfile/load.input failed", err)
	}

	var fieldOtherTypes []sqltools.FieldOtherType
	importOtherTypeFile, err := os.ReadFile("../testfile/fieldOtherType.json")
	if err != nil {
		t.Fatal("read param importOtherType file fail:", err)
	}
	if err := json.Unmarshal(importOtherTypeFile, &fieldOtherTypes); err != nil {
		t.Fatal("importOtherType to json fail:", err)
	}

	needTag := []string{"gorm", "json", "toml"}
	upperFirstLetter := false
	humpNaming := true
	if err := sqltools.GenerateCodeFile("../model", "model", string(sql), needTag, upperFirstLetter, humpNaming, fieldOtherTypes); err != nil {
		t.Fatal("generate code file fail:", err)
	}
}
