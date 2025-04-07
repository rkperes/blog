package main

import (
	"flag"
	"log"

	"github.com/rkperes/blog/internal/adapter/htmx"
	"github.com/rkperes/blog/internal/adapter/httpsrv"
	"github.com/rkperes/blog/internal/adapter/repository"
)

var _port = flag.Int("p", 3000, "port")

func main() {
	flag.Parse()
	port := *_port

	repository, err := repository.NewSQLiteRepository()
	if err != nil {
		log.Fatal(err)
	}

	srv := httpsrv.NewServer(httpsrv.ServerParams{
		SessionRepository: repository,
	})

	htmxHandler := htmx.NewHandler(repository)
	htmxHandler.RegisterRoutes(srv)

	log.Printf("Serving http://localhost:%d", port)
	err = srv.Serve(port)
	log.Fatal(err)
}
