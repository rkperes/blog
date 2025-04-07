package htmx

import (
	"net/http"

	"github.com/donseba/go-htmx"
	"github.com/rkperes/blog/internal/adapter/httpauth"
	"github.com/rkperes/blog/internal/adapter/httpsrv"
	"github.com/rkperes/blog/internal/core/ports"
)

type addMiddlewareFunc func(next http.Handler) http.Handler

type Handler struct {
	htmx *htmx.HTMX

	addMiddleware func(http.Handler) http.Handler

	auth        *httpauth.HTTPAuth
	sessionRepo ports.SessionRepository
}

func NewHandler(sessionRepo ports.SessionRepository) *Handler {
	auth := httpauth.New(sessionRepo)
	return &Handler{
		htmx:          htmx.New(),
		addMiddleware: auth.MaybeAuth,
		auth:          auth,
		sessionRepo:   sessionRepo,
	}
}

type RouteRegister interface {
	RegisterRoute(route httpsrv.Route) error
	RegisterRoutes(route []httpsrv.Route) error
}

func (h *Handler) RegisterRoutes(reg RouteRegister) error {
	reg.RegisterRoutes(h.routes())

	return nil
}

func (h *Handler) routes() []httpsrv.Route {
	routes := []httpsrv.Route{
		{
			Path:    "/",
			Handler: http.HandlerFunc(h.Home),
		},
		{
			Path:    "/search",
			Handler: http.HandlerFunc(h.SearchPokemon),
		},
		{
			Path:    "/login",
			Handler: http.HandlerFunc(h.Login),
		},
		{
			Path:    "/logout",
			Handler: http.HandlerFunc(h.Logout),
		},
	}

	return h.wrapRoutesWithMiddleware(routes)
}

func (h *Handler) wrapRoutesWithMiddleware(routes []httpsrv.Route) []httpsrv.Route {
	for idx, route := range routes {
		routes[idx] = h.wrapRouteWithMiddleware(route)
	}
	return routes
}

func (h *Handler) wrapRouteWithMiddleware(route httpsrv.Route) httpsrv.Route {
	if h.addMiddleware == nil {
		return route
	}

	route.Handler = h.addMiddleware(route.Handler)
	return route
}
