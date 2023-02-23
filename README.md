# sql-to-code
go read db table to code

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
```

↑↑↑
outDir is generate go code folder path

tableNames is the user needs to configure the table name array that needs to generate the go file

packageName is generate go code belong to package

# run
code run is go run ./main.go

exe than go build

### document
github.com/gangming/sql2struct