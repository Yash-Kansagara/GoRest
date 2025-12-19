package middlewares

import (
	"log"
	"net"
	"net/http"

	utils "github.com/Yash-Kansagara/GoRest/internal/Utils"
)

/*
rate limits requests based on client id (ip, name, token etc)
reqPerSec = number of requests per second for each unique client id
burst = max requests at once allowed when enough tokens are available
*/
func RateLimiterMiddleware(next http.Handler, reqPerSec int, burst int) http.Handler {
	rateLimiter := utils.NewRateLimiter(10, 10)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err == nil {
			if rateLimiter.Allow(host) {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests) // code 429
			}
		} else {
			// log it but allow request anyway, can be restricted
			log.Println("RateLimiter: Error getting host from remote address", r.RemoteAddr, err)
			next.ServeHTTP(w, r)
		}
	})
}
