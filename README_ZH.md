# sql-to-code
(GO语言) 读数据库生成结构体代码工具

这个工具对比其它使用的数据库sql生成代码工具厉害在于其可配置其它数据库类型

比如point类型, 其它工具使用时是无法生成的,但是此工具可以配置自定义类型,让你的代码生成更快捷,更方便

# 准备
使用此工具前需要安装goimports指令
```
go get -v golang.org/x/tools/cmd/goimports

go install golang.org/x/tools/cmd/goimports
```

设置配置文件 (config.toml)
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
db模块配置既数据库模块配置:

host = 数据库地址

port = 数据库端口号

user = 数据库访问用户名

password = 数据库访问密码

database = 数据库库名

cfg模块配置为生成代码方面的配置:

outDir = 代码输出的文件夹路径

tableNames = [对应读的表名数组]

packageName = 输出的代码属于哪一个包

needTag = [需要生成的tag名称数组]

upperFirstLetter = 布尔型值, 生成的tag值是否首字母大写, 这个参数对于gorm类型的tag不起作用

humpNaming = 布尔型值, 生成的tag值是否使用驼峰命名规则, 这个参数对于gorm类型的tag不起作用

importOtherType = 引入的其他类型的配置文件地址

### 引入的其他类型的配置文件说明
其他类型配置是一个json数组

例如sqltools/sqltools_test.go测试方法

testfile/fieldOtherType.json为例:

```
[{
    "importPath": "",
    "fieldType": "IPoint",
    "dbType": "point"
}]
```

其他类型配置文件中的内容是一个json数组,这个json数组可以指定此类型需导入的包, 代码中使用的类型名称, 数据库对应的类型名称

依据此配置即可让工具理解你的意图,将此类型转换为你所指定的类型

# 运行
如果你是直接拉取的此项目,那么你只需要配置你对应的config.toml后使用 go run /.main.go 即可运行

如果你需要将此工具交给无环境的用户,那么只需要build此项目后,将对应的配置文件与生成的工具同目录,然后配置你的配置文件即可

### 参考文档
https://github.com/gangming/sql2struct

http://www.javashuo.com/article/p-eraqqmsn-a.html
