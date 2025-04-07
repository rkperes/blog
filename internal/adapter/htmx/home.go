package htmx

import (
	"net/http"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	// initiate a new htmx handler
	hh := h.htmx.NewHandler(w, r)

	// set the headers for the response, see docs for more options
	hh.PushURL("http://push.url")
	hh.ReTarget("#ReTarged")

	// mock session
	// TODO: delete this
	s, err := h.sessionRepo.CreateSession(r.Context(), "test-user-session")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var authCookie string = "session_token"
	http.SetCookie(w, &http.Cookie{
		Name:     authCookie,
		Value:    string(s.ID),
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
	// end mock session

	// write the output like you normally do.
	// check inspector tool in browser to see that the headers are set.
	http.ServeFile(w, r, "./pages/index.html")
}
