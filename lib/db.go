package lib

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

// GetSQL 依据表名获取建表语句
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

// GetTableNames 获取表名列表
func GetTableNames() ([]string, error) {
	result := make([]string, 0)
	rows, err := DB.Query("show tables")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}
		result = append(result, tableName)
	}
	return result, nil
}
