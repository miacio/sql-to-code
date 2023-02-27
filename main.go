package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/miacio/sql-to-code/app"
	"github.com/miacio/sql-to-code/lib"
	"github.com/miacio/sql-to-code/log"
	"github.com/miacio/sql-to-code/sqltools"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.GetLogger().Errorf("system panic msg: %v", r)
		}
	}()

	if lib.CfgParam.WebServer {
		webServer()
	} else {
		generateCode()
	}
}

func webServer() {
	r := gin.Default()
	r.GET("/", app.PingApp.Pong)
	r.GET("/ping", app.PingApp.Pong)

	webPort := fmt.Sprintf(":%d", lib.Application.Port)
	if lib.Application.UseHttps {
		r.Use(app.TlsHandler(webPort))
		if err := r.RunTLS(webPort, lib.Application.PemFile, lib.Application.KeyFile); err != nil {
			log.GetLogger().Errorf("Application server RunTLS fail: %v", err)
			return
		}
	}
	if err := r.Run(webPort); err != nil {
		log.GetLogger().Errorf("Application server Run fail: %v", err)
		return
	}
}

func generateCode() {
	sqls, err := lib.GetSQL(lib.CfgParam.TableNames...)
	if err != nil {
		log.GetLogger().Errorf("get db sql fail: %v", err)
		return
	}

	for _, sql := range sqls {
		err := sqltools.GenerateCodeFile(lib.CfgParam.OutDir, lib.CfgParam.PackageName, sql, lib.CfgParam.NeedTag, lib.CfgParam.UpperFirstLetter, lib.CfgParam.HumpNaming, lib.FieldOtherTypes)
		if err != nil {
			log.GetLogger().Error("generate code file fail: %v", err)
			continue
		}
	}
}
