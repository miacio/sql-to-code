package lib

var (
	MySql2Type = map[string]string{
		"bool": "bool",

		"char":       "string",
		"varchar":    "string",
		"tinytext":   "string",
		"text":       "string",
		"mediumtext": "string",
		"longtext":   "string",
		"json":       "string",

		"tinyint":  "int8",
		"smallint": "int64",
		"int":      "int",
		"integer":  "int",
		"bigint":   "int64",

		"float":   "float32",
		"double":  "float64",
		"numeric": "float64",
		"decimal": "float64",

		"date":      "time.Time",
		"time":      "time.Time",
		"datetime":  "time.Time",
		"timestamp": "time.Time",
	}

	MySqlUIntType = map[string]string{
		"tinyint":  "uint8",
		"smallint": "uint16",
		"int":      "uint",
		"integer":  "uint",
		"bigint":   "uint64",
	}
)
