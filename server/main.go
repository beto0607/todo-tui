package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"todo.home/server/core"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}

	dbConnection := connectToDB()
	defer disconnectDB(dbConnection)

	startServer()

	log.Println("final")
}

func startServer() {
	port := os.Getenv("APP_PORT")
	if len(port) == 0 {
		port = "9091"
	}
	host := os.Getenv("APP_HOST")
	if len(host) == 0 {
		host = "todo.local"
	}

	server := core.InitServer(host, port)
	log.Print("Server configured")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go func() {
		log.Println("Starting server on " + server.Addr)
		server.ListenAndServe()
	}()

	<-ctx.Done()

	server.Shutdown(context.TODO())
	log.Println("Server shutdown")
}

func connectToDB() *core.DBConnection {
	mongoDbConnectionString := os.Getenv("MONGODB_CONNECTIONSTRING")
	dbConnection, err := core.InitDB(mongoDbConnectionString)
	if err != nil {
		log.Fatal("Could not connect to DB")
	}
	log.Print("DB connected correctly")
	return dbConnection
}

func disconnectDB(dbConnection *core.DBConnection) {
	err := dbConnection.Disconnect(context.Background())
	if err != nil {
		log.Fatal("Could not disconnect to DB")
	}
	log.Print("DB disconnected correctly")
}
