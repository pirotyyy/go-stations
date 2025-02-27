package middleware

import (
	"fmt"
	"net/http"
)

func Recovery(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// ServeHTTP内でpanicが起こるとrecover
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println("Recover!: ", err)
			}
		}()
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
