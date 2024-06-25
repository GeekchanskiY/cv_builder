package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime/trace"

	"github.com/GeekchanskiY/cv_builder/internal"
	database "github.com/GeekchanskiY/cv_builder/pkg/db"
	"github.com/GeekchanskiY/cv_builder/pkg/router"
)

func main() {

	// Logging file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer file.Close()

	// Logging to file and stdout
	logger_output := io.MultiWriter(file, os.Stdout)
	log.SetOutput(logger_output)

	// Trace file
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Start trace
	if err := trace.Start(f); err != nil {
		fmt.Printf("failed to start trace: %v\n", err)
		return
	}
	defer trace.Stop()

	// Connect to database
	connection, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	internal.Samples(connection)

	log.Println("Server starting...")

	router := router.CreateRoutes()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("server_port")), router))

}
