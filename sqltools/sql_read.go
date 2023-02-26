package sqltools

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/miacio/sql-to-code/lib"
)

// GenerateTable 生成表结构体
// packageName 结构体生成后所在的包名
// sql 建表语句
// needTag 需要生成的字段标签 gorm标签与后置配置无关
// upperFirstLetter 字段标签中名称首字母是否大写
// humpNaming 字段标签中是否使用驼峰命名
func GenerateTable(packageName, sql string, needTag []string, upperFirstLetter, humpNaming bool, fieldOtherTypes []FieldOtherType) Table {
	sqlLines := make([]string, 0)
	for _, sqlLine := range strings.Split(sql, "\n") {
		if strings.TrimSpace(sqlLine) != "" {
			sqlLines = append(sqlLines, sqlLine)
		}
	}
	sql = strings.Join(sqlLines, "\n")
	tableName := GetTableName(string(sql))
	tableComment := GetTableComment(string(sql))
	parimaryKey := GetPrimaryKey(string(sql))

	table := Table{
		Package:   packageName,
		Name:      lib.HumpNaming(tableName),
		TableName: tableName,
		Comment:   tableComment,
		Fields:    make([]Field, 0),
	}

	imports := make([]string, 0)
	importMap := make(map[string]struct{}, 0)
	lines := GetFieldSql(string(sql))
	for _, line := range lines {
		field := FieldSqlToStruct(line)
		if parimaryKey != "" && field.FieldName == parimaryKey {
			field.IsPrimaryKey = true
		}
		field.Name = lib.HumpNaming(field.FieldName)
		if field.IsUnsigned {
			field.Type = MySqlUIntType[field.FieldType]
		} else {
			field.Type = MySql2Type[field.FieldType]
		}
		if field.Type == "" {
			for _, fieldOtherType := range fieldOtherTypes {
				if fieldOtherType.DbType == field.FieldType {
					field.Type = fieldOtherType.FieldType

					if fieldOtherType.ImportPath != "" {
						if _, ok := importMap[field.Type]; !ok {
							imports = append(imports, fieldOtherType.ImportPath)
							importMap[field.Type] = struct{}{}
						}
					}
				}
			}
		}

		if _, ok := importMap[field.Type]; strings.Contains(field.Type, "time") && !ok {
			imports = append(imports, "time")
			importMap[field.Type] = struct{}{}
		}

		tags := make([]string, 0)
		for _, need := range needTag {
			switch need {
			case "gorm":
				gormTag := field.GetGormTag()
				tags = append(tags, gormTag)
			default:
				tag := field.GenerateTag(need, upperFirstLetter, humpNaming)
				tags = append(tags, tag)
			}
		}
		field.FieldTag = "`" + strings.Join(tags, " ") + "`"
		table.Fields = append(table.Fields, field)
	}
	table.Imports = imports
	return table
}

// GenerateCodeFile 生成代码文件
func GenerateCodeFile(dir, pkName, sql string, needTag []string, upperFirstLetter, humpNaming bool, fieldOtherTypes []FieldOtherType) error {
	table := GenerateTable(pkName, sql, needTag, upperFirstLetter, humpNaming, fieldOtherTypes)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	code, err := table.ToCode()
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir, table.TableName+".go")
	fd, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer fd.Close()
	_, err = fd.Write([]byte(code))
	if err != nil {
		return err
	}
	_, err = exec.Command("goimports", "-l", "-w", dir).Output()
	if err != nil {
		return err
	}
	_, err = exec.Command("gofmt", "-l", "-w", dir).Output()
	if err != nil {
		return err
	}
	return nil
}
