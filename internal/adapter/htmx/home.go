package htmx

import (
	"net/http"
	"path/filepath"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	// initiate a new htmx handler
	hh := h.htmx.NewHandler(w, r)

	// set the headers for the response, see docs for more options
	hh.PushURL("http://push.url")
	hh.ReTarget("#ReTarged")

	// write the output like you normally do.
	// check inspector tool in browser to see that the headers are set.
	http.ServeFile(w, r, "./pages/index.html")
}

func (h *Handler) HTMXPageHandler(page string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// // initiate a new htmx handler
		// hh := h.htmx.NewHandler(w, r)

		// // set the headers for the response, see docs for more options
		// hh.PushURL("http://push.url")
		// hh.ReTarget("#ReTarged")

		// write the output like you normally do.
		// check inspector tool in browser to see that the headers are set.
		http.ServeFile(w, r, filepath.Join("./pages/", page))
	}
}
