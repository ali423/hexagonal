package app

import (
	"fmt"
	"github.com/ali423/hexagonal/cmd/shotener/config"
	api2 "github.com/ali423/hexagonal/internal/api"
	"github.com/ali423/hexagonal/internal/infra/db"
	log3 "github.com/ali423/hexagonal/internal/infra/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
	"gorm.io/gorm"
	log2 "log"
)

func SetupLogger(c *config.Config) log3.Logger {
	return log3.NewZapLog(c)
}

func SetupDB(c *config.Config) (*db.DB, error) {
	d := db.NewDB(&db.Config{
		DBName:     c.DBName,
		DBUsername: c.DBUsername,
		DBPassword: c.DBPassword,
		DBHostname: c.DBHostname,
		DBPort:     c.DBPort,
	})

	if c.DBType == config.DBMySql {
		err := d.SetupMySQL()
		if err != nil {
			return nil, err
		}
	} else if c.DBType == config.DBPostgres {
		err := d.SetupPG()
		if err != nil {
			return nil, err
		}
	} else if c.DBType == config.DBSqlite {
		err := d.SetupSQLite()
		if err != nil {
			return nil, err
		}
	}
	return d, nil
}

func SetupRouter(c *config.Config, db *gorm.DB) {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	docs.SwaggerInfo.BasePath = "/api/v1"
	url := ginSwagger.URL(fmt.Sprintf("http://%v:%v/swagger/doc.json", c.AppAddress, c.AppPort)) // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// ROUTES
	api2.RegisterAllRoutes(v1,
		api2.NewShortenerAPIs(db),
	)

	srvAdd := fmt.Sprintf("%s:%d", c.AppAddress, c.AppPort)
	log2.Fatal(router.Run(srvAdd))
}
