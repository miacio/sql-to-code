package router

import "github.com/gin-gonic/gin"

type sqlToolsRouter struct{}

var SQLToolsRouter Router = (*sqlToolsRouter)(nil)

func (*sqlToolsRouter) Execute(c *gin.Engine) {
	sqltools := c.Group("/sqltools")
	{
		sqltools.POST("/generter")
	}
}
