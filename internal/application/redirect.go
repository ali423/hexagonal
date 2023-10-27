package application

import (
	shortner "github.com/ali423/hexagonal/internal/shortener"
	"gorm.io/gorm"
)

type RedirectService struct {
	infra Infrastructure
}

func NewRedirectService(db *gorm.DB) *RedirectService {
	return &RedirectService{
		infra: NewInfrastructure(db),
	}
}

func (r *RedirectService) Find(code string) (redirect *shortner.Redirect, err error) {
	return r.infra.RedirectRepository.Find(code)
}

func (r *RedirectService) Store(redirect *shortner.Redirect) error {
	return r.infra.RedirectRepository.Store(redirect)
}
