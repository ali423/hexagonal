package application

import (
	"github.com/ali423/hexagonal/internal/repository/mysql"
	"gorm.io/gorm"
)

type Infrastructure struct {
	RedirectRepository *mysql.RedirectRepository
}

func NewInfrastructure(db *gorm.DB) Infrastructure {
	return Infrastructure{
		RedirectRepository: mysql.NewRedirectRepository(db),
	}
}
