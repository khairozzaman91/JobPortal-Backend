package middlewares

import "net/http"

type Middlewares func(http.Handler) http.Handler

type Manager struct {
	middlewares []Middlewares
}

func NewManager() *Manager {
	return &Manager{
		middlewares: make([]Middlewares, 0),
	}
}

func (m *Manager) Use(mws ...Middlewares) {
	m.middlewares = append(m.middlewares, mws...)
}

func (m *Manager) Wraper(next http.Handler) http.Handler {
	n := next
	for i := len(m.middlewares) - 1; i >= 0; i-- {
		middleware := m.middlewares[i]
		n = middleware(n)
	}
	return n
}

func (m *Manager) With(handler http.Handler, middlewares ...Middlewares) http.Handler {
	n := handler
	for i := len(middlewares) - 1; i >= 0; i-- {
		middleware := middlewares[i]

		n = middleware(n)
	}
	return n
}
