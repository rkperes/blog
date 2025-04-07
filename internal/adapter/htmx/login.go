package htmx

import (
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/rkperes/blog/components"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		slog.Error("failed to parse form", slog.Any("error", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	formUser := r.FormValue("user")
	formPw := r.FormValue("password")
	slog.Info("login", slog.String("user", formUser), slog.String("password", formPw))

	// TODO: check password and retrieve user ID
	userID := uuid.New()

	s, err := h.sessionRepo.CreateSession(r.Context(), userID.String())
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

	components.Login(s).Render(r.Context(), w)
}
