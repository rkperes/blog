package htmx

import (
	"context"
	"net/http"

	"github.com/rkperes/blog/components"
	"github.com/rkperes/blog/internal/core/services/authctx"
)

func (h *Handler) Check(w http.ResponseWriter, r *http.Request) {
	session, isLoggedIn := authctx.Session(r.Context())
	components.Check(isLoggedIn, string(session.UserID)).Render(context.Background(), w)
}
