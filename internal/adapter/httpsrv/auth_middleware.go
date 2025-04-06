package httpsrv

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/rkperes/blog/internal/core/domain/errs"
	"github.com/rkperes/blog/internal/core/ports"
	"github.com/rkperes/blog/internal/core/services/authctx"
)

const (
	authCookie string = "session_token"
)

type authMiddleware struct {
	sessionRepository ports.SessionRepository
}

func newAuthMiddleware(sessionRepository ports.SessionRepository) *authMiddleware {
	return &authMiddleware{
		sessionRepository: sessionRepository,
	}
}

func (am *authMiddleware) MustAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(am.handle(next))
}

func (am *authMiddleware) handle(next http.Handler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ck, err := r.Cookie(authCookie)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				slog.Debug("no cookie found")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			slog.Error("unexpected error getting cookies", slog.Any("error", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		session, err := am.sessionRepository.GetSession(ck.Value)
		if err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				slog.Warn("session not found", slog.String("sessionID", ck.Value), slog.Any("error", err))
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			slog.Error("unexpected error getting session", slog.String("sessionID", ck.Value), slog.Any("error", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ctx = authctx.WithSession(ctx, session)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
