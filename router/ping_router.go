package router

import (
	"github.com/gin-gonic/gin"
	"github.com/miacio/sql-to-code/app"
)

type pingRouter struct{}

var PingRouter Router = (*pingRouter)(nil)

func (*pingRouter) Execute(c *gin.Engine) {
	c.GET("/", app.PingApp.Pong)
	c.GET("/ping", app.PingApp.Pong)
}
