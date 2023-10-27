package application

import shortner "github.com/ali423/hexagonal/internal/shortener"

type RedirectServiceInterface interface {
	Find(code string) (*shortner.Redirect, error)
	Store(redirect *shortner.Redirect) error
}
