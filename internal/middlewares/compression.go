package middlewares

import (
	"log"
	"net/http"
	"strings"

	utils "github.com/Yash-Kansagara/GoRest/internal/Utils"
)

func CompressionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Compression->")
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			w.Header().Set("Content-Encoding", "gzip")
			gzWriter, close := utils.NewGzipWriter(w)
			defer close()
			next.ServeHTTP(gzWriter, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
