package api

import "github.com/gin-gonic/gin"

type Route interface {
	// RegisterRoutes receives a gin router and registers routes of a resource.
	// gin is the only supported router at this moment. We may support other routers later.
	RegisterRoutes(rg *gin.RouterGroup)
}

func RegisterAllRoutes(rg *gin.RouterGroup, rs ...Route) {
	for _, r := range rs {
		r.RegisterRoutes(rg)
	}
}
