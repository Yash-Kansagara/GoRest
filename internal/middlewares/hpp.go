package middlewares

import (
	"net/http"
	"slices"
)

type HppMiddlewareConfig struct {
	VerifyBodyForm      bool
	VerifyQuery         bool
	AllowedQueryKeys    []string
	AllowedBodyFormKeys []string
}

func HPPMiddleware(next http.Handler, config *HppMiddlewareConfig) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// first verify url query, it will get passed to r.Form as well when r.ParseForm() is done
		if config.VerifyQuery {
			query := r.URL.Query()
			for k, _ := range query {
				if slices.Contains(config.AllowedQueryKeys, k) == false {
					query.Del(k)
				}
			}
			r.URL.RawQuery = query.Encode()
		}

		// verify form data
		// r.Form = url query + body's formdata when Content-Type is x-www-form-urlencoded
		// r.PostForm = body's formdata when Content-Type is x-www-form-urlencoded
		if config.VerifyBodyForm {

			r.ParseForm() // parses only if Content-Type is x-www-form-urlencoded
			for k, _ := range r.PostForm {
				if slices.Contains(config.AllowedBodyFormKeys, k) == false {
					delete(r.PostForm, k)
					delete(r.Form, k)
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}
