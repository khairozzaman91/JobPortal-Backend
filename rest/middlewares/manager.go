package middlewares

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalmiddlewares []Middleware
}

func NewManager() *Manager {
	mngr := Manager{
		globalmiddlewares: make([]Middleware, 0),
	}
	return &mngr
}

func (mngr *Manager) With(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {

		n := next

		for i := len(middlewares) - 1; i >= 0; i-- {
			middleware := middlewares[i]
			n = middleware(n)
		}
		return n
	}
}
