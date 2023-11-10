package api

import "github.com/gin-gonic/gin"

type Route interface {
	RegisterRoutes(rg *gin.RouterGroup)
}

func RegisterAllRoutes(rg *gin.RouterGroup, rs ...Route) {
	for _, r := range rs {
		r.RegisterRoutes(rg)
	}
}
