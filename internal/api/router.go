package api

import (
	"github.com/ali423/hexagonal/internal/api/rest"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShortenerAPIs struct {
	redirectHandler *rest.RedirectHandler
}

func NewShortenerAPIs(db *gorm.DB) *ShortenerAPIs {
	return &ShortenerAPIs{
		redirectHandler: rest.NewRedirectHandler(db),
	}
}

func (s *ShortenerAPIs) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/swagger", func(c *gin.Context) {
		c.File("./swagger.yml")
	})

	redirect := rg.Group("/redirect")
	{
		redirect.POST("/create", s.redirectHandler.CreateRedirect)
		redirect.GET("/:code", s.redirectHandler.GetRedirect)
	}
}
