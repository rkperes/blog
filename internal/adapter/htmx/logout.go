package htmx

import (
	"net/http"

	"github.com/rkperes/blog/components"
	"github.com/rkperes/blog/internal/core/domain"
)

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	// TODO: invalidate session

	var authCookie string = "session_token"
	http.SetCookie(w, &http.Cookie{
		Name:     authCookie,
		Value:    "",
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	components.Login(domain.NoSession).Render(r.Context(), w)
}
