package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime/trace"

	"github.com/GeekchanskiY/cv_builder/pkg/db"
	rt "github.com/GeekchanskiY/cv_builder/pkg/router"
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

	// Start server
	r := rt.CreateRoutes()
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", os.Getenv("server_host"), os.Getenv("server_port")),
		Handler: r,
	}
	log.Println("Server started!")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
