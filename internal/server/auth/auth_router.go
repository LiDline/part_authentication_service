package auth

import (
	"encoding/json"
	"net/http"
	"test/internal/constants"
	authservices "test/internal/server/auth/services"
	customTypes "test/internal/types"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func AuthRouter() *chi.Mux {
	r := chi.NewRouter()

	validate := validator.New()

	r.Post(constants.AUTH_LOGIN, func(w http.ResponseWriter, r *http.Request) {

		var req customTypes.LoginRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if err := validate.Struct(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := authservices.CreateTokens(req)

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		w.Write([]byte(user))
	})

	return r
}
