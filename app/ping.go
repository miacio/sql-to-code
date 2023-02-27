package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type pingApp interface {
	Pong(c *gin.Context)
}

type pingAppImpl struct {
}

var PingApp pingApp = (*pingAppImpl)(nil)

func (*pingAppImpl) Pong(c *gin.Context) {
	c.JSONP(http.StatusOK, gin.H{"code": "200", "msg": "pong"})
}
