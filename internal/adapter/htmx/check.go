package htmx

import (
	"context"
	"net/http"

	"github.com/rkperes/blog/components"
)

func (h *Handler) Check(w http.ResponseWriter, r *http.Request) {
	components.Check(true, "test").Render(context.Background(), w)
}
