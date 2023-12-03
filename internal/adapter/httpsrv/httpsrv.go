package httpsrv

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

// Server is the http server entity.
// It implements the http serving and multiplexing.
type Server struct {
	r chi.Router
}

// NewServer creates a new server.
func NewServer() *Server {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	return &Server{
		r: r,
	}
}

type Route struct {
	Path    string
	Handler http.Handler
}

func (s *Server) RegisterRoute(route Route) error {
	s.r.Handle(route.Path, route.Handler)

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
