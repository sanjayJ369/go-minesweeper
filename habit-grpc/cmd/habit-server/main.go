package main

import (
	"habit-grpc/internal/server"
	"habit-grpc/log"
	"os"
)

const port = 28710 // port

func main() {
	lgr := log.NewLogger(os.Stdout)

	srv := server.New(lgr)

	if err := srv.ListenAndServer(port); err != nil {
		lgr.Logf("error while running the server:%w", err)
		os.Exit(1)
	}
}
