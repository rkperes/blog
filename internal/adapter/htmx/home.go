package htmx

import "net/http"

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	// initiate a new htmx handler
	hh := h.htmx.NewHandler(w, r)

	// set the headers for the response, see docs for more options
	hh.PushURL("http://push.url")
	hh.ReTarget("#ReTarged")

	// write the output like you normally do.
	// check inspector tool in browser to see that the headers are set.
	_, _ = hh.Write([]byte("OK"))
}
