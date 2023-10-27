package mysql

import (
	"github.com/ali423/hexagonal/internal/shortener"
	"gorm.io/gorm"
	"time"
)

type RedirectRepository struct {
	DB *gorm.DB
}

type Redirect struct {
	Id        int       `gorm:"primary_key"`
	Code      string    `gorm:"column:code"`
	Url       string    `gorm:"column:url"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (r *Redirect) TableName() string {
	return "redirects"
}

func (r *Redirect) ModelToDto() (redirect shortener.Redirect) {
	return shortener.Redirect{
		Code: r.Code,
		Url:  r.Url,
	}
}

func NewRedirectRepository(db *gorm.DB) *RedirectRepository {
	return &RedirectRepository{
		DB: db,
	}
}

func (rr *RedirectRepository) Find(code string) (redirect *shortener.Redirect, err error) {
	var redirectModel Redirect
	err = rr.DB.First(&redirectModel, "code = ?", code).Error
	redirectDto := redirectModel.ModelToDto()
	return &redirectDto, err
}

func (rr *RedirectRepository) Store(redirect *shortener.Redirect) error {
	redirectModel := Redirect{
		Code:      redirect.Code,
		Url:       redirect.Url,
		CreatedAt: time.Now(),
	}
	return rr.DB.Create(&redirectModel).Error
}
