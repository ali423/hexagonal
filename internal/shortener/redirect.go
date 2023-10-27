package shortener

type Redirect struct {
	Code string
	Url  string
}

type RedirectRepositoryInterface interface {
	Find(code string) (redirect *Redirect, err error)
	Store(redirect *Redirect) error
}
