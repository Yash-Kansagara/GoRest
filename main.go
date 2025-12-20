package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Yash-Kansagara/GoRest/internal/db"
	"github.com/Yash-Kansagara/GoRest/internal/server"
	"github.com/joho/godotenv"
)

func main() {

	// load env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Failed to load Environment")
	}

	// connect to db
	if _, err := db.ConnectDB(); err != nil {
		panic("could not connect to DB")
	}

	// graceful shutdown of process
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	// start http server
	go server.Start()

	// wait for interrupt to shutdown
	<-ctx.Done()
	log.Println("Shutdown")
}
