package middleware

import "github.com/gorilla/mux"

type Manager struct {
	router *mux.Router
}

func NewManager(router *mux.Router) *Manager {
	return &Manager{
		router: router,
	}
}

func (m Manager) Apply() {
	m.router.Use(mux.CORSMethodMiddleware(m.router))

	m.router.Use(panicRecovery)
}
