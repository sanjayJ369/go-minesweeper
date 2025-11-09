package main

import (
	"echo-server/logger"
	"echo-server/maelstrom"
	"os"
)

func main() {
	lgr := logger.NewStreamLogger(os.Stderr)
	srv := maelstrom.NewEchoServer(os.Stdin, os.Stdout, lgr)
	srv.ListenAndServe()
}
