package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", ":8080", "listening address")
	dir    = flag.String("dir", ".", "directory to server")
)

func main() {
	flag.Parse()
	log.Printf("listening to %q....", listen)
	log.Fatal(http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir))))
}
