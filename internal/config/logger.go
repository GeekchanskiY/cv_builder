package config

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/trace"

	"go.uber.org/fx"
)

func SetupLogger(lc fx.Lifecycle) {
	// Logging file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	// Logging to file and stdout
	logger_output := io.MultiWriter(file, os.Stdout)
	log.SetOutput(logger_output)

	// Trace file
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}

	// Start trace
	if err := trace.Start(f); err != nil {
		fmt.Printf("failed to start trace: %v\n", err)
		return
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			trace.Stop()
			f.Close()
			file.Close()
			return nil
		},
	})
}
