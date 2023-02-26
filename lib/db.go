package lib

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

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
