package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/ginframe/internel/handler"
)

// ApiRouter http api routers
func ApiRouter(r *gin.Engine) {

	r.GET("/ping", handler.Ping)

}
