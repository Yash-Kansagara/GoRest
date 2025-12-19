package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Yash-Kansagara/GoRest/internal/server"
)

func main() {
	flag.Parse()
	// graceful shutdown of process
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	// start http server
	go server.Start()

	// wait for interrupt
	log.Println("main waiting...")
	<-ctx.Done()
	log.Println("Shutting down...")
}
