package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/TechBowl-japan/go-stations/model"
)

func AccessLog(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		accessLog := model.AccessLog{
			Timestamp: time.Now(),
			Path:      r.URL.Path,
		}
		// var accessLog model.AccessLog

		// set os name
		key := model.Key("os")
		if osValue, ok := r.Context().Value(key).(string); ok {
			accessLog.OS = osValue
		}

		defer func() {
			// set latency
			accessLog.Latency = time.Since(accessLog.Timestamp).Milliseconds()
			// accessLogJson, err := json.Marshal(accessLog)
			// if err != nil {
			// 	log.Println(err)
			// 	return
			// }
			fmt.Printf("%#v\n", accessLog)
		}()
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
