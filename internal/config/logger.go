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
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	// Logging to file and stdout
	loggerOutput := io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(loggerOutput)

	// Trace file
	traceFile, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}

	// Start trace
	if err = trace.Start(traceFile); err != nil {
		fmt.Printf("failed to start trace: %v\n", err)
		return
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			trace.Stop()
			err = traceFile.Close()
			if err != nil {
				log.Fatal(err)
			}
			err = logFile.Close()
			if err != nil {
				log.Fatal(err)
			}
			return nil
		},
	})
}
