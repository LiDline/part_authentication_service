package healthcheck

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HealthCheckRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	return r
}
