package app

import "github.com/gin-gonic/gin"

type sqlToolsApp interface {
	Generter(c *gin.Context) // 生成结构体操作
}

type sqlToolsAppImpl struct{}

var SQLToolsApp sqlToolsApp = (*sqlToolsAppImpl)(nil)

// 生成结构体操作
func (*sqlToolsAppImpl) Generter(c *gin.Context) {

}
