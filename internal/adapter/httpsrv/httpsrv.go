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

type ServerParams struct {
}

// NewServer creates a new server.
func NewServer(p ServerParams) *Server {
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
	h := route.Handler

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
