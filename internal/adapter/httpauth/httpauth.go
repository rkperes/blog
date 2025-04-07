package httpauth

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/rkperes/blog/internal/core/domain"
	"github.com/rkperes/blog/internal/core/domain/errs"
	"github.com/rkperes/blog/internal/core/ports"
	"github.com/rkperes/blog/internal/core/services/authctx"
)

const (
	authCookie = "session_token"
)

type HTTPAuth struct {
	sessionRepository ports.SessionRepository
}

func New(sessionRepository ports.SessionRepository) *HTTPAuth {
	return &HTTPAuth{
		sessionRepository: sessionRepository,
	}
}

// MustAuth is a middleware that checks if the user is authenticated
// and populates the request context with the session information.
// It returns a 401 Unauthorized status if the user is not authenticated.
func (am *HTTPAuth) MustAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(am.handleAuth(next, true))
}

// MaybeAuth is a middleware that checks if the user is authenticated
// and populates the request context with the session information.
// It continues the request even if the session is missing, not setting the context then.
func (am *HTTPAuth) MaybeAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(am.handleAuth(next, false))
}

func (am *HTTPAuth) handleAuth(next http.Handler, must bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := am.getSession(r)
		if err != nil {
			if !errors.Is(err, errNoAuth) {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if must {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		ctx := r.Context()
		if session != domain.NoSession {
			ctx = authctx.WithSession(ctx, session)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

var errNoAuth = errors.New("no auth means found")

func (am *HTTPAuth) getSession(r *http.Request) (domain.Session, error) {
	ck, err := r.Cookie(authCookie)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return domain.NoSession, errNoAuth
		} else {
			return domain.NoSession, err
		}
	}

	session, err := am.sessionRepository.GetSession(r.Context(), ck.Value)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			slog.Warn("session not found", slog.String("sessionID", ck.Value), slog.Any("error", err))
			return domain.NoSession, errNoAuth
		}
		slog.Error("unexpected error getting session", slog.String("sessionID", ck.Value), slog.Any("error", err))
	}

	return session, nil
}
