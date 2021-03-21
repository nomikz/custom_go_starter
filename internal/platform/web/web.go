package web

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

type App struct {
	mux *chi.Mux
	log *log.Logger
}

func NewApp(log *log.Logger) *App {
	return &App{
		mux: chi.NewRouter(),
		log: log,
	}
}

func (a *App) Handle(method, pattern string, h Handler) {

	fn := func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			a.log.Printf("ERROR : %s", err)

			if err := RespondError(w, err); err != nil {
				a.log.Printf("ERROR : %s", err)
			}
		}
	}

	a.mux.MethodFunc(method, pattern, fn)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
