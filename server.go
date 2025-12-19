package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func handleRoot(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Works")
}

func main() {

	// graceful shutdown of process
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	// start http server
	go func() {
		http.HandleFunc("/", handleRoot)
		log.Println("http server starting...")
		err := http.ListenAndServe("127.0.0.1:9090", nil)
		if err != nil {
			log.Fatalln("Error starting http server")
		}
	}()

	// wait for interrupt
	log.Println("main waiting...")
	<-ctx.Done()
	log.Println("Shutting down...")
}
