package rest

import (
	"github.com/ali423/hexagonal/internal/application"
	"github.com/ali423/hexagonal/internal/shortener"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func NewRedirectHandler(db *gorm.DB) *RedirectHandler {
	return &RedirectHandler{
		RedirectService: application.NewRedirectService(db),
	}
}

type RedirectHandlerInterface interface {
	CreateRedirect(*gin.Context)
	GetRedirect(*gin.Context)
}

type RedirectHandler struct {
	RedirectService application.RedirectServiceInterface
}

type CreateRedirectRequest struct {
	Code string `json:"code" binding:"required"`
	Url  string `json:"url" binding:"required"`
}

type RedirectResponse struct {
	Code string `json:"Code"`
	Url  string `json:"Url"`
}

func CreateRedirectRequestToRequestStruct(req CreateRedirectRequest) shortener.Redirect {
	return shortener.Redirect{
		Code: req.Code,
		Url:  req.Url,
	}
}

func RedirectToResponse(redirect shortener.Redirect) RedirectResponse {
	return RedirectResponse{
		Code: redirect.Code,
		Url:  redirect.Url,
	}
}

func (r *RedirectHandler) CreateRedirect(c *gin.Context) {

	var newRedirect CreateRedirectRequest
	if err := c.ShouldBindJSON(&newRedirect); err != nil {
		ResponseFail(c, http.StatusBadRequest, err)
		return
	}
	redirect := CreateRedirectRequestToRequestStruct(newRedirect)
	if err := r.RedirectService.Store(&redirect); err != nil {
		ResponseFail(c, http.StatusBadRequest, err)
		return
	}
	ResponseSuccess(c, http.StatusCreated, "redirect link Created")
}

func (r *RedirectHandler) GetRedirect(c *gin.Context) {

	code := c.Param("code")
	redirect, err := r.RedirectService.Find(code)
	if err != nil {
		ResponseFail(c, http.StatusBadRequest, err)
		return
	}
	myResponse := RedirectToResponse(*redirect)

	ResponseSuccess(c, http.StatusCreated, myResponse)
}
