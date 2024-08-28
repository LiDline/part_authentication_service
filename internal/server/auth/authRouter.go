package auth

import (
	"net/http"
	"test/internal/constants"

	"github.com/go-chi/chi/v5"
)

func AuthRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post(constants.AuthLogin, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Auth Login OK"))
	})

	return r
}
