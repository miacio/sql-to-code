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
tableName="user"
packageName="model"
```