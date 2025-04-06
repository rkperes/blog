package htmx

import (
	"net/http"

	"github.com/donseba/go-htmx"
	"github.com/donseba/go-htmx/middleware"
	"github.com/rkperes/blog/internal/adapter/httpsrv"
)

type Handler struct {
	htmx          *htmx.HTMX
	addMiddleware addMiddlewareFunc
}

type addMiddlewareFunc func(next http.Handler) http.Handler

func NewHandler() *Handler {
	addMW := middleware.MiddleWare

	return &Handler{
		htmx:          htmx.New(),
		addMiddleware: addMW,
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
			Path:         "/check",
			Handler:      http.HandlerFunc(h.Check),
			AuthRequired: true,
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
