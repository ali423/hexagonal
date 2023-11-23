package test

import (
	"github.com/ali423/hexagonal/cmd/shotener/app"
	"github.com/ali423/hexagonal/internal/application"
	"github.com/ali423/hexagonal/internal/repository/mysql"
	shortner "github.com/ali423/hexagonal/internal/shortener"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"sync"
	"testing"
	"time"
)

var onceRedirectTest sync.Once

type locationTest struct{}

var locationTestInstance = &locationTest{}
var redirectCode = "google"
var redirectUrl = "http://google.com/"

var redirectCreateData = shortner.Redirect{
	Code: "test_code",
	Url:  "test_url",
}

func Test_FindRedirectByCode(t *testing.T) {
	locationTestInstance.LocationTestInit()
	var redirect *shortner.Redirect
	d := app.A().GetDB().DB

	newRedirect := NewRedirectTestHandler(d)

	redirect, err := newRedirect.RedirectService.Find(redirectCode)
	assert.Equal(t, nil, err)
	assert.Equal(t, redirect.Url, redirectUrl)
}

func Test_CreateRedirect(t *testing.T) {
	locationTestInstance.LocationTestInit()
	d := app.A().GetDB().DB
	newRedirect := NewRedirectTestHandler(d)

	err := newRedirect.RedirectService.Store(&redirectCreateData)
	assert.Equal(t, nil, err)

	var storedRedirect mysql.Redirect
	d.Table("(?) AS r", d.Model(&mysql.Redirect{})).
		Order("id desc").
		Find(&storedRedirect)
	assert.NotNil(t, storedRedirect)
	assert.Equal(t, redirectCreateData.Url, storedRedirect.Url)
	assert.Equal(t, redirectCreateData.Code, storedRedirect.Code)
}

type RedirectTestHandler struct {
	RedirectService application.RedirectServiceInterface
}

func NewRedirectTestHandler(db *gorm.DB) *RedirectTestHandler {
	return &RedirectTestHandler{
		RedirectService: application.NewRedirectService(db),
	}
}

func (l *locationTest) LocationTestInit() {
	SetupTestSuite()
	onceRedirectTest.Do(func() {
		d := app.A().GetDB()
		err := d.DB.Migrator().AutoMigrate(
			&mysql.Redirect{},
		)
		if err != nil {
			panic(err)
		}
		ClearTableByList([]interface{}{
			&mysql.Redirect{},
		})
		d.DB.Create(&mysql.Redirect{
			Id:        1,
			Code:      redirectCode,
			Url:       redirectUrl,
			CreatedAt: time.Now(),
		})
	})

}
