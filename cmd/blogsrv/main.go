package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	htmx "github.com/donseba/go-htmx"
	"github.com/donseba/go-htmx/middleware"
)

var _port = flag.Int("p", 3000, "port")

type App struct {
	htmx *htmx.HTMX
}

func main() {
	flag.Parse()
	port := *_port

	// new app with htmx instance
	app := &App{
		htmx: htmx.New(),
	}

	mux := http.NewServeMux()
	// wrap the htmx example middleware around the http handler
	mux.Handle("/", middleware.MiddleWare(http.HandlerFunc(app.Home)))

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	log.Fatal(err)
}

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	// initiate a new htmx handler
	h := a.htmx.NewHandler(w, r)

	// set the headers for the response, see docs for more options
	h.PushURL("http://push.url")
	h.ReTarget("#ReTarged")

	// write the output like you normally do.
	// check inspector tool in browser to see that the headers are set.
	_, _ = h.Write([]byte("OK"))
}
