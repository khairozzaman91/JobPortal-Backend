package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {

	logger := func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Println(r.Method, r.URL.Path, time.Since(start))

	}

	return http.HandlerFunc(logger)
}
