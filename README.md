# sql-to-code
go read db table to code tool [中文](https://github.com/miacio/sql-to-code/blob/master/README_ZH.md)

# prepare
install goimports
```
go get -v golang.org/x/tools/cmd/goimports

go install golang.org/x/tools/cmd/goimports
```

set config.toml file
```
[db]
host="127.0.0.1"
port=3306
user="root"
password="123456"
database="test"

[cfg]
outDir="../model"
tableNames=["user"]
packageName="model"
needTag=["gorm", "json"]
upperFirstLetter=false
humpNaming=false
importOtherType="./fieldOtherType.json"
```

↑↑↑
outDir is generate go code folder path

tableNames is the user needs to configure the table name array that needs to generate the go file

packageName is generate go code belong to package

needTag is generate go code tag array

upperFirstLetter is set tag value first letter is upper, gorm is not in this upperFirstLetter

humpNaming is set tag value use hump naming func, gorm is not in this humpNaming

importOtherType is import other database type config, see test: sqltools/sqltools_test.go func

### (importOtherType.json) - this json is array

object param importPath is import other package url

fieldType is other package param

dbType is in database type

```
[{
    "importPath": "",
    "fieldType": "IPoint",
    "dbType": "point"
}]
```

# run
code run is go run ./main.go

exe than go build

### document
https://github.com/gangming/sql2struct

http://www.javashuo.com/article/p-eraqqmsn-a.html