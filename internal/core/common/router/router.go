package router

import "github.com/gin-gonic/gin"

func Post(r *gin.RouterGroup, path string, handler func(c *gin.Context)) {
	r.POST(path, handler)
}
