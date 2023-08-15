package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/TechBowl-japan/go-stations/model"
)

func BasicAuth(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// get basic auth from request
		userId, password, ok := r.BasicAuth()

		// change this condition to your own authentication
		if !ok || userId != os.Getenv("BASIC_AUTH_USER_ID") || password != os.Getenv("BASIC_AUTH_PASSWORD") {
			// if authentication failed, return 401 error
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)

			// どちらが適切なんだろう
			// w.Write([]byte(`{"message": "Unauthorized"}`))
			errorMessage := model.ErrUnauthorized{Message: "Unauthorized"}
			if err := json.NewEncoder(w).Encode(errorMessage); err != nil {
				log.Println(err)
			}

			return
		}
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
