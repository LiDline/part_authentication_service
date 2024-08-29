package auth

import (
	"encoding/json"
	"net/http"
	"test/internal/constants"
	customTypes "test/internal/types"

	"github.com/go-chi/chi/v5"
)

func AuthRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post(constants.AUTH_LOGIN, func(w http.ResponseWriter, r *http.Request) {

		var req customTypes.LoginRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)

			return
		}

		// user := authRouterServices.GetUserByGUID(req)

		// if !user {
		// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)

		// 	return
		// }

		w.Write([]byte("Auth Login OK"))
	})

	return r
}
