package router

import (
	"github.com/gin-gonic/gin"
)

var routers []Router

func init() {
	routers = make([]Router, 0)
	AddRouters(PingRouter)
}

// Router
type Router interface {
	Execute(c *gin.Engine) // execute router
}

// AddRouters
func AddRouters(rou ...Router) {
	routers = append(routers, rou...)
}

// RunRouter
func RunRouter(r *gin.Engine) {
	for _, route := range routers {
		route.Execute(r)
	}
}
