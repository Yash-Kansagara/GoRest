package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Yash-Kansagara/GoRest/internal/db"
	"github.com/Yash-Kansagara/GoRest/internal/server"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Failed to load Environment")
	}

	if _, err := db.ConnectDB(); err != nil {
		panic("could not connect to DB")
	}

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
