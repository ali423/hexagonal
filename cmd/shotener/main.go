package main

import (
	"github.com/ali423/hexagonal/cmd/shotener/app"
	"github.com/ali423/hexagonal/cmd/shotener/config"
	"github.com/ali423/hexagonal/internal/infra/db"
)

func main() {
	cl := config.ViperLoader{}
	c, err := cl.Load()
	if err != nil {
		panic(err)
	}
	d, err := app.SetupDB(c)
	if err != nil {
		panic(err)
	}

	l := app.SetupLogger(c)
	r := db.NewDefaultRepository(d)
	app.NewApp(c, l, r, d).SetApp()

	app.SetupRouter(c, d.DB)
}
