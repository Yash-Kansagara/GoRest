package server

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Yash-Kansagara/GoRest/internal/middlewares"
	"golang.org/x/net/http2"
)

var enableTLS bool = true

func loadTLS() *tls.Config {
	conf := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
	return conf
}

// start http server
func Start() {

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":9090",
		Handler: ApplyMiddlewares(mux),
	}

	if enableTLS {
		srv.TLSConfig = loadTLS()
	}

	defer srv.Close()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/product", productHandler)

	err := http2.ConfigureServer(srv, &http2.Server{})
	if err != nil {
		log.Fatal("Failed Confuguring http2 server", err)
	}
	fmt.Println("server running...")

	if enableTLS {
		err = srv.ListenAndServeTLS("cert.pem", "key.pem")
		if err != nil {
			log.Fatal("Failed starting https server", err)
		}
	} else {
		err = srv.ListenAndServe()
		if err != nil {
			log.Fatal("Failed starting http server", err)
		}
	}

}

func ApplyMiddlewares(mux *http.ServeMux) http.Handler {
	// applied/runs from bottom
	handler := middlewares.CompressionMiddleware(mux)
	handler = middlewares.SecurityHeaders(handler)
	handler = middlewares.Cors(handler)
	return handler
}

func logReqDetails(req *http.Request) {
	log.Println("Path: ", req.URL.Path)
	log.Println("Method: ", req.Method)
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	logReqDetails(r)
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("products GET"))
	case http.MethodPost:
		w.Write([]byte("products POST"))
	case http.MethodPatch:
		w.Write([]byte("products PATCH"))
	case http.MethodPut:
		w.Write([]byte("products PUT"))
	case http.MethodDelete:
		w.Write([]byte("products DELETE"))
	}
}

func handleRoot(res http.ResponseWriter, req *http.Request) {
	logReqDetails(req)

	io.WriteString(res, "Works\n")
}
