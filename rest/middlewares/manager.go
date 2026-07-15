package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	middlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		middlewares: make([]Middleware, 0),
	}
}


func (m *Manager) Use(mws ...Middleware) {
	m.middlewares = append(m.middlewares, mws...)
}

func (m *Manager) Wraper(next http.Handler) http.Handler {
	n := next
	for i := len(m.middlewares)-1; i >= 0; i-- {
		middleware := m.middlewares[i]
		n = middleware(n)
	}
	return n
}


func (m *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	n := handler
	for i := len(middlewares)-1; i >= 0; i-- {
		middleware := middlewares[i]
		n = middleware(n)
	}
	return n
}