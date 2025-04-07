package htmx

import (
	"net/http"

	"github.com/rkperes/blog/components"
	"github.com/rkperes/blog/internal/core/services/authctx"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	session := authctx.Session(r.Context())
	components.Page("Home", session, nil).Render(r.Context(), w)
}
