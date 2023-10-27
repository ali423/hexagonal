package app

import (
	"github.com/ali423/hexagonal/cmd/shotener/config"
	"github.com/ali423/hexagonal/internal/infra/db"
	"github.com/ali423/hexagonal/internal/infra/logger"
	"sync"
	"time"
)

var once sync.Once
var app *App

type App struct {
	Config     *config.Config
	Logger     logger.Logger
	Repository db.Repository
	DB         *db.DB
}

func NewApp(c *config.Config, l logger.Logger, repo db.Repository, d *db.DB) *App {
	return &App{
		Config:     c,
		Logger:     l,
		Repository: repo,
		DB:         d,
	}
}

// SetApp fills the current instance of App for further
func (a *App) SetApp() {
	once.Do(func() {
		app = a
	})
}

func (a *App) GetDB() *db.DB {
	return a.DB
}

// A returns a previously created App instance
func A() *App {
	time.Local, _ = time.LoadLocation("Asia/Tehran")
	return app
}
