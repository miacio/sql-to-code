package lib

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/miacio/sql-to-code/log"
	"github.com/spf13/viper"
)

// c_db 数据库配置[db]
type c_db struct {
	Host     string `toml:"host"`     // 数据库地址
	Port     uint   `toml:"port"`     // 数据库端口
	User     string `toml:"user"`     // 数据库用户名
	Password string `toml:"password"` // 数据库密码
	Database string `toml:"database"` // 数据库名
	Charset  string `toml:"charset"`  // 数据库连接字符类型
}

// c_cfg 参数配置[cfg]
type c_cfg struct {
	OutDir           string   `toml:"outDir"`           // 代码输出目录
	TableNames       []string `toml:"tableName"`        // 数据库表名
	PackageName      string   `toml:"packageName"`      // 输出代码的库名
	NeedTag          []string `toml:"needTag"`          // 生成对应的标签 gorm 标签不参与下边的首字母是否大写及是否驼峰
	UpperFirstLetter bool     `toml:"upperFirstLetter"` // 生成的标签值首字母是否大写
	HumpNaming       bool     `toml:"humpNaming"`       // 生成标签值是否使用驼峰命名
	ImportOtherType  string   `toml:"importOtherType"`  // 引用其它类型(文件地址)
	WebServer        bool     `toml:"webServer"`        // 是否以web服务模式启动
}

// c_application 服务项配置[application]
// TODO: 服务器模式参数配置
type c_application struct {
	Port     uint   `toml:"port"`     // 服务端口
	UseHttps bool   `toml:"useHttps"` // 是否使用HTTPS模式启动服务
	PemFile  string `toml:"pemFile"`  // pem文件地址
	KeyFile  string `toml:"keyFile"`  // key文件地址
}

// FieldOtherType 字段其它类型引入
type FieldOtherType struct {
	ImportPath string `json:"importPath" toml:"importPath"` // 引用的包地址
	FieldType  string `json:"fieldType" toml:"fieldType"`   // 代码中字段类型
	DbType     string `json:"dbType" toml:"dbType"`         // 数据库对应的类型
}

var (
	dbParam         c_db
	CfgParam        c_cfg
	Application     c_application
	FieldOtherTypes []FieldOtherType
)

func logger() {
	lo := map[string]log.Level{
		"debug.log": log.DebugLevel,
		"info.log":  log.InfoLevel,
		"error.log": log.ErrorLevel,
	}
	log.Init("./out", 256, 10, 7, false, lo)
}

// 全局lib启动器操作初始化文件等
func init() {
	logger()

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("toml")
	runPath, _ := os.Getwd()
	v.AddConfigPath(runPath)

	err := v.ReadInConfig()
	if err != nil {
		log.GetLogger().Errorf("read config fail: %v", err)
		return
	}

	if err := v.UnmarshalKey("cfg", &CfgParam); err != nil {
		log.GetLogger().Errorf("viper read cfg param fail: %v", err)
		return
	}

	// 是否以Web服务模式启动, 如果是则拉取application参数, 并关闭数据库连接模式及其他参数配置模式
	if CfgParam.WebServer {
		if err := v.UnmarshalKey("application", &Application); err != nil {
			log.GetLogger().Errorf("viper read application param fail: %v", err)
		}
		return
	}

	// dbParam 数据库配置项
	if err := v.UnmarshalKey("db", &dbParam); err != nil {
		log.GetLogger().Errorf("viper read db param fail: %v", err)
		return
	}

	// 外部应用项
	if CfgParam.ImportOtherType != "" {
		importOtherTypeFile, err := os.ReadFile(CfgParam.ImportOtherType)
		if err != nil {
			log.GetLogger().Errorf("viper read cfg param importOtherType file fail: %v", err)
			return
		}
		if err := json.Unmarshal(importOtherTypeFile, &FieldOtherTypes); err != nil {
			log.GetLogger().Errorf("viper read cfg param importOtherType file fail: %v", err)
			return
		}
	}

	// 数据库连接项
	if err := dbParam.LinkDB(); err != nil {
		log.GetLogger().Errorf("db link fail: %v", err)
		return
	}
}

// GetDSN
func (c *c_db) GetDSN(parseTime, loc string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s", c.User, c.Password, c.Host, c.Port, c.Database, c.Charset, parseTime, loc)
}

// LinkDB
func (c *c_db) LinkDB() error {
	dsn := c.GetDSN("True", "Local")
	pool, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	pool.SetMaxOpenConns(64)
	pool.SetMaxIdleConns(16)
	pool.SetConnMaxLifetime(100 * time.Second)
	DB = pool
	return nil
}
