package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/mileusna/useragent"
)

func Os(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		k := model.Key("os")
		us := useragent.Parse(r.UserAgent())
		ctx := context.WithValue(r.Context(), k, fmt.Sprintf("%v", us))
		h.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
