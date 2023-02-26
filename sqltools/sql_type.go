package sqltools

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

	BaseType = map[string]struct{}{
		"bool":    {},
		"string":  {},
		"int":     {},
		"int8":    {},
		"int16":   {},
		"int32":   {},
		"int64":   {},
		"float32": {},
		"float64": {},
		"uint":    {},
		"uint8":   {},
		"uint16":  {},
		"uint32":  {},
		"uint64":  {},
	}
)
