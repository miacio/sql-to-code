package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/miacio/sql-to-code/lib"
	"github.com/miacio/sql-to-code/log"
	"github.com/spf13/viper"
)

type c_db struct {
	Host     string `toml:"host"`     // 数据库地址
	Port     uint   `toml:"port"`     // 数据库端口
	User     string `toml:"user"`     // 数据库用户名
	Password string `toml:"password"` // 数据库密码
	Database string `toml:"database"` // 数据库名
	Charset  string `toml:"charset"`  // 数据库连接字符类型
}

type c_cfg struct {
	OutDir      string   `toml:"outDir"`      // 代码输出目录
	TableNames  []string `toml:"tableName"`   // 数据库表名
	PackageName string   `toml:"packageName"` // 输出代码的库名
}

var (
	dbParam  c_db
	cfgParam c_cfg
)

func init() {
	lo := map[string]log.Level{
		"debug.log": log.DebugLevel,
		"info.log":  log.InfoLevel,
		"error.log": log.ErrorLevel,
	}
	log.Init("./out", 256, 10, 7, false, lo)
}

// GetDSN
func (c *c_db) GetDSN(charset, parseTime, loc string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s", c.User, c.Password, c.Host, c.Port, c.Database, charset, parseTime, loc)
}

// LinkDB
func (c *c_db) LinkDB() error {
	dsn := c.GetDSN("utf8mb4", "True", "Local")
	pool, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	pool.SetMaxOpenConns(64)
	pool.SetMaxIdleConns(16)
	pool.SetConnMaxLifetime(100 * time.Second)
	lib.DB = pool
	return nil
}

func main() {
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

	if err := v.UnmarshalKey("db", &dbParam); err != nil {
		log.GetLogger().Errorf("viper read db param fail: %v", err)
		return
	}

	if err := v.UnmarshalKey("cfg", &cfgParam); err != nil {
		log.GetLogger().Errorf("viper read cfg param fail: %v", err)
		return
	}

	if err := dbParam.LinkDB(); err != nil {
		log.GetLogger().Errorf("db link fail: %v", err)
		return
	}

	sqls, err := lib.GetSQL(cfgParam.TableNames...)
	if err != nil {
		log.GetLogger().Errorf("get db sql fail: %v", err)
		return
	}
	for _, sql := range sqls {
		err := lib.GenerateCodeFile(cfgParam.OutDir, cfgParam.PackageName, sql)
		if err != nil {
			log.GetLogger().Error("generate code file fail: %v", err)
			continue
		}
	}
}
