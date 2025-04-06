package httpsrv

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rkperes/blog/internal/core/ports"
)

// Server is the http server entity.
// It implements the http serving and multiplexing.
type Server struct {
	r chi.Router

	authMw *authMiddleware
}

type ServerParams struct {
	sessionRepository ports.SessionRepository
}

// NewServer creates a new server.
func NewServer(p ServerParams) *Server {
	authMw := newAuthMiddleware(p.sessionRepository)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	return &Server{
		r:      r,
		authMw: authMw,
	}
}

type Route struct {
	Path         string
	Handler      http.Handler
	AuthRequired bool
}

func (s *Server) RegisterRoute(route Route) error {
	h := route.Handler

	if route.AuthRequired {
		h = s.authMw.MustAuth(h)
	}

	s.r.Handle(route.Path, h)

	return nil
}

func (s *Server) RegisterRoutes(routes []Route) error {
	for _, route := range routes {
		if err := s.RegisterRoute(route); err != nil {
			return err
		}
	}

	return nil
}

// Serve serves the http(s) server in the given port.
// It is a synchronous (blocking) call.
func (s *Server) Serve(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), s.r)
}
