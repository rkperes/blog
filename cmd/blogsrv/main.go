package main

import (
	"flag"
	"log"

	"github.com/rkperes/blog/internal/adapter/htmx"
	"github.com/rkperes/blog/internal/adapter/httpsrv"
)

var _port = flag.Int("p", 3000, "port")

func main() {
	flag.Parse()
	port := *_port

	srv := httpsrv.NewServer()

	htmxHandler := htmx.NewHandler()
	htmxHandler.RegisterRoutes(srv)

	log.Printf("Serving http://localhost:%d", port)
	err := srv.Serve(port)
	log.Fatal(err)
}
