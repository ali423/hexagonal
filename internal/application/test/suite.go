package test

import (
	"fmt"
	"github.com/ali423/hexagonal/cmd/shotener/app"
	"github.com/ali423/hexagonal/cmd/shotener/config"
	"github.com/ali423/hexagonal/internal/infra/db"
	log3 "github.com/ali423/hexagonal/internal/infra/logger"
	"gorm.io/gorm"
	"sync"
)

func SetupTestSuite() {
	(&sync.Once{}).Do(func() {

		c := &config.Config{}

		d := db.NewDB(&db.Config{DBName: "file::memory:?cache=shared"})
		if err := d.SetupSQLite(); err != nil {
			panic(err)
		}

		l := log3.NewZapLog(&config.Config{Writers: "stdout"})

		r := db.NewDefaultRepository(d)
		fmt.Println(c, l, r)

		app.NewApp(c, l, r, d).SetApp()
	})
}

func ClearTable(model interface{}) {
	stmt := &gorm.Statement{DB: app.A().GetDB().DB}
	err := stmt.Parse(model)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table cleared: ", stmt.Schema.Table)
	app.A().GetDB().DB.Exec(fmt.Sprintf("DELETE FROM %s;", stmt.Schema.Table))
}

func ClearTableByList(models []interface{}) {
	for _, model := range models {
		ClearTable(model)
	}
}
